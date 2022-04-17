package main

import (
    "context"
    "os"
    "os/signal"
    "syscall"
    "time"

    _ "github.com/xiaoqidun/entps"
    // required by schema hooks.
    _ "github.com/xmlking/grpc-starter-kit/ent/runtime"
    _ "github.com/xmlking/toolkit/logger/auto"

    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
    grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
    "github.com/rs/zerolog/log"
    "github.com/sercand/kuberesolver"
    profilev1 "github.com/xmlking/grpc-starter-kit/gen/go/gkit/service/account/profile/v1"
    userv1 "github.com/xmlking/grpc-starter-kit/gen/go/gkit/service/account/user/v1"
    "github.com/xmlking/grpc-starter-kit/internal/config"
    "github.com/xmlking/grpc-starter-kit/internal/constants"
    "github.com/xmlking/grpc-starter-kit/internal/middleware/translog"
    "github.com/xmlking/grpc-starter-kit/internal/version"
    "github.com/xmlking/grpc-starter-kit/service/account/registry"
    broker "github.com/xmlking/toolkit/broker/cloudevents"
    "github.com/xmlking/toolkit/middleware/rpclog"
    appendTags "github.com/xmlking/toolkit/middleware/tags/append"
    forwardTags "github.com/xmlking/toolkit/middleware/tags/forward"
    "github.com/xmlking/toolkit/server"
    "github.com/xmlking/toolkit/util/endpoint"
    "github.com/xmlking/toolkit/util/tls"
    "golang.org/x/sync/errgroup"
    "google.golang.org/grpc"
    "google.golang.org/grpc/backoff"
    "google.golang.org/grpc/credentials"
)

func main() {
    serviceName := constants.ACCOUNT_SERVICE
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

    // Initialize DI Container
    ctn, err := registry.NewContainer(ctx, cfg)
    defer ctn.Clean()
    if err != nil {
        log.Fatal().Msgf("failed to build container: %v", err)
    }
    translogPublisher := ctn.Resolve("translog-publisher").(broker.Publisher)
    // Handlers
    userHandler := ctn.Resolve("user-handler").(userv1.UserServiceServer)
    profileHandler := ctn.Resolve("profile-handler").(profilev1.ProfileServiceServer)

    // ServerOption
    grpcOps := []grpc.ServerOption{
        grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
            // Execution is done in left-to-right order
            // keep around type Interceptors first,
            rpclog.UnaryServerInterceptor(),
            grpc_validator.UnaryServerInterceptor(),
            appendTags.UnaryServerInterceptor(appendTags.WithPairs(constants.FromServiceKey, constants.ACCOUNT_SERVICE)),
            forwardTags.UnaryServerInterceptor(forwardTags.WithForwardTags(constants.TraceIDKey, constants.TenantIDKey)),
            translog.UnaryServerInterceptor(translogPublisher, serviceName),
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

    // DialOptions
    // var dialOptions []grpc.DialOption
    dialOptions := []grpc.DialOption{
        grpc.WithAuthority(cfg.Services.Greeter.Authority),
        grpc.WithConnectParams(grpc.ConnectParams{Backoff: backoff.Config{BaseDelay: 5 * time.Second}, MinConnectTimeout: 5 * time.Second}),
    }
    var ucInterceptors []grpc.UnaryClientInterceptor

    tlsConf := cfg.Features.TLS
    if tlsConf.Enabled {
        if creds, err := tls.NewTLSConfig(efs, tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.ServerName, cfg.Features.TLS.Password); err != nil {
            log.Fatal().Err(err).Msg("Failed to create tlsConf")
        } else {
            dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
        }
    } else {
        dialOptions = append(dialOptions, grpc.WithInsecure())
    }

    if cfg.Features.Rpclog.Enabled {
        ucInterceptors = append(ucInterceptors, rpclog.UnaryClientInterceptor())
    }

    if len(ucInterceptors) > 0 {
        dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(ucInterceptors...)))
    }

    listener, err := endpoint.GetListener(cfg.Services.Account.Endpoint)
    if err != nil {
        log.Fatal().Stack().Err(err).Msg("error creating listener")
    }
    srv := server.NewServer(ctx, server.ServerName(serviceName), server.WithListener(listener), server.WithServerOptions(grpcOps...))

    // greeterClientCon, err := srv.Client(cfg.Services.Greeter.Endpoint, server.ClientName(constants.GREETER_SERVICE), server.WithDialOptions(dialOptions...))

    gSrv := srv.Server()
    // Register Handlers
    userv1.RegisterUserServiceServer(gSrv, userHandler)
    profilev1.RegisterProfileServiceServer(gSrv, profileHandler)

    // Start broker/gRPC daemon services
    log.Info().Object("build_info", version.GetBuildInfo()).Send()
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
