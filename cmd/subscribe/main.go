package main

import (
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/shared/broker"
	"github.com/xmlking/grpc-starter-kit/shared/config"
)

var (
	cfg = config.GetConfig()
)

func main() {
	broker.DefaultBroker = broker.NewBroker()

	myHandler := func(e broker.Event) error {
		msg := e.Message()
		log.Info().Interface("event.Message.Body", msg.Body).Send()
		log.Info().Interface("event.Message.Header", msg.Header).Send()
		log.Info().Interface("event.Message", msg).Send()
		log.Info().Interface("event.Topic", e.Topic()).Send()
		_ = e.Ack()
		return nil
	}

	var subs []broker.Subscriber

	// like adding micro.NewSubscribe(...)
	sub, err := broker.Subscribe("ingestion-in-dev", myHandler, broker.Queue("ingestion-in-dev"))
	if err != nil {
		log.Error().Err(err).Msg("Subscribe to `ingestion-in-dev` Topic failed")
	}
	subs = append(subs, sub)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Info().Msg("Got to Go...")
	//_ = sub.Unsubscribe()
	//_ = broker.Disconnect()

	// close all subs and then connection.
	for _, sub := range subs {
		log.Info().Msgf("Unsubscribing from topic: %s", sub.Topic())
		sub.Unsubscribe()
	}

	if err := broker.Disconnect(); err != nil {
		log.Fatal().Err(err).Msg("Unexpected disconnect error")
	}
}
