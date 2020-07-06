package main

import (
    cloudevents "github.com/cloudevents/sdk-go/v2"
    "github.com/rs/zerolog/log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"

    "github.com/xmlking/grpc-starter-kit/shared/middleware/rpclog"
    appendTags "github.com/xmlking/grpc-starter-kit/shared/middleware/tags/append"
    "github.com/xmlking/grpc-starter-kit/shared/middleware/translog"
    "github.com/xmlking/grpc-starter-kit/toolkit/service"
    "github.com/xmlking/grpc-starter-kit/toolkit/util/tls"

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
        tlsConf, err := tls.NewTLSConfig(cfg.Features.Tls.CertFile, cfg.Features.Tls.KeyFile, cfg.Features.Tls.CaFile, cfg.Features.Tls.ServerName)
        if err != nil {
            log.Fatal().Err(err).Msg("failed to create cert")
        }
        serverCert := credentials.NewTLS(tlsConf)
        grpcOps = append(grpcOps, grpc.Creds(serverCert))
    }

    srv := service.NewService(
        service.Name(serviceName),
        service.Version(cfg.Services.Account.Version),
        service.WithGrpcEndpoint(cfg.Services.Account.Endpoint),
        service.WithGrpcOptions(grpcOps...),
        // service.WithBrokerOptions(...),
    )
    // create a gRPC server object
    grpcServer := srv.Server()

    // Register Handlers
    userv1.RegisterUserServiceServer(grpcServer, userHandler)
    profilev1.RegisterProfileServiceServer(grpcServer, profileHandler)

    // start the server
    if err := srv.Start(); err != nil {
        log.Fatal().Err(err).Send()
    }
}
