package forward

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor forward tags from incoming to outgoing
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	options := defaultOptions()
	for _, o := range opts {
		o(&options)
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var kv []string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if md.Len() > 0 {
				for _, k := range options.tags {
					if v := md.Get(k); len(v) > 0 {
						kv = append(kv, k, v[0])
					}
				}
			}
		}
		if len(kv) > 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, kv...)
		}

		return handler(ctx, req)
	}
}
