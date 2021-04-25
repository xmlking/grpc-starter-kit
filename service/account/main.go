package main

import (
	_ "github.com/xmlking/toolkit/logger/auto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"github.com/sercand/kuberesolver"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	"github.com/xmlking/grpc-starter-kit/internal/middleware/translog"
	profilev1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
	userv1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
	"github.com/xmlking/grpc-starter-kit/service/account/registry"
	broker "github.com/xmlking/toolkit/broker/cloudevents"
	"github.com/xmlking/toolkit/middleware/rpclog"
	appendTags "github.com/xmlking/toolkit/middleware/tags/append"
	forwardTags "github.com/xmlking/toolkit/middleware/tags/forward"
	"github.com/xmlking/toolkit/service"
	"github.com/xmlking/toolkit/util/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	serviceName := constants.ACCOUNT_SERVICE
	cfg := config.GetConfig()
	efs := config.GetFileSystem()

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

	translogPublisher := ctn.Resolve("translog-publisher").(broker.Publisher)

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
			forwardTags.UnaryServerInterceptor(forwardTags.WithForwardTags(constants.TraceIDKey, constants.TenantIDKey)),
			translog.UnaryServerInterceptor(translogPublisher, serviceName),
		)),
	}

	if cfg.Features.TLS.Enabled {
		tlsConf, err := tls.NewTLSConfig(efs, cfg.Features.TLS.CertFile, cfg.Features.TLS.KeyFile, cfg.Features.TLS.CaFile, cfg.Features.TLS.ServerName, cfg.Features.TLS.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cert")
		}
		serverCert := credentials.NewTLS(tlsConf)
		grpcOps = append(grpcOps, grpc.Creds(serverCert))
	}

	// DialOptions
	var dialOptions []grpc.DialOption
	var ucInterceptors []grpc.UnaryClientInterceptor

	tlsConf := cfg.Features.TLS
	if tlsConf.Enabled {
		if creds, err := tls.NewTLSConfig(efs, tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.ServerName, cfg.Features.TLS.Password); err != nil {
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
