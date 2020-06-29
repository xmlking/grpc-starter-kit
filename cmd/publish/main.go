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

    header := map[string]string{"sumo": "demo"}
    body   := []byte("ABC€")
    msg := broker.Message{Header: header, Body: body}

    if err := broker.Publish("ingestion-in-dev", &msg); err != nil {
        log.Error().Err(err).Send()
    }
}
