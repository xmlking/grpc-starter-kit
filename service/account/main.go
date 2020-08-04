package main

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/rs/zerolog/log"
	"github.com/sercand/kuberesolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/service"
	"github.com/xmlking/toolkit/util/tls"

	"github.com/xmlking/grpc-starter-kit/shared/middleware/translog"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	profilev1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
	userv1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"github.com/xmlking/grpc-starter-kit/service/account/registry"

	appendTags "github.com/xmlking/toolkit/middleware/tags/append"
	forwardTags "github.com/xmlking/toolkit/middleware/tags/forward"

	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func main() {
	serviceName := constants.ACCOUNT_SERVICE
	cfg := config.GetConfig()

	// Register kuberesolver to grpc.
	// This line should be before calling registry.NewContainer(cfg)
	if config.IsProduction() {
		kuberesolver.RegisterInCluster()
	}

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}

	translogPublisher := ctn.Resolve("translog-publisher").(cloudevents.Client)

	// Handlers
	userHandler := ctn.Resolve("user-handler").(userv1.UserServiceServer)
	profileHandler := ctn.Resolve("profile-handler").(profilev1.ProfileServiceServer)

	// ServerOption
	grpcOps := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			// Execution is done in left-to-right order
			// keep around type Interceptors first,
			rpclog.UnaryServerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			appendTags.UnaryServerInterceptor(appendTags.WithPairs(constants.FromServiceKey, constants.ACCOUNT_SERVICE)),
			forwardTags.UnaryServerInterceptor(forwardTags.WithForwardTags(constants.TraceIDKey, constants.TenantIdKey)),
			translog.UnaryServerInterceptor(translogPublisher, serviceName),
		)),
	}

	if cfg.Features.Tls.Enabled {
		tlsConf, err := tls.NewTLSConfig(cfg.Features.Tls.CertFile, cfg.Features.Tls.KeyFile, cfg.Features.Tls.CaFile, cfg.Features.Tls.ServerName, cfg.Features.Tls.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cert")
		}
		serverCert := credentials.NewTLS(tlsConf)
		grpcOps = append(grpcOps, grpc.Creds(serverCert))
	}

	// DialOptions
	var dialOptions []grpc.DialOption
	var ucInterceptors []grpc.UnaryClientInterceptor

	tlsConf := cfg.Features.Tls
	if tlsConf.Enabled {
		if creds, err := tls.NewTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.ServerName, cfg.Features.Tls.Password); err != nil {
			log.Fatal().Err(err).Msg("Failed to create tlsConf")
		} else {
			dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
		}
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	if cfg.Features.Rpclog.Enabled {
		ucInterceptors = append(ucInterceptors, rpclog.UnaryClientInterceptor())
	}

	if len(ucInterceptors) > 0 {
		dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(ucInterceptors...)))
	}

	srv := service.NewService(
		service.Name(serviceName),
		service.Version(cfg.Services.Account.Version),
		service.WithGrpcEndpoint(cfg.Services.Account.Endpoint),
		service.WithGrpcOptions(grpcOps...),
		service.WithDialOptions(dialOptions...),
		// service.WithBrokerOptions(...),
	)

	// create a gRPC server object
	grpcServer := srv.Server()
	// greeterClientCon, err := srv.Client(service.Remote{ Endpoint: cfg.Services.Greeter.Endpoint, ServiceConfig: cfg.Services.Greeter.ServiceConfig } )

	// Register Handlers
	userv1.RegisterUserServiceServer(grpcServer, userHandler)
	profilev1.RegisterProfileServiceServer(grpcServer, profileHandler)

	// start the server
	log.Info().Msg(config.GetBuildInfo())
	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
