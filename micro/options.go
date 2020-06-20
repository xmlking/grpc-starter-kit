package micro

import (
	"context"

	"google.golang.org/grpc"
)

// Options for micro service
type Options struct {
	Config Interface
	Client *grpc.ClientConn
	Server *grpc.Server

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	Signal bool
}
