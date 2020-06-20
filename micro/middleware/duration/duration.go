package duration

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor logs the time taken by the service.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		started := time.Now().UTC()

		resp, err := handler(ctx, req)

		finished := time.Now().UTC()

		log.Debug().
			Str("module", "duration").
			Str("method", info.FullMethod).
			Str("start", started.Format(time.RFC3339Nano)).
			Str("finish", finished.Format(time.RFC3339Nano)).
			Str("duration", finished.Sub(started).String()).
			Str("duration", finished.Sub(started).String()).
			Msg("")

		return resp, err
	}
}
