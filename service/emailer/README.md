# Emailer Service

This is the Emailer service

Showcase

1. Async service

## Usage

A Makefile is included for convenience

### Build the binary

```bash
make build TARGET=emailer TYPE=service VERSION=v0.1.1
```

### Run the service

```bash
make run-emailer
# or
go run service/emailer/main.go
```

### Build a docker image

```bash
make docker TARGET=emailer TYPE=service VERSION=v0.1.1
```

### Test the service

```bash
curl -X POST \
    -H "content-type: application/json"  \
    -H "ce-specversion: 1.0"  \
    -H "ce-source: curl-command"  \
    -H "ce-type: curl.demo"  \
    -H "ce-id: 123-abc"  \
    -d '{"Subject": "Sumo", "To": "sumo@demo.com"}' \
    http://localhost:8082
```
