package main

import (
	_ "github.com/xmlking/toolkit/logger/auto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/rs/zerolog/log"
	"github.com/sercand/kuberesolver"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/service"
	"github.com/xmlking/toolkit/util/tls"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//"github.com/xmlking/grpc-starter-kit/internal/telemetry/metrics"
	"github.com/xmlking/grpc-starter-kit/internal/telemetry/tracing"
	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/play/handler"
)

func main() {
	serviceName := constants.PLAY_SERVICE
	cfg := config.GetConfig()
	efs := config.GetFileSystem()

	// Register kuberesolver to grpc
	if config.IsProduction() {
		kuberesolver.RegisterInCluster()
	}

	if cfg.Features.Tracing.Enabled {
		closeFn := tracing.InitTracing(cfg.Features.Tracing)
		defer closeFn()
	}

	//if cfg.Features.Metrics.Enabled {
	//	exporter := metrics.InitMetrics(cfg.Features.Metrics)
	//	defer exporter.Stop()
	//}

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
