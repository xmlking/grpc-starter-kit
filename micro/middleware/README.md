# Middleware

1. rpclog
2. duration
3. validator

## Get

```bash
go get github.com/xmlking/grpc-starter-kit/micro
```

## Usage

```go
import (
    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
    grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
    "github.com/xmlking/grpc-starter-kit/micro/middleware/rpclog"
)

server := grpc.NewServer(
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
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
