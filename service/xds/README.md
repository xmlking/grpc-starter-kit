# xds

This is the xDS service

## Usage

### Build the binary

```bash
make build TARGET=xds TYPE=service
# then run with custom env
CONFY_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/xds-service
```

### Run the service

```bash
make run-xds
# or
go run service/xds/main.go
```

### Build a docker image

```bash
make docker TARGET=xds TYPE=service VERSION=v0.1.1
```

### Test

