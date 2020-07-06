package append

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/xmlking/grpc-starter-kit/shared/constants"
)

// UnaryServerInterceptor appends metadata to outgoing
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	options := defaultOptions()
	for _, o := range opts {
		o(&options)
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md := options.getMetadata()
		var pairs []string
		for k, vv := range md {
			pairs = append(pairs, k, vv[0])
		}

		if options.traceIdEnabled {
			pairs = append(pairs, constants.TraceIDKey, uuid.New().String())
		}

		if len(pairs) > 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, pairs...)
		}

		return handler(ctx, req)
	}
}

// UnaryClientInterceptor appends metadata to outgoing
func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	options := defaultOptions()
	for _, o := range opts {
		o(&options)
	}

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md := options.getMetadata()
		println(len(md))
		var pairs []string
		for k, vv := range md {
			pairs = append(pairs, k, vv[0])
		}
		println(len(pairs))

		if options.traceIdEnabled {
			pairs = append(pairs, constants.TraceIDKey, uuid.New().String())
		}

		if len(pairs) > 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, pairs...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
