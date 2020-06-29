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
    broker.DefaultBroker = broker.NewBroker();

    myHandler := func(e broker.Event) error {
        log.Info().Interface("event", e).Send()
        e.Ack()
        return  nil
    }

    sub, err := broker.Subscribe("ingestion-in-dev", myHandler, broker.Queue("ingestion-in-dev"));
    if err != nil {
        log.Error().Err(err).Send()
    }
    sub.Unsubscribe()
}

