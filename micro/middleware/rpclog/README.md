# RpcLogger

A set of Log wrappers for `Client`, `Server`  for `grpc` and `broker` for debugging.

## Usage

```go
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			rpclog.UnaryServerInterceptor(rpclog.WithExcludeMethods("/grpc.health.v1.Health/Check", "/api.MyService/*")),
		)),
	)
```
