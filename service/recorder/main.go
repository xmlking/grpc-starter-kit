package main

import (
	"github.com/rs/zerolog/log"

	"github.com/xmlking/toolkit/broker"
	"github.com/xmlking/toolkit/service"

	"github.com/xmlking/grpc-starter-kit/service/recorder/registry"
	"github.com/xmlking/grpc-starter-kit/service/recorder/subscriber"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
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
		service.WithBrokerOptions(
			broker.Name("mkit.broker.recorder"),
			broker.WithEndpoint(cfg.Services.Recorder.Endpoint),
		),
		// service.WithBrokerOptions(...),
	)
	srv.AddSubscriber(transactionSubscriber.HandleSend)

	// Start server!
	log.Info().Msg(config.GetBuildInfo())
	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
