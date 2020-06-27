package main

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/xmlking/grpc-starter-kit/shared/middleware/rpclog"
	appendTags "github.com/xmlking/grpc-starter-kit/shared/middleware/tags/append"
	"github.com/xmlking/grpc-starter-kit/shared/middleware/translog"

	profilev1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
	userv1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
	greeterv1 "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/account/handler"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/xmlking/grpc-starter-kit/service/account/registry"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	forwardTags "github.com/xmlking/grpc-starter-kit/shared/middleware/tags/forward"

	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func main() {
	serviceName := constants.ACCOUNT_SERVICE
	cfg := config.GetConfig()

	lis, err := config.GetListener(cfg.Services.Account.Endpoint)
	if err != nil {
		log.Fatal().Msgf("failed to create listener: %v", err)
	}

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}

	translogPublisher := ctn.Resolve("translog-publisher").(cloudevents.Client)
	emailPublisher := ctn.Resolve("email-publisher").(cloudevents.Client)
	greeterSrvClient := ctn.Resolve("greeter-client").(greeterv1.GreeterServiceClient)

	// Handlers
	userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), emailPublisher, greeterSrvClient)
	profileHandler := ctn.Resolve("profile-handler").(profilev1.ProfileServiceServer)

	// create a gRPC server object
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			// Execution is done in left-to-right order
			// keep around type Interceptors first,
			rpclog.UnaryServerInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			appendTags.UnaryServerInterceptor(appendTags.WithPairs(constants.FromServiceKey, constants.ACCOUNT_SERVICE)),
			forwardTags.UnaryServerInterceptor(forwardTags.WithForwardTags(constants.TraceIDKey, constants.TenantIdKey)),
			translog.UnaryServerInterceptor(translogPublisher, serviceName),
		)),
	)

	// Register Handlers
	userv1.RegisterUserServiceServer(grpcServer, userHandler)
	profilev1.RegisterProfileServiceServer(grpcServer, profileHandler)

	// Add HealthChecks
	hsrv := health.NewServer()
	for name := range grpcServer.GetServiceInfo() {
		hsrv.SetServingStatus(name, grpc_health_v1.HealthCheckResponse_SERVING)
	}
	grpc_health_v1.RegisterHealthServer(grpcServer, hsrv)

	// Start server!
	reflection.Register(grpcServer)
	println(config.GetBuildInfo())
	log.Info().Msgf("Server (%s) started at: %s, secure: %t", serviceName, lis.Addr(), cfg.Features.Tls.Enabled)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Send()
	}
}
