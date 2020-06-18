# Greeter Service

This is the Greeter service

Showcase

1. Async service

## Usage

### Build the binary

```bash
make build TARGET=greeter TYPE=service
# then run with custom env
CONFIGOR_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/greeter-service
```

### Run the service

```bash
make run-greeter
# or
go run service/greeter/main.go
```

### Build a docker image

```bash
make docker TARGET=greeter TYPE=service VERSION=v0.1.1
```

### Test the service

```bash
# start greeter service first
make run-greeter

# test with grpc cli
grpcurl -plaintext -proto proto/mkit/service/greeter/v1/greeter.proto list
grpcurl -plaintext -proto proto/mkit/service/greeter/v1/greeter.proto describe
grpcurl -plaintext -proto proto/mkit/service/greeter/v1/greeter.proto -d '{"name": "sumo"}' localhost:8081  mkit.service.greeter.v1.GreeterService/Hello
