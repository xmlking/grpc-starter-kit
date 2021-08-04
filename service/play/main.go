package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/rs/zerolog/log"
	"github.com/sercand/kuberesolver"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	_ "github.com/xmlking/toolkit/logger/auto"
	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/server"
	"github.com/xmlking/toolkit/telemetry/metrics"
	"github.com/xmlking/toolkit/telemetry/tracing"
	"github.com/xmlking/toolkit/util/endpoint"
	"github.com/xmlking/toolkit/util/tls"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/play/handler"
)

func main() {
	serviceName := constants.PLAY_SERVICE
	cfg := config.GetConfig()
	efs := config.GetFileSystem()

	appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer stop()

	g, ctx := errgroup.WithContext(appCtx)

	// Register kuberesolver to grpc.
	// This line should be before calling registry.NewContainer(cfg)
	if config.IsProduction() {
		kuberesolver.RegisterInCluster()
	}

	if cfg.Features.Tracing.Enabled {
		closeFn := tracing.InitTracing(ctx, cfg.Features.Tracing)
		defer closeFn()
	}

	if cfg.Features.Metrics.Enabled {
		closeFn := metrics.InitMetrics(ctx, cfg.Features.Metrics)
		defer closeFn()
	}

	var unaryInterceptors = []grpc.UnaryServerInterceptor{grpc_validator.UnaryServerInterceptor()}
	var streamInterceptors = []grpc.StreamServerInterceptor{grpc_validator.StreamServerInterceptor()}

	if cfg.Features.Tracing.Enabled {
		unaryInterceptors = append(unaryInterceptors, otelgrpc.UnaryServerInterceptor())
		streamInterceptors = append(streamInterceptors, otelgrpc.StreamServerInterceptor())
	}
	if cfg.Features.Rpclog.Enabled {
		// keep it last in the interceptor chain
		unaryInterceptors = append(unaryInterceptors, rpclog.UnaryServerInterceptor())
	}

	// ServerOption
	grpcOps := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
	}

	if cfg.Features.TLS.Enabled {
		tlsConf, err := tls.NewTLSConfig(efs, cfg.Features.TLS.CertFile, cfg.Features.TLS.KeyFile, cfg.Features.TLS.CaFile, cfg.Features.TLS.ServerName, cfg.Features.TLS.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cert")
		}
		serverCert := credentials.NewTLS(tlsConf)
		grpcOps = append(grpcOps, grpc.Creds(serverCert))
	}

	listener, err := endpoint.GetListener(cfg.Services.Play.Endpoint)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("error creating listener")
	}
	srv := server.NewServer(appCtx, server.ServerName(serviceName), server.WithListener(listener), server.WithServerOptions(grpcOps...))

	gSrv := srv.Server()

	greeterHandler := handler.NewGreeterHandler()
	// attach the Greeter service to the server
	greeterv1.RegisterGreeterServiceServer(gSrv, greeterHandler)

	// Start broker/gRPC daemon services
	log.Info().Msg(config.GetBuildInfo())
	log.Info().Msgf("Server(%s) starting at: %s, secure: %t, pid: %d", serviceName, listener.Addr(), cfg.Features.TLS.Enabled, os.Getpid())

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
