package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/service/recorder/registry"
	"github.com/xmlking/grpc-starter-kit/service/recorder/subscriber"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
	"github.com/xmlking/grpc-starter-kit/shared/eventing"
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

	ceClient := eventing.NewSinkClient(cfg.Services.Recorder.Endpoint)

	// Start server!
	println(config.GetBuildInfo())
	log.Info().Msgf("Server (%s) started at: %s, secure: %t", serviceName, cfg.Services.Recorder.Endpoint, cfg.Features.Tls.Enabled)
	if err := ceClient.StartReceiver(context.Background(), transactionSubscriber.HandleSend); err != nil {
		log.Fatal().Err(err).Send()
	}
}
