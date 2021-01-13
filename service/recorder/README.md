# Recorder

Recorder service saves `TransactionEvents` to configured go-micro `store`.
Transactions are published by other micro services i.e., `account`, `emailer`, `greeter`

## Usage

### Build the binary

```bash
make build TARGET=recorder TYPE=service
# then run with custom env
CONFIG_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/recorder-service
```

### Run the service

> (optional) set broker to googlepubsub

```bash
make run-recorder
# or
go run service/recorder/main.go
```

### Build a docker image

```bash
make docker TARGET=recorder TYPE=service VERSION=v0.1.1
```

### Test the service

```bash
```


## TODO
- Store in RocksDB?
