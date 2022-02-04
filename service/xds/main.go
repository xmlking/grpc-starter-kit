package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	"github.com/xmlking/grpc-starter-kit/internal/version"
	_ "github.com/xmlking/toolkit/logger/auto"
	"github.com/xmlking/toolkit/server"
	"github.com/xmlking/toolkit/util/endpoint"
	"github.com/xmlking/toolkit/util/tls"
	"github.com/xmlking/toolkit/xds"
	"github.com/xmlking/toolkit/xds/callbacks"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const (
	grpcKeepaliveTime        = 30 * time.Second
	grpcKeepaliveTimeout     = 5 * time.Second
	grpcKeepaliveMinTime     = 30 * time.Second
	grpcMaxConcurrentStreams = 1000000
)

func main() {
	serviceName := constants.XDS_SERVICE
	cfg := config.GetConfig()
	tlsConfig := cfg.Features.TLS
	xdsConfig := cfg.Xds
	efs := config.GetFileSystem()

	appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer stop()

	// we use this context to cancel all sibling goroutines in the errgroup, when one of the sibling goroutine throw error.
	// it also bubble down signals from parent appCtx to all child goroutines tree.
	g, ctx := errgroup.WithContext(appCtx)

	// gRPC golang library sets a very small upper bound for the number gRPC/h2
	// streams over a single TCP connection. If a proxy multiplexes requests over
	// a single connection to the management server, then it might lead to
	// availability problems. Keepalive timeouts based on connection_keepalive parameter https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/examples#dynamic
	grpcOps := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(xdsConfig.MaxConcurrentStreams),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    grpcKeepaliveTime,
			Timeout: grpcKeepaliveTimeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             grpcKeepaliveMinTime,
			PermitWithoutStream: true,
		}),
	}

	if cfg.Features.TLS.Enabled {
		tlsConf, err := tls.NewTLSConfig(efs, tlsConfig.CertFile, tlsConfig.KeyFile, tlsConfig.CaFile, tlsConfig.ServerName, tlsConfig.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cert")
		}
		serverCert := credentials.NewTLS(tlsConf)
		grpcOps = append(grpcOps, grpc.Creds(serverCert))
	}

	listener, err := endpoint.GetListener(cfg.Services.Xds.Endpoint)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("error creating listener")
	}
	srv := server.NewServer(ctx, server.ServerName(serviceName), server.WithListener(listener), server.WithServerOptions(grpcOps...))

	gSrv := srv.Server()
	refresh := xds.NewRefresher(ctx, xdsConfig.SourceType, xds.WithNodeID(xdsConfig.NodeID), xds.WithRefreshInterval(xdsConfig.RefreshInterval),
		xds.WithFS(efs),
		xds.WithHostnames(xdsConfig.DNS.Hostnames),
		xds.WithNamespace(xdsConfig.Kubernetes.Namespace),
	)

	var cb serverv3.Callbacks
	if cfg.Features.Metrics.Enabled {
		cb = callbacks.NewOTelCallbacks()
	} else {
		cb = callbacks.NewDefaultCallbacks()
	}

	adsSrv := serverv3.NewServer(ctx, refresh.GetCache(), cb)
	// register services
	discoverygrpc.RegisterAggregatedDiscoveryServiceServer(gSrv, adsSrv)

	// Start broker/gRPC daemon services
	log.Info().Object("build_info", version.GetBuildInfo()).Send()
	log.Info().Msgf("Server(%s) starting at: %s, secure: %t, pid: %d", serviceName, listener.Addr(), cfg.Features.TLS.Enabled, os.Getpid())

	g.Go(func() error {
		return refresh.Start()
	})
	g.Go(func() error {
		return srv.Start()
	})

	go func() {
		if err := g.Wait(); err != nil {
			log.Fatal().Stack().Err(err).Msgf("Unexpected error for service: %s", cfg.Services.Xds.Endpoint)
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
