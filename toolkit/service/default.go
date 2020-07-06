package service

import (
    "context"
    "net"
    "time"

    "github.com/pkg/errors"
    "github.com/rs/zerolog/log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/health"
    "google.golang.org/grpc/health/grpc_health_v1"
    "google.golang.org/grpc/reflection"

    "github.com/xmlking/grpc-starter-kit/shared/config"
    configPB "github.com/xmlking/grpc-starter-kit/shared/proto/config/v1"
    "github.com/xmlking/grpc-starter-kit/toolkit/broker"
    "github.com/xmlking/grpc-starter-kit/toolkit/util/endpoint"
    "github.com/xmlking/grpc-starter-kit/toolkit/util/signals"
)

const (
    DefaultName    = "mkit.service.default"
    DefaultVersion = "latest"
    DefaultAddress = ":0"
    // DefaultShutdownTimeout defines the default timeout given to the service when calling Shutdown.
    DefaultShutdownTimeout = time.Minute * 1
)

type service struct {
    subscribers []interface{}
    opts        Options
    cfg         configPB.Configuration
    grpcServer  *grpc.Server
    broker      broker.Broker
}

func newService(opts ...Option) Service {
    // Default Options
    options := Options{
        Name:            DefaultName,
        Version:         DefaultVersion,
        ShutdownTimeout: DefaultShutdownTimeout,
        // Set up signals so we handle the first shutdown signal gracefully.
        Context: signals.NewContext(),
    }
    s := service{opts: options}
    s.cfg = config.GetConfig()
    s.ApplyOptions(opts...)

    s.grpcServer = grpc.NewServer(s.opts.GrpcOptions...)

    return &s
}

func (s *service) ApplyOptions(opts ...Option) {
    // process options
    for _, o := range opts {
        o(&s.opts)
    }
}

func (s *service) AddSubscriber(fn interface{}) {
    s.subscribers = append(s.subscribers, fn)
}

func (s *service) Options() Options {
    return s.opts
}

func (s *service) Server() *grpc.Server {
    return s.grpcServer
}

func (s *service) Shutdown() error {
    return nil
}

func (s *service) Start() (err error) {
    //println(config.GetBuildInfo())

    // eg, egCtx := errgroup.WithContext(s.opts.Context)
    ctx, cancel := context.WithCancel(s.opts.Context)
    defer cancel()

    errCh := make(chan error, 1)

    // Start
    log.Info().Msg(config.GetBuildInfo())

    // Start GrpcServer
    // Add HealthChecks
    hsrv := health.NewServer()
    for name := range s.grpcServer.GetServiceInfo() {
        hsrv.SetServingStatus(name, grpc_health_v1.HealthCheckResponse_SERVING)
    }
    grpc_health_v1.RegisterHealthServer(s.grpcServer, hsrv)
    // TODO: User our own custom health implementation, instead of using built-in health server
    // https://github.com/GoogleCloudPlatform/grpc-gke-nlb-tutorial/blob/master/echo-grpc/health/health.go

    var listener net.Listener
    if s.opts.GrpcEndpoint == "" {
        listener, err = net.Listen("tcp", DefaultAddress)
        if err != nil {
            return errors.Wrap(err, "Failed to create listener")
        }
    } else {
        listener, err = endpoint.GetListener(s.opts.GrpcEndpoint)
        if err != nil {
            return errors.Wrap(err, "Failed to create listener")
        }
    }
    log.Info().Msgf("Server (%s) starting at: %s, secure: %t", s.opts.Name, listener.Addr(), s.cfg.Features.Tls.Enabled)
    go func() {
        reflection.Register(s.grpcServer)
        errCh <- s.grpcServer.Serve(listener)
    }()

    // Start Broker
    if len(s.subscribers) > 0 {
        s.opts.BrokerOptions = append(s.opts.BrokerOptions, broker.Context(ctx))
        s.broker = broker.NewBroker(s.opts.BrokerOptions...)

        log.Info().Msgf("Broker (%s) starting at: %s, secure: %t", s.broker.Options().Name, s.broker.Options().Endpoint, s.cfg.Features.Tls.Enabled)
        for _, receiver := range s.subscribers {
            // eg.Go()
            go func(ctx context.Context) {
                errCh <- s.broker.CeClient().StartReceiver(ctx, receiver)
            }(ctx)
        }
    }

    // This will block until either a signal arrives or one of the grouped functions
    // returns an error.
    // <-egCtx.Done()

    // Stop either if the receiver stops (sending to errCh) or if stopCh is closed.
    select {
    case err := <-errCh:
        return err
    case <-ctx.Done():
        break
    }

    // do any more resources closing here
    s.grpcServer.GracefulStop()
    cancel()

    select {
    case <-errCh:
        s.grpcServer.Stop()
        log.Debug().Msg("Gracefully shutdown")
    case <-time.After(DefaultShutdownTimeout):
        log.Error().Msg("Failed to shutdown within grace period")
        return errors.New("Failed to shutdown within grace period")
    }

    return nil
}
