# API Testing 

API testing with **grpcurl** via Envoy


## Test the service

1. start greeter service
   `make run-greeter`
2. start envoy via docker-compose
    `docker-compose up envoy` or `docker-compose up envoy_secure_backend` when **tls** is enabled in [config.yaml](/config/config.yaml)

> envoy via docker-compose will be exposing '9901', '9090', '9444' ports

### Test API directly (go greeter)
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
### Test API via envoy
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

### Test API via Envoy with TLS, and client cert

1. set **tls** enabled **true** in [config.yaml](/config/config.yaml) and start greeter service
   `make run-greeter`
2. start envoy via docker-compose
    `docker-compose up envoy_secure_backend_cleint_cert`
    
```bash
# test with good cleint cert
grpcurl -cacert=config/certs/ca-cert.pem \
-cert=config/certs/client-cert.pem \
-key=config/certs/client-key.pem \
-protoset <(buf image build -o -) \
-d '{"name": "sumo"}' localhost:9444 mkit.service.greeter.v1.GreeterService/Hello

# test with wrong client cert. This will fail! You will see `TLS error:` in envoy logs
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

# ???

curl 'http://localhost:9090/mkit.service.greeter.v1.GreeterService/Hello' \
-H 'Accept: application/grpc-web-text' \
-H 'X-User-Agent: grpc-web-javascript/0.1' \
-H 'X-Grpc-Web: 1' \
-H 'Content-Type: application/grpc-web-text' \
-H 'Origin: http://localhost:4200' \
-H 'Sec-Fetch-Site: same-site' \
-H 'Sec-Fetch-Mode: cors' \
--data-binary 'AAAAAAYKBHN1bW8=' --compressed
```

### Envoy Health checks

```bash
curl -v http://localhost:9090/healthz \
-H 'x-envoy-livenessprobe: healthz'

curl -v https://localhost:9444/healthz \
-H 'x-envoy-livenessprobe: healthz'
```
