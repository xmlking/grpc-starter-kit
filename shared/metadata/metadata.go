package metadata

import (
    "context"
    "strings"

    "google.golang.org/grpc/metadata"
)

// metadata constants
const (
    // There are certain requirements for metadata to be passed in the http header:
    // gRPC recommended Key format: `lowercase alphanumeric characters and hyphen`
    TraceIDKey     = "mkit-trace-id"
    FromServiceKey = "mkit-from-service"
    TenantIdKey    = "mkit-tenant-id"
)

func GetFirst(md metadata.MD, key string) (val string, ok bool) {
    var vals []string
    // attempt to get as is
    if vals, ok = md[key]; ok && len(val) > 0 {
        val = vals[0]
    }

    // attempt to get lower case
    if vals, ok = md[strings.Title(key)]; ok && len(val) > 0 {
        val = vals[0]
    }
    return
}

func GetTenantID(ctx context.Context) (string, bool) {
    if md, ok := metadata.FromIncomingContext(ctx); ok {
        return GetFirst(md, TenantIdKey)
    }
    return "", false
}

func GetFromService(ctx context.Context) (string, bool) {
    if md, ok := metadata.FromIncomingContext(ctx); ok {
        return GetFirst(md, FromServiceKey)
    }
    return "", false
}

func GetTraceID(ctx context.Context) (string, bool) {
    if md, ok := metadata.FromIncomingContext(ctx); ok {
        return GetFirst(md, TraceIDKey)
    }
    return "", false
}

func SetTraceID(ctx context.Context, tranId string) context.Context {
    return metadata.AppendToOutgoingContext(ctx, TraceIDKey, tranId)
}
