package main

import (
	"github.com/rs/zerolog/log"

	"github.com/xmlking/toolkit/broker/cloudevents"
	"github.com/xmlking/toolkit/service"

	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	_ "github.com/xmlking/grpc-starter-kit/internal/logger"
	"github.com/xmlking/grpc-starter-kit/service/recorder/registry"
	"github.com/xmlking/grpc-starter-kit/service/recorder/subscriber"
)

func main() {
	serviceName := constants.RECORDER_SERVICE
	cfg := config.GetConfig()

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}
	transactionSubscriber := ctn.Resolve("transaction-subscriber").(*subscriber.TransactionSubscriber)

	srv := service.NewService(
		service.Name(serviceName),
		service.Version(cfg.Services.Recorder.Version),
		// service.WithBrokerOptions(...),
	)
	bkr := broker.NewBroker(broker.Name(serviceName))
	_, _ = bkr.NewSubscriber(cfg.Services.Recorder.Endpoint, transactionSubscriber.HandleSend)
	err = bkr.Start()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to start the Broker: %s", cfg.Services.Recorder.Endpoint)
	}

	// Start server!
	log.Info().Msg(config.GetBuildInfo())
	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
