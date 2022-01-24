package config

import (
	"io/fs"
	"os"
	"strings"
	"sync"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/toolkit/confy"
	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/util/tls"
	"github.com/xmlking/toolkit/util/xfs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"

	embed "github.com/xmlking/grpc-starter-kit"
)

var (
	once       sync.Once
	efs        fs.FS
	cfg        Configuration
	configLock = new(sync.RWMutex)
)

// Init you can call `Init()` explicitly one-time when app start, or `GetConfig()` implicitly initialize it.
func Init() {
	once.Do(func() { // <-- atomic, does not allow repeating
		efs = xfs.FS(embed.StaticConfig)
		configFiles, exists := os.LookupEnv("CONFY_FILES")
		if !exists {
			configFiles = "config/config.yml"
		}

		confy.DefaultConfy = confy.NewConfy(confy.WithFS(efs), confy.WithErrorOnUnmatchedKeys())

		log.Info().Msgf("loading config files: %s", configFiles)
		if err := confy.Load(&cfg, strings.Split(configFiles, ",")...); err != nil {
			if strings.Contains(err.Error(), "no such file") {
				log.Panic().Err(err).Msgf("missing config file at %s", configFiles)
			} else {
				log.Fatal().Err(err).Send()
			}
		}
	})
}

/**
  Helper Functions
*/

// GetFileSystem helper
func GetFileSystem() fs.FS {
	configLock.RLock()
	defer configLock.RUnlock()
	return efs
}

// GetConfig helper
func GetConfig() Configuration { // FIXME: return a deep copy?
	configLock.RLock()
	defer configLock.RUnlock()
	// HINT: initialize for the first time
	if confy.DefaultConfy == nil {
		Init()
	}
	return cfg
}

// IsProduction helper
func IsProduction() bool {
	return confy.GetEnvironment() == "production"
}

// IsSecure helper
func IsSecure() bool {
	configLock.RLock()
	defer configLock.RUnlock()
	return cfg.Features.TLS.Enabled
}

// GetClientConn helper
func GetClientConn(service *Service, ucInterceptors []grpc.UnaryClientInterceptor) (clientConn *grpc.ClientConn, err error) {
	configLock.RLock()
	defer configLock.RUnlock()

	var dialOptions []grpc.DialOption
	//var ucInterceptors []grpc.UnaryClientInterceptor

	tlsConf := cfg.Features.TLS
	if tlsConf.Enabled {
		if creds, err := tls.NewTLSConfig(efs, tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.ServerName, tlsConf.Password); err != nil {
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
