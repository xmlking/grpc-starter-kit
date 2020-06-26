# RpcLogger

A set of interceptors for `Client` and `Server` for debug **logging** request/response/errors.

## Usage

```go
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			rpclog.UnaryServerInterceptor(rpclog.WithExcludeMethods("/grpc.health.v1.Health/Check", "/api.MyService/*")),
		)),
	)
```
