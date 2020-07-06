package broker

import (
    "context"
)

type Option func(*Options)

type Options struct {
    Name    string

    Endpoint string

    // Alternative options
    Context context.Context
}

// Name of the service
func Name(n string) Option {
    return func(o *Options) {
        o.Name = n
    }
}

func WithEndpoint(endpoint string) Option {
    return func(o *Options) {
        o.Endpoint = endpoint
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


