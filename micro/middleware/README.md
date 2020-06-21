# Middleware

1. rpclog
2. duration
3. validator

## Get

```bash
go get github.com/xmlking/grpc-starter-kit/micro
```

## Usage

Execution is done in left-to-right order.

Add interceptors in following order
1. Around interceptors - from outer to inner — e.g., duration,  retry  
2. Before interceptors - rate-limit, auth, validation , tagging 
3. After interceptors - rpclog, translog, recovery

```go
import (
    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
    grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
    "github.com/xmlking/grpc-starter-kit/micro/middleware/rpclog"
)

server := grpc.NewServer(
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
        // Execution is done in left-to-right order
        grpc_validator.UnaryServerInterceptor(),
        // keep it last in the interceptor chain
        rpclog.UnaryServerInterceptor(),
    )),
    grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
        // keep it last in the interceptor chain
        rpclog.StreamServerInterceptor()
    )),
)
```
