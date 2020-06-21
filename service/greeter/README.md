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

1. start greeter service first
   `make run-greeter`
2. start envoy via docker compose
    `docker-compose up envoy`

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

#### test API via envoy with TLS, and client cert

>  make sure `require_client_certificate: true` is enabled in `listeners.yaml` for following tests:

```bash
# test with good cleint cert
grpcurl -cacert=config/certs/ca-cert.pem \
-cert=config/certs/client-cert.pem \
-key=config/certs/client-key.pem \
-protoset <(buf image build -o -) \
-d '{"name": "sumo"}' localhost:9444 mkit.service.greeter.v1.GreeterService/Hello

# test with wrong client cert. This will fail!
grpcurl -cacert=config/certs/ca-cert.pem \
-cert=config/certs/upstream-cert.pem \
-key=config/certs/upstream-key.pem \
-protoset <(buf image build -o -) \
-d '{"name": "sumo"}' localhost:9444 mkit.service.greeter.v1.GreeterService/Hello

# testing with request data from file.
grpcurl -cacert=config/certs/ca-cert.pem \
-protoset <(buf image build -o -) \
-v -H trans_id=abc123 \
-d @ localhost:9444 mkit.service.greeter.v1.GreeterService/Hello <test/echo-request.json
```

### Test gRPC-Web

```bash
# without TLS

 curl 'http://localhost:9090/mkit.service.greeter.v1.GreeterService/Hello' \
 -H 'Content-Type: application/grpc-web+proto' \
 -H 'X-Grpc-Web: 1' \
 -H 'custom-header-1: value1' \
 -H 'Accept: */*' \
 -H 'Connection: keep-alive' \
 --data-binary $'\x00\x00\x00\x00\x05\n\x03abc' --compressed
```

### Envoy Health checks

```bash
curl -v http://localhost:9090/healthz \
-H 'x-envoy-livenessprobe: healthz'

curl -v https://localhost:9444/healthz \
-H 'x-envoy-livenessprobe: healthz'
```
