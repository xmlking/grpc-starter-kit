package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/sercand/kuberesolver"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/instrumentation/grpctrace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/rs/zerolog/log"
	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/service"
	"github.com/xmlking/toolkit/util/tls"

	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/play/handler"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
	"github.com/xmlking/grpc-starter-kit/shared/telemetry/metrics"
	"github.com/xmlking/grpc-starter-kit/shared/telemetry/tracing"
)

func main() {
	serviceName := constants.PLAY_SERVICE
	cfg := config.GetConfig()

	// Register kuberesolver to grpc
	if config.IsProduction() {
		kuberesolver.RegisterInCluster()
	}

	if cfg.Features.Tracing.Enabled {
		closeFn := tracing.InitTracing(cfg.Features.Tracing)
		defer closeFn()
	}

	if cfg.Features.Metrics.Enabled {
		exporter := metrics.InitMetrics(cfg.Features.Metrics)
		defer exporter.Stop()
	}

	var unaryInterceptors = []grpc.UnaryServerInterceptor{grpc_validator.UnaryServerInterceptor()}
	var streamInterceptors = []grpc.StreamServerInterceptor{grpc_validator.StreamServerInterceptor()}

	if cfg.Features.Tracing.Enabled {
		unaryInterceptors = append(unaryInterceptors, grpctrace.UnaryServerInterceptor(global.Tracer("")))
		streamInterceptors = append(streamInterceptors, grpctrace.StreamServerInterceptor(global.Tracer("")))
	}
	if cfg.Features.Rpclog.Enabled {
		// keep it last in the interceptor chain
		unaryInterceptors = append(unaryInterceptors, rpclog.UnaryServerInterceptor())
	}

	grpcOps := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
	}

	if cfg.Features.Tls.Enabled {
		tlsConf, err := tls.NewTLSConfig(cfg.Features.Tls.CertFile, cfg.Features.Tls.KeyFile, cfg.Features.Tls.CaFile, cfg.Features.Tls.ServerName, cfg.Features.Tls.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cert")
		}
		serverCert := credentials.NewTLS(tlsConf)
		grpcOps = append(grpcOps, grpc.Creds(serverCert))
	}

	srv := service.NewService(
		service.Name(serviceName),
		service.Version(cfg.Services.Play.Version),
		service.WithGrpcEndpoint(cfg.Services.Play.Endpoint),
		service.WithGrpcOptions(grpcOps...),
	)
	// create a gRPC server object
	grpcServer := srv.Server()

	// create a server instance
	greeterHandler := handler.NewGreeterHandler()

	// attach the Greeter service to the server
	greeterv1.RegisterGreeterServiceServer(grpcServer, greeterHandler)

	// start the server
	log.Info().Msg(config.GetBuildInfo())
	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
