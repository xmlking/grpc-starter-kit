# Mock

Running mock gRPC server from proto files for testing.

### Prerequisites

```bash
# gRPC mock server for testing
npm i bloomrpc-mock -g
# install `bloomrpc` via `brew` into ~/Applications)
brew install --cask --appdir=~/Applications bloomrpc
```

> use certs generated from [mtls](../../config/certs/README.md)

### Run

```bash
bloomrpc-mock service/greeter/proto/greeter/greeter.proto
# Or
bloomrpc-mock e2e/account.bloomrpc.proto \
-r config/base/secrets/certs/upstream-cert.pem \
-k config/base/secrets/certs/client-key.pem,config/base/secrets/certs/client-cert.pem \
-i ~/go/src  -i /usr/local/Cellar/protobuf/3.11.2/include \
-i ~/Developer/Work/go/grpc-starter-kit
```
