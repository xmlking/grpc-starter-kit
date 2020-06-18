package main

import (
    "context"

    "github.com/rs/zerolog/log"

    "github.com/xmlking/grpc-starter-kit/service/emailer/registry"
    "github.com/xmlking/grpc-starter-kit/service/emailer/subscriber"
    "github.com/xmlking/grpc-starter-kit/shared/config"
    "github.com/xmlking/grpc-starter-kit/shared/constants"
    _ "github.com/xmlking/grpc-starter-kit/shared/constants"
    "github.com/xmlking/grpc-starter-kit/shared/eventing"
    _ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func main() {
    serviceName := constants.EMAILER_SERVICE
    cfg := config.GetConfig()

    // Initialize DI Container
    ctn, err := registry.NewContainer(cfg)
    defer ctn.Clean()
    if err != nil {
        log.Fatal().Msgf("failed to build container: %v", err)
    }
    emailSubscriber := ctn.Resolve("emailer-subscriber").(*subscriber.EmailSubscriber)

    client := eventing.NewSinkClient(cfg.Services.Emailer.Endpoint)

    // Start server!
    println(config.GetBuildInfo())
    log.Info().Msgf("Server (%s) started at: %s", serviceName, cfg.Services.Emailer.Endpoint)
    if err := client.StartReceiver(context.Background(), emailSubscriber.HandleSend); err != nil {
       log.Fatal().Err(err).Send();
    }
}
