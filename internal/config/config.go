package config

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/toolkit/configurator"
	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/util/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"

	//k8s resolver
	_ "github.com/tcfw/go-grpc-k8s-resolver"
)

var (
	cfg        Configuration
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
	configFiles, exists := os.LookupEnv("CONFIG_FILES")
	if !exists {
		configFiles = "/config/config.yml"
	}

	configurator.DefaultConfigurator = configurator.NewConfigurator(configurator.WithPkger(), configurator.WithErrorOnUnmatchedKeys())

	log.Info().Msgf("loading config files: %s", configFiles)
	if err := configurator.Load(&cfg, strings.Split(configFiles, ",")...); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			log.Panic().Err(err).Msgf("missing config file at %s", configFiles)
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

func GetConfig() Configuration { // FIXME: return a deep copy?
	configLock.RLock()
	defer configLock.RUnlock()
	return cfg
}

func IsProduction() bool {
	return configurator.GetEnvironment() == "production"
}

func IsSecure() bool {
	configLock.RLock()
	defer configLock.RUnlock()
	return cfg.Features.Tls.Enabled
}

func GetClientConn(service *Service, ucInterceptors []grpc.UnaryClientInterceptor) (clientConn *grpc.ClientConn, err error) {
	configLock.RLock()
	defer configLock.RUnlock()

	var dialOptions []grpc.DialOption
	//var ucInterceptors []grpc.UnaryClientInterceptor

	tlsConf := cfg.Features.Tls
	if tlsConf.Enabled {
		if creds, err := tls.NewTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.ServerName, tlsConf.Password); err != nil {
			return nil, err
		} else {
			dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
		}
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	if service.Authority != "" {
		dialOptions = append(dialOptions, grpc.WithAuthority(service.Authority))
	}

	if service.ServiceConfig != "" {
		dialOptions = append(dialOptions, grpc.WithDefaultServiceConfig(service.ServiceConfig))
	}

	if cfg.Features.Rpclog.Enabled {
		ucInterceptors = append(ucInterceptors, rpclog.UnaryClientInterceptor())
	}

	if len(ucInterceptors) > 0 {
		dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(ucInterceptors...)))
	}

	resolver.SetDefaultScheme("k8s") //optional

	clientConn, err = grpc.Dial(service.Endpoint, dialOptions...)
	return
}
