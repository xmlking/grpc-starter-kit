package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/rs/zerolog/log"
	"github.com/soheilhy/cmux"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/greeter/handler"
	_ "github.com/xmlking/toolkit/logger/auto"
	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/server"
	"github.com/xmlking/toolkit/util/endpoint"
	"github.com/xmlking/toolkit/util/tls"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	serviceName := constants.GREETER_SERVICE
	cfg := config.GetConfig()
	efs := config.GetFileSystem()

	appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer stop()

	g, ctx := errgroup.WithContext(appCtx)

	// ServerOption
	grpcOps := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			// keep it last in the interceptor chain
			rpclog.UnaryServerInterceptor(),
		)),
	}

	if cfg.Features.TLS.Enabled {
		tlsConf, err := tls.NewTLSConfig(efs, cfg.Features.TLS.CertFile, cfg.Features.TLS.KeyFile, cfg.Features.TLS.CaFile, cfg.Features.TLS.ServerName, cfg.Features.TLS.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cert")
		}
		serverCert := credentials.NewTLS(tlsConf)
		grpcOps = append(grpcOps, grpc.Creds(serverCert))
	}

	listener, err := endpoint.GetListener(cfg.Services.Greeter.Endpoint)
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

// cmux example
func main_cmux() {
	serviceName := constants.GREETER_SERVICE
	cfg := config.GetConfig()

	lis, err := endpoint.GetListener(cfg.Services.Greeter.Endpoint)
	if err != nil {
		log.Fatal().Msgf("failed to create listener: %v", err)
	}

	// Create a cmux.
	mux := cmux.New(lis)
	// Match connections in order:
	grpcL := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := mux.Match(cmux.HTTP1Fast())

	// Create your protocol servers.
	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			// keep it last in the interceptor chain
			rpclog.UnaryServerInterceptor(),
		)))
	greeterv1.RegisterGreeterServiceServer(grpcS, handler.NewGreeterHandler())

	// Register http Handlers
	httpS := &http.Server{
		Handler: handler.NewHTTPHandler(),
	}

	// Add HealthChecks
	hsrv := health.NewServer()
	for name := range grpcS.GetServiceInfo() {
		hsrv.SetServingStatus(name, grpc_health_v1.HealthCheckResponse_SERVING)
	}
	grpc_health_v1.RegisterHealthServer(grpcS, hsrv)
	// TODO: User our own custom health implementation, instead of using built-in health server
	// https://github.com/GoogleCloudPlatform/grpc-gke-nlb-tutorial/blob/master/echo-grpc/health/health.go

	// Use the muxed listeners for your servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	// Start server!
	reflection.Register(grpcS)
	println(config.GetBuildInfo())
	log.Info().Msgf("Server (%s) started at: %s, secure: %t", serviceName, lis.Addr(), cfg.Features.TLS.Enabled)
	mux.Serve()
}
