package rpclog

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor is an example server-side request logger middleware
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	options := defaultOptions()
	for _, o := range opts {
		o(&options)
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// exclude methods
		if options.matchMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		incomingMd := metautils.ExtractIncoming(ctx)
		resp, err := handler(ctx, req)
		outgoingMd := metautils.ExtractOutgoing(ctx)

		log.Debug().
			Str("module", "rpclog").
			Str("method", info.FullMethod).
			Interface("incoming_md:", incomingMd).
			Interface("req", req).
			Interface("outgoing_md:", outgoingMd).
			Interface("resp", resp).
			Interface("err", err).
			Msg("Server-Side")

		return resp, err
	}
}

// UnaryClientInterceptor is an example client-side request logger middleware
func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// exclude methods
		if options.matchMethod(method) {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		incomingMd := metautils.ExtractIncoming(ctx)
		err := invoker(ctx, method, req, reply, cc, opts...)
		outgoingMd := metautils.ExtractOutgoing(ctx)

		log.Debug().
			Str("module", "rpclog").
			Str("method", method).
			Interface("incoming_md:", incomingMd).
			Interface("req", req).
			Interface("outgoing_md:", outgoingMd).
			Interface("resp", reply).
			Interface("err", err).
			Msg("Client-Side")

		return err
	}
}
