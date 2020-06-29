package main

import (
    "context"

    "cloud.google.com/go/pubsub"
    "github.com/rs/zerolog/log"
)



func main() {
    // os.Setenv("PUBSUB_EMULATOR_HOST", "http://localhost:8085")
    // os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "~/my-json.json")

    ctx := context.Background()
    client, err := pubsub.NewClient(ctx, "micro-starter-kit")
    if err != nil {
        log.Error().Err(err).Send()
    }

    t := client.Topic("streaming-input")

    result := t.Publish(ctx, &pubsub.Message{
        Data: []byte("ABC€"),
    })
    // Block until the result is returned and a server-generated
    // ID is returned for the published message.
    id, err := result.Get(ctx)
    if err != nil {
        log.Error().Err(err).Send()
    }
    log.Info().Msgf("Published a message; msg ID: %v\n", id)
}


