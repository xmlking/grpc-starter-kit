# xds

This is the Greeter service

## Usage

### Build the binary

```bash
make build TARGET=greeter TYPE=service
# then run with custom env
CONFY_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/greeter-service
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

### Test

