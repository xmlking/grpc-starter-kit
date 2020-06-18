package main

import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/balancer/roundrobin"

    //"google.golang.org/grpc/balancer/roundrobin"

    "github.com/rs/zerolog/log"

    profilev1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
    userv1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
    greeterv1 "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
    "github.com/xmlking/grpc-starter-kit/service/account/handler"

    "github.com/xmlking/grpc-starter-kit/service/account/registry"
    "github.com/xmlking/grpc-starter-kit/service/account/repository"
    "github.com/xmlking/grpc-starter-kit/shared/eventing"

    _ "github.com/jinzhu/gorm/dialects/sqlite"

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

    // Publisher publish to "mkit.service.emailer"
    greeterConn, err := grpc.Dial(cfg.Services.Greeter.Endpoint, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
    if err != nil {
       log.Fatal().Msgf("did not connect: %s", err)
    }
    defer greeterConn.Close()
    println(greeterConn.Target())

    publisher := eventing.NewSourceClient(cfg.Services.Emailer.Endpoint)
    // greeterSrv Client to call "mkit.service.greeter"
    greeterSrvClient := greeterv1.NewGreeterServiceClient(greeterConn)

    // // Handlers
    userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
    profileHandler := ctn.Resolve("profile-handler").(profilev1.ProfileServiceServer)

    // create a gRPC server object
    grpcServer := grpc.NewServer()
    // attach the Ping service to the server
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
