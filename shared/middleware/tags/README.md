# Tags 

A set of interceptors for `Client` and `Server` for **appending** or **forwarding** selective metadata tags in the context.

1. append (client/server) --> append tags e.g., "mkit-from-service" to outgoing
2. forward (server) -->  copy tags e.g., "mkit-trace-id" from incoming to outgoing



## Usage

```go
    // Server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			append.UnaryServerInterceptor(append.WithPairs(constants.FromServiceKey, "account-service")),
            forward.UnaryServerInterceptor(forward.WithForwardTags(constants.TraceIDKey)),
		)),
	)
    // Client 
    var ucInterceptors = []grpc.UnaryClientInterceptor{
        appendTags.UnaryClientInterceptor(appendTags.WithTraceID(), appendTags.WithPairs(constants.FromServiceKey, constants.ACCOUNT_CLIENT)),
    }
    dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(ucInterceptors...)))
    clientConn, err = grpc.Dial(service.Endpoint, dialOptions...)
```

Access **Tags** in your handler

```go
outgoingMd := metautils.ExtractIncoming(ctx)
println( outgoingMd.Get(constants.TraceIDKey)))
println( outgoingMd.Get(constants.TenantIdKey)))
println( outgoingMd.Get(constants.FromServiceKey)))
```
