package config

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/configor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/xmlking/grpc-starter-kit/micro/middleware/rpclog"
	configPB "github.com/xmlking/grpc-starter-kit/shared/proto/config"
	uTLS "github.com/xmlking/grpc-starter-kit/shared/util/tls"
)

var (
	Configor   *configor.Configor
	cfg        configPB.Configuration
	configLock = new(sync.RWMutex)

	// Version is populated by govvv in compile-time.
	Version = "untouched"
	// BuildDate is populated by govvv.
	BuildDate string
	// GitCommit is populated by govvv.
	GitCommit string
	// GitBranch is populated by govvv.
	GitBranch string
	// GitState is populated by govvv.
	GitState string
	// GitSummary is populated by govvv.
	GitSummary string
)

// VersionMsg is the message that is shown after process started.
const versionMsg = `
version     : %s
build date  : %s
go version  : %s
go compiler : %s
platform    : %s/%s
git commit  : %s
git branch  : %s
git state   : %s
git summary : %s
`

func init() {
	configPath, exists := os.LookupEnv("CONFIGOR_FILE_PATH")
	if !exists {
		configPath = "/config/config.yaml"
	}

	Configor = configor.New(&configor.Config{UsePkger: true, ErrorOnUnmatchedKeys: true})
	log.Info().Msgf("loading configuration from file: %s", configPath)
	if err := Configor.Load(&cfg, configPath); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Panic().Err(err).Msgf("missing config file at %s", configPath)
		} else {
			log.Fatal().Err(err).Send()
		}
	}
}

/**
  Helper Functions
*/

func GetBuildInfo() string {
	return fmt.Sprintf(versionMsg, Version, BuildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
		GitCommit, GitBranch, GitState, GitSummary)
}

func GetConfig() configPB.Configuration { // FIXME: return a deep copy?
	configLock.RLock()
	defer configLock.RUnlock()
	return cfg
}

func CreateServerCerts() (tlsConfig *tls.Config, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	tlsConf := cfg.Features.Tls
	return uTLS.GetTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.Servername)
}

func IsProduction() bool {
	return Configor.GetEnvironment() == "production"
}

func IsSecure() bool {
	configLock.RLock()
	defer configLock.RUnlock()
	return cfg.Features.Tls.Enabled
}

func GetCeClient() {
}

func GetClientConn(service *configPB.Service /*from_service*/) (clientConn *grpc.ClientConn, err error) {
	configLock.RLock()
	defer configLock.RUnlock()

	var dialOptions []grpc.DialOption
	var ucInterceptors []grpc.UnaryClientInterceptor

	tlsConf := cfg.Features.Tls
	if tlsConf.Enabled {
		if creds, err := uTLS.GetTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.Servername); err != nil {
			return nil, err
		} else {
			dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
		}
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	if service.ServiceConfig != "" {
		dialOptions = append(dialOptions, grpc.WithDefaultServiceConfig(service.ServiceConfig))
	}

	if cfg.Features.Rpclog.Enabled {
		ucInterceptors = append(ucInterceptors, rpclog.UnaryClientInterceptor())
	}
	// TODO append from_service Interceptor passed above, carry over existing tags

	if len(ucInterceptors) > 0 {
		dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(ucInterceptors...)))
	}
	clientConn, err = grpc.Dial(service.Endpoint, dialOptions...)
	if err != nil {
		log.Fatal().Msgf("Failed connect to: %s, error: %s", service.Endpoint, err)
	}
	return
}

func GetListener(endpoint string) (lis net.Listener, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	var target Target
	target, err = ParseTarget(endpoint)
	if err != nil {
		return
	}

	switch target.Scheme {
	case "unix":
		return net.Listen("unix", target.Path)
	case "tcp", "dns", "http", "https", "kubernetes":
		if target.Port == "" {
			target.Port = "0"
		}

		tlsConf := cfg.Features.Tls
		if tlsConf.Enabled {
			if tlsConfig, err := uTLS.GetTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.Servername); err != nil {
				return nil, err
			} else {
				return tls.Listen("tcp", fmt.Sprintf(":%s", target.Port), tlsConfig)
			}
		} else {
			return net.Listen("tcp", fmt.Sprintf(":%s", target.Port))
		}
	default:
		return nil, errors.New(fmt.Sprintf("unknown scheme: %s in endpoint: %s", target.Scheme, endpoint))
	}
}

type Target struct {
	Scheme string
	Host   string
	Port   string
	Path   string
}

// ParseTarget splits target into a Target struct containing scheme, host, port and path.
// If target is not a valid scheme://host:port/path, it returns error.
func ParseTarget(target string) (ret Target, err error) {
	if strings.HasPrefix(target, "unix://") {
		spl := strings.SplitN(target, "://", 2)
		ret.Scheme = spl[0]
		ret.Path = spl[1]
		return
	}
	target = strings.Replace(target, ":///", "://", 1)
	var u *url.URL
	if u, err = url.Parse(target); err == nil {
		ret.Scheme = u.Scheme
		ret.Path = u.Path
		if u.Host != "" {
			var portError error
			if ret.Host, ret.Port, portError = net.SplitHostPort(u.Host); portError != nil {
				log.Debug().Err(portError)
				ret.Host = u.Host
			}
		}
	}
	return
}
