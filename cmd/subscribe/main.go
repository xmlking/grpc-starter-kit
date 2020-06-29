package main

import (
    "context"

    "cloud.google.com/go/pubsub"
    "github.com/rs/zerolog/log"
)


func main() {
    //os.Setenv("PUBSUB_EMULATOR_HOST", "http://localhost:8085")
    // os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "~/my-json.json")

    ctx := context.Background()
    c, err := pubsub.NewClient(ctx, "micro-starter-kit") // Go autodetects the PUBSUB_EMULATOR_HOST variable
    if err != nil {
        log.Error().Err(err).Send()
    }

    sub := c.Subscription("streaming-input")

    err = sub.Receive(ctx, func(cctx context.Context, msg *pubsub.Message) {
     println(msg)
    })
    log.Error().Err(err).Send()
}
