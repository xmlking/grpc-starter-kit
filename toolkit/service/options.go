package service

import (
    "context"
    "time"

    "google.golang.org/grpc"

    "github.com/xmlking/grpc-starter-kit/toolkit/broker"
)

type Option func(*Options)

type Options struct {
    Name    string
    Version string

    GrpcEndpoint string
    GrpcOptions []grpc.ServerOption

    BrokerOptions []broker.Option

    // ShutdownTimeout defines the timeout given to the http.Server when calling Shutdown.
    // If nil, DefaultShutdownTimeout is used.
    ShutdownTimeout time.Duration
    // Alternative options
    Context context.Context

    // Before and After funcs
    //BeforeStart []func() error
    //BeforeStop  []func() error
    //AfterStart  []func() error
    //AfterStop   []func() error

}

// Name of the service
func Name(n string) Option {
    return func(o *Options) {
        o.Name = n
    }
}

// Version of the service
func Version(v string) Option {
    return func(o *Options) {
        o.Version = v
    }
}

// Context specifies a context for the service.
// Can be used to signal shutdown of the service
// Can be used for extra option values.
func Context(ctx context.Context) Option {
    return func(o *Options) {
        o.Context = ctx
    }
}

func WithGrpcOptions(opts ...grpc.ServerOption) Option {
    return func(o *Options) {
        o.GrpcOptions = opts
    }
}

// WithGrpcEndpoint specifies the net.Listener endpoint to use instead of the default
func WithGrpcEndpoint(endpoint string) Option {
    return func(o *Options) {
        o.GrpcEndpoint = endpoint
    }
}

func WithBrokerOptions(opts ...broker.Option) Option {
    return func(o *Options) {
        o.BrokerOptions = opts
    }
}

func WithShutdownTimeout(timeout time.Duration) Option {
    return func(o *Options) {
        o.ShutdownTimeout = timeout
    }
}

