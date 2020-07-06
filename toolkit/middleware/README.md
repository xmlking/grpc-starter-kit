# Middleware

1. rpclog
2. duration
3. [validator](github.com/grpc-ecosystem/go-grpc-middleware/validator)
4. [grpc-opentracing](https://github.com/grpc-ecosystem/grpc-opentracing)

## Get

```bash
go get github.com/xmlking/grpc-starter-kit/micro
```

## Usage

Interceptors will be executed **from left to right**: e.g., logging, monitoring and auth.
```go
grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(loggingUnary, monitoringUnary, authUnary),)
```

Add interceptors in following order
1. Around interceptors - from outer to inner — e.g., duration,  retry  
2. Before interceptors - rate-limit, auth, validation , tagging 
3. After interceptors - rpclog, translog, recovery

```go
import (
    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
    grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
    "github.com/xmlking/grpc-starter-kit/shared/middleware/rpclog"
)

server := grpc.NewServer(
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
        // Execution is done in left-to-right order
        grpc_validator.UnaryServerInterceptor(),
        // keep it last in the interceptor chain
        rpclog.UnaryServerInterceptor(rpclog.WithExcludeMethods("/grpc.health.v1.Health/Check", "/api.MyService/*")),
    )),
    grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
        // keep it last in the interceptor chain
        rpclog.StreamServerInterceptor()
    )),
)
```
