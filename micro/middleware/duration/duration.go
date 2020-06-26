package duration

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor logs the time taken on server-side
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		var deadline string
		if d, ok := ctx.Deadline(); ok {
			deadline = d.Format(time.RFC3339)
		}

		started := time.Now().UTC()

		resp, err := handler(ctx, req)

		finished := time.Now().UTC()

		log.Debug().
			Str("module", "duration").
			Str("method", info.FullMethod).
			Str("start", started.Format(time.RFC3339Nano)).
			Str("finish", finished.Format(time.RFC3339Nano)).
			Str("duration", finished.Sub(started).String()). // time.Since(started).String()) ?
			Str("deadline", deadline).
			Msg("Server-Side Duration")

		return resp, err
	}
}

// UnaryClientInterceptor logs the time taken on client-side
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		var deadline string
		if d, ok := ctx.Deadline(); ok {
			deadline = d.Format(time.RFC3339)
		}

		started := time.Now().UTC()

		err := invoker(ctx, method, req, reply, cc, opts...)

		finished := time.Now().UTC()

		log.Debug().
			Str("module", "duration").
			Str("method", method).
			Str("start", started.Format(time.RFC3339Nano)).
			Str("finish", finished.Format(time.RFC3339Nano)).
			Str("duration", finished.Sub(started).String()). // time.Since(started).String()) ?
			Str("deadline", deadline).
			Msg("Client-Side Duration")

		return err
	}
}
