package main

import (
    "github.com/rs/zerolog/log"
    "github.com/xmlking/grpc-starter-kit/shared/broker"
    "github.com/xmlking/grpc-starter-kit/shared/config"
)

var (
    cfg = config.GetConfig()
)

func main() {
    // os.Setenv("PUBSUB_EMULATOR_HOST", "http://localhost:8085")
    // os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "~/my-json.json")

    broker.DefaultBroker = broker.NewBroker(broker.ProjectID("my-project-id"));

    myHandler := func(e broker.Event) error {
        log.Info().Interface("event", e).Send()
        return  nil
    }

    sub, err := broker.Subscribe("ingestion-in-dev", myHandler, broker.Queue("ingestion-in-dev"));
    if err != nil {
        log.Error().Err(err).Send()
    }
    sub.Unsubscribe()
}

