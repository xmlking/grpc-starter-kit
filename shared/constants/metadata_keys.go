package constants

// metadata constants
const (
    // There are certain requirements for metadata to be passed in the http header:
    // gRPC recommended Key format: `lowercase alphanumeric characters and hyphen`
    TraceIDKey     = "mkit-trace-id"
    FromServiceKey = "mkit-from-service"
    TenantIdKey    = "mkit-tenant-id"
)
