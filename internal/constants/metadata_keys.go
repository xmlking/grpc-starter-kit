package constants

// metadata constants
const (
	// There are certain requirements for metadata to be passed in the http header:
	// gRPC recommended Key format: `lowercase alphanumeric characters and hyphen`
	TraceIDKey     = "gkit-trace-id"
	FromServiceKey = "gkit-from-service"
	TenantIDKey    = "gkit-tenant-id"
)
