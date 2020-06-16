# ceDemo Service

This is the CE Demo service

Showcase

1. Async service

## Usage

A Makefile is included for convenience

Build the binary

```bash
make build TARGET=cedemo TYPE=service VERSION=v0.1.1
```

Run the service

```bash
make run-cedemo
# or
go run service/cedemo/main.go
```

Build a docker image

```bash
make docker TARGET=cedemo TYPE=service VERSION=v0.1.1
```

Test the service

```bash
curl -X POST \
    -H "content-type: application/json"  \
    -H "ce-specversion: 1.0"  \
    -H "ce-source: curl-command"  \
    -H "ce-type: curl.demo"  \
    -H "ce-id: 123-abc"  \
    -d '{"name":"Sumo"}' \
    http://localhost:8080
```


## Reference
- https://github.com/knative/docs/blob/master/docs/serving/samples/cloudevents/cloudevents-go/README.md
- https://github.com/knative/docs/blob/master/docs/serving/samples/cloudevents/cloudevents-go/cloudevents.go
- https://github.com/spencer-p/moroncloudevents
