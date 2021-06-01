package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"

    "github.com/xmlking/toolkit/grpc"
    _ "github.com/xmlking/toolkit/logger/auto"
    "golang.org/x/sync/errgroup"

    "github.com/rs/zerolog/log"
    "github.com/xmlking/grpc-starter-kit/internal/config"
    "github.com/xmlking/grpc-starter-kit/internal/constants"
    "github.com/xmlking/grpc-starter-kit/service/emailer/registry"
    "github.com/xmlking/grpc-starter-kit/service/emailer/subscriber"
    "github.com/xmlking/toolkit/broker/cloudevents"
)

func main() {
    // println(os.Getpid())
    appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
    defer stop()

    g, ctx := errgroup.WithContext(appCtx)

    broker.DefaultBroker = broker.NewBroker(ctx)

	cfg := config.GetConfig()

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}
	emailSubscriber := ctn.Resolve("emailer-subscriber").(*subscriber.EmailSubscriber)

    if err := broker.AddSubscriber(cfg.Services.Emailer.Endpoint,  emailSubscriber.HandleSend); err != nil {
        log.Fatal().Err(err).Msgf("Failed subscribing to Topi %s", cfg.Services.Emailer.Endpoint)
    }

    // gSrv := grpc.NewServer(appCtx, grpc.WithGrpcEndpoint(cfg.Services.Emailer.Endpoint));
    gSrv := grpc.NewServer(appCtx);

    g.Go(func() error {
        return broker.Start()
    })

    g.Go(func() error {
        return gSrv.Start()
    })


	// Start server! constants.EMAILER_SERVICE
	log.Info().Msg(config.GetBuildInfo())
    go func() {
        if err := g.Wait(); err != nil {
            log.Fatal().Stack().Err(err).Msgf("Unexpected error for service: %s",  cfg.Services.Emailer.Endpoint)
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
    gSrv.Stop()
    os.Exit(1)
}
