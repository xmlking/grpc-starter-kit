package main

import (
	"context"
	"os"
	"os/signal"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/shared/broker"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

var (
	cfg = config.GetConfig()
)

func main() {
	broker.DefaultBroker = broker.NewBroker()

	myHandler := func(ctx context.Context, msg *pubsub.Message) error {
		//md, _ := metadata.FromContext(ctx)
		//log.Info().Interface("md", md).Send()
		log.Info().Interface("event.Message.ID", msg.ID).Send()
		log.Info().Interface("event.Message.Attributes", msg.Attributes).Send()
		log.Info().Interface("event.Message.Data", msg.Data).Send()

		log.Info().Interface("event.Message", msg).Send()
		msg.Ack() // or msg.Nack() // or return error for autoAck
		return nil
	}

	err := broker.Subscribe("ingestion-in-dev", myHandler, broker.Queue("ingestion-in-dev"))
	if err != nil {
		log.Error().Err(err).Msg("Failed subscribing to Topic: ingestion-in-dev")
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Info().Msg("Got to Go...")
	// close all subs and then connection.
	if err := broker.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("Unexpected disconnect error")
	}
}
