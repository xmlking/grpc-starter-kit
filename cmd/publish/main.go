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
    // broker.DefaultBroker = broker.NewBroker(broker.ProjectID("my-project-id")); // use cfg.pubsub.ProjectID
    broker.DefaultBroker = broker.NewBroker();

    header := map[string]string{"sumo": "demo"}
    body   := []byte("ABC€")
    msg := broker.Message{Header: header, Body: body}

    if err := broker.Publish("ingestion-in-dev", &msg); err != nil {
        log.Error().Err(err).Send()
    }
}
