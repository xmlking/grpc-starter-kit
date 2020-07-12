package main

import (
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/rs/zerolog/log"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/xmlking/toolkit/middleware/rpclog"
	"github.com/xmlking/toolkit/service"
	"github.com/xmlking/toolkit/util/endpoint"
	"github.com/xmlking/toolkit/util/tls"

	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/greeter/handler"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func main() {
	serviceName := constants.GREETER_SERVICE
	cfg := config.GetConfig()

	grpcOps := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			// keep it last in the interceptor chain
			rpclog.UnaryServerInterceptor(),
		)),
	}

	if cfg.Features.Tls.Enabled {
		tlsConf, err := tls.NewTLSConfig(cfg.Features.Tls.CertFile, cfg.Features.Tls.KeyFile, cfg.Features.Tls.CaFile, cfg.Features.Tls.ServerName, cfg.Features.Tls.Password)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cert")
		}
		serverCert := credentials.NewTLS(tlsConf)
		// serverCert, err := credentials.NewServerTLSFromFile(cfg.Features.Tls.CertFile, cfg.Features.Tls.KeyFile)
		//if err != nil {
		//    log.Fatal().Err(err).Msg("failed to create cert")
		//}
		grpcOps = append(grpcOps, grpc.Creds(serverCert))
	}

	srv := service.NewService(
		service.Name(serviceName),
		service.Version(cfg.Services.Greeter.Version),
		service.WithGrpcEndpoint(cfg.Services.Greeter.Endpoint),
		service.WithGrpcOptions(grpcOps...),
		// service.WithBrokerOptions(...),
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
		Handler: handler.NewHttpHandler(),
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
	log.Info().Msgf("Server (%s) started at: %s, secure: %t", serviceName, lis.Addr(), cfg.Features.Tls.Enabled)
	mux.Serve()
}
