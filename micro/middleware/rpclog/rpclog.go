package rpclog

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/xmlking/grpc-starter-kit/shared/constants"
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

		resp, err := handler(ctx, req)

		log.Debug().
			Str("module", "debug").
			Str("method", info.FullMethod).
			Str("trace_id", metautils.ExtractIncoming(ctx).Get(constants.TraceIDKey)).
			Str("tenant_id", metautils.ExtractIncoming(ctx).Get(constants.TenantIdKey)).
			Str("from_service", metautils.ExtractIncoming(ctx).Get(constants.FromServiceKey)).
			Interface("req", req).
			Interface("resp", resp).
			Interface("err", err).
			Msg("Server-Side rpclog")

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

		err := invoker(ctx, method, req, reply, cc, opts...)

		log.Debug().
			Str("module", "debug").
			Str("method", method).
			Str("trace_id", metautils.ExtractOutgoing(ctx).Get(constants.TraceIDKey)).
			Str("tenant_id", metautils.ExtractOutgoing(ctx).Get(constants.TenantIdKey)).
			Str("from_service", metautils.ExtractOutgoing(ctx).Get(constants.FromServiceKey)).
			Interface("req", req).
			Interface("resp", reply).
			Interface("err", err).
			Msg("Client-Side rpclog")

		return err
	}
}
