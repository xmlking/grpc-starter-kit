package rpclog

import (
    "context"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc"

    "github.com/xmlking/grpc-starter-kit/shared/metadata"
)

// UnaryServerInterceptor is an example server-side request logger middleware
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {

    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

        resp, err := handler(ctx, req)

        traceId, _ := metadata.GetTraceID(ctx)
        tenantId, _ := metadata.GetTenantID(ctx)
        fromService, _ := metadata.GetFromService(ctx)
        log.Debug().
            Str("module", "debug").
            Str("method", info.FullMethod).
            Str("trace_id", traceId).
            Str("tenant_id", tenantId).
            Str("from_service", fromService).
            Interface("req", req).
            Interface("resp", resp).
            Interface("err", err).
            Msg("Server-Side Handler")

        return resp, err
    }
}

// UnaryClientInterceptor is an example client-side request logger middleware
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {

    return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

        err := invoker(ctx, method, req, reply, cc, opts...)

        traceId, _ := metadata.GetTraceID(ctx)
        tenantId, _ := metadata.GetTenantID(ctx)
        fromService, _ := metadata.GetFromService(ctx)
        log.Debug().
            Str("module", "debug").
            Str("method", method).
            Str("trace_id", traceId).
            Str("tenant_id", tenantId).
            Str("from_service", fromService).
            Interface("req", req).
            Interface("reply", reply).
            Msg("Client-Side Call")

        return err
    }
}
