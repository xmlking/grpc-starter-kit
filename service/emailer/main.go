package main

import (
    "context"

    cloudevents "github.com/cloudevents/sdk-go/v2"

    "github.com/rs/zerolog/log"

    "github.com/xmlking/grpc-starter-kit/shared/broker"
    _ "github.com/xmlking/grpc-starter-kit/shared/constants"
    _ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func receive(event cloudevents.Event) {
    // do something with event.
    log.Info().Msgf("%s", event)
}

func main() {
    client := broker.DefaultClient
    if err := client.StartReceiver(context.Background(), receive); err != nil {
        log.Fatal().Err(err).Send();
    }
}
