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

1. start greeter service
   `make run-greeter`
2. start envoy via docker-compose
    `docker-compose up envoy` or `docker-compose up envoy_secure_backend` when **tls** is enabled in [config.yaml](/config/config.yaml)

> envoy via docker-compose will be exposing '9901', '9090', '9444' ports

####  test API directly  (go greeter)
```bash
### plaintext ###
grpcurl -plaintext -proto proto/mkit/service/greeter/v1/greeter.proto list
grpcurl -plaintext -proto proto/mkit/service/greeter/v1/greeter.proto describe
grpcurl -plaintext -proto proto/mkit/service/greeter/v1/greeter.proto -d '{"name": "sumo"}' localhost:8081  mkit.service.greeter.v1.GreeterService/Hello
# OR 
grpcurl -plaintext \
-protoset <(buf image build -o -) \
-d '{"name": "sumo"}' 0.0.0.0:8081 mkit.service.greeter.v1.GreeterService/Hello

### TLS ###
grpcurl -insecure \
-protoset <(buf image build -o -) \
-d '{"name": "sumo"}' 0.0.0.0:8081 mkit.service.greeter.v1.GreeterService/Hello
```
#### test API via envoy
```bash
### plaintext ###
grpcurl -plaintext  \
-protoset <(buf image build -o -) \
-d '{"name": "sumo"}' 0.0.0.0:9090 mkit.service.greeter.v1.GreeterService/Hello

### TLS ###
grpcurl -cacert=config/certs/ca-cert.pem \
-protoset <(buf image build -o -) \
-d '{"name": "sumo"}' localhost:9444 mkit.service.greeter.v1.GreeterService/Hello
```

For full examples, see [Testing with grpcurl](/docs/testing/grpcurl.md)
