package main

import (
    "net/http"

    "github.com/rs/zerolog/log"
    "github.com/soheilhy/cmux"
    "google.golang.org/grpc"
    "google.golang.org/grpc/health"
    "google.golang.org/grpc/health/grpc_health_v1"

    "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
    "github.com/xmlking/grpc-starter-kit/service/greeter/handler"
    "github.com/xmlking/grpc-starter-kit/shared/config"
    "github.com/xmlking/grpc-starter-kit/shared/constants"
    _ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func main1() {
    serviceName := constants.GREETER_SERVICE
    cfg := config.GetConfig()

    lis, err := config.GetListener(cfg.Services.Greeter.Endpoint)
    if err != nil {
        log.Fatal().Msgf("failed to create listener: %v", err)
    }

    // Create a cmux.
    mux := cmux.New(lis)
    // Match connections in order:
    grpcL := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
    httpL := mux.Match(cmux.HTTP1Fast())


    // Create your protocol servers.
    grpcS := grpc.NewServer()
    greeterv1.RegisterGreeterServiceServer(grpcS, handler.NewGreeterHandler())

    // Register http Handlers
    httpS := &http.Server{
        Handler: handler.NewHttpHandler(),
    }

      hsrv := health.NewServer()
      for name, _ := range grpcS.GetServiceInfo() {
         hsrv.SetServingStatus(name, grpc_health_v1.HealthCheckResponse_SERVING)
      }
      grpc_health_v1.RegisterHealthServer(grpcS, hsrv)


    println(config.GetBuildInfo())
    println(serviceName)

    // Use the muxed listeners for your servers.
    go grpcS.Serve(grpcL)
    go httpS.Serve(httpL)

    // Start serving!
    mux.Serve()
}

func main() {
    serviceName := constants.GREETER_SERVICE
    cfg := config.GetConfig()

    println(serviceName)
    lis, err := config.GetListener(cfg.Services.Greeter.Endpoint)
    if err != nil {
        log.Fatal().Msgf("failed to create listener: %v", err)
    }

    // create a server instance
    s := handler.NewGreeterHandler()
    // create a gRPC server object
    grpcServer := grpc.NewServer()
    // attach the Ping service to the server
    greeterv1.RegisterGreeterServiceServer(grpcServer, s)
    // start the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatal().Err(err).Send()
    }
}
