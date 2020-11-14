package main

import (
	"github.com/rs/zerolog/log"

	"github.com/xmlking/toolkit/broker"
	"github.com/xmlking/toolkit/service"

	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	_ "github.com/xmlking/grpc-starter-kit/internal/constants"
	_ "github.com/xmlking/grpc-starter-kit/internal/logger"
	"github.com/xmlking/grpc-starter-kit/service/emailer/registry"
	"github.com/xmlking/grpc-starter-kit/service/emailer/subscriber"
)

func main() {
	serviceName := constants.EMAILER_SERVICE
	cfg := config.GetConfig()

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}
	emailSubscriber := ctn.Resolve("emailer-subscriber").(*subscriber.EmailSubscriber)

	srv := service.NewService(
		service.Name(serviceName),
		service.Version(cfg.Services.Emailer.Version),
		service.WithBrokerOptions(
			broker.Name("mkit.broker.emailer"),
			broker.WithEndpoint(cfg.Services.Emailer.Endpoint),
		),
		// service.WithBrokerOptions(...),
	)
	srv.AddSubscriber(emailSubscriber.HandleSend)

	// Start server!
	log.Info().Msg(config.GetBuildInfo())
	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
