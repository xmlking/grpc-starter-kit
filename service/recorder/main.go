package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	"github.com/xmlking/grpc-starter-kit/internal/version"
	"github.com/xmlking/grpc-starter-kit/service/recorder/registry"
	"github.com/xmlking/grpc-starter-kit/service/recorder/subscriber"
	"github.com/xmlking/toolkit/broker/cloudevents"
	_ "github.com/xmlking/toolkit/logger/auto"
	"github.com/xmlking/toolkit/server"
	"github.com/xmlking/toolkit/util/endpoint"
	"golang.org/x/sync/errgroup"
)

func main() {
	serviceName := constants.RECORDER_SERVICE
	cfg := config.GetConfig()

	appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer stop()

	g, ctx := errgroup.WithContext(appCtx)

	broker.DefaultBroker = broker.NewBroker(ctx, broker.Name(serviceName))

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}
	transactionSubscriber := ctn.Resolve("transaction-subscriber").(*subscriber.TransactionSubscriber)

	if err := broker.AddSubscriber(cfg.Services.Recorder.Endpoint, transactionSubscriber.HandleSend); err != nil {
		log.Fatal().Err(err).Msgf("Failed subscribing to Topi %s", cfg.Services.Recorder.Endpoint)
	}

	listener, err := endpoint.GetListener("tcp:///:0" /*cfg.Services.Recorder.Endpoint*/)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("error creating listener")
	}
	srv := server.NewServer(ctx, server.ServerName(serviceName), server.WithListener(listener))

	// Start broker/gRPC daemon services
	log.Info().Object("build_info", version.GetBuildInfo()).Send()
	log.Info().Msgf("Server(%s) starting at: %s, secure: %t, pid: %d", serviceName, listener.Addr(), cfg.Features.TLS.Enabled, os.Getpid())

	g.Go(func() error {
		return broker.Start()
	})

	g.Go(func() error {
		return srv.Start()
	})

	go func() {
		if err := g.Wait(); err != nil {
			log.Fatal().Stack().Err(err).Msgf("Unexpected error for service: %s", cfg.Services.Emailer.Endpoint)
		}
		log.Info().Msg("Goodbye.....")
		os.Exit(0)
	}()

	// Listen for the interrupt signal.
	<-appCtx.Done()

	// notify user of shutdown
	switch ctx.Err() {
	case context.DeadlineExceeded:
		log.Info().Str("cause", "timeout").Msg("Shutting down gracefully, press Ctrl+C again to force")
	case context.Canceled:
		log.Info().Str("cause", "interrupt").Msg("Shutting down gracefully, press Ctrl+C again to force")
	}

	// Restore default behavior on the interrupt signal.
	stop()

	// Perform application shutdown with a maximum timeout of 1 minute.
	timeoutCtx, cancel := context.WithTimeout(context.Background(), constants.DefaultShutdownTimeout)
	defer cancel()

	// force termination after shutdown timeout
	<-timeoutCtx.Done()
	log.Error().Msg("Shutdown grace period elapsed. force exit")
	// force stop any daemon services here:
	srv.Stop()
	os.Exit(1)
}
