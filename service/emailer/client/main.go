package main

import (
    "context"

    //cepubsub "github.com/cloudevents/sdk-go/protocol/pubsub/v2"
    cloudevents "github.com/cloudevents/sdk-go/v2"

    "github.com/rs/zerolog/log"

    "github.com/xmlking/grpc-starter-kit/shared/broker"
    _ "github.com/xmlking/grpc-starter-kit/shared/constants"
    _ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func main() {
    client := broker.DefaultClient

    // Create an Event.
    event :=  cloudevents.NewEvent()
    event.SetSource("example/uri")
    event.SetType("example.type")
    event.SetData(cloudevents.ApplicationJSON, map[string]string{"hello": "world"})

    // Set a target.
    ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")

    // Send that Event.
    if result := client.Send(ctx, event); !cloudevents.IsACK(result) {
        log.Fatal().Msgf("failed to send, %v", result)
    }
}
