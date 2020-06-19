package main

import (
    cloudevents "github.com/cloudevents/sdk-go/v2"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc"

    "github.com/xmlking/grpc-starter-kit/micro/middleware/rpclog"
    profilev1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
    userv1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
    greeterv1 "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
    "github.com/xmlking/grpc-starter-kit/service/account/handler"

    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/xmlking/grpc-starter-kit/service/account/registry"
    "github.com/xmlking/grpc-starter-kit/service/account/repository"

    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
    grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

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

    publisher := ctn.Resolve("email-publisher").(cloudevents.Client)
    greeterSrvClient := ctn.Resolve("greeter-client").(greeterv1.GreeterServiceClient)

    // Handlers
    userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
    profileHandler := ctn.Resolve("profile-handler").(profilev1.ProfileServiceServer)

    // create a gRPC server object
    grpcServer := grpc.NewServer(
        grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
            grpc_validator.UnaryServerInterceptor(),
            // keep it last in the interceptor chain
            rpclog.UnaryServerInterceptor(),
        )),
    )

    // Register Handlers
    userv1.RegisterUserServiceServer(grpcServer, userHandler)
    profilev1.RegisterProfileServiceServer(grpcServer, profileHandler)

    // start the server
    println(config.GetBuildInfo())
    log.Info().Msgf("Server (%s) started at: %s", serviceName, lis.Addr())
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatal().Err(err).Send()
    }
}
