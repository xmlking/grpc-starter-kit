# Play

Experimenting with [OpenTelemetry](https://opentelemetry.io/) using [Google Cloud Trace](https://cloud.google.com/trace) and [Google Cloud Monitoring](https://cloud.google.com/monitoring).


## Usge

### Run

```bash
make run-play
# or
go run service/play/main.go

# run with `production` mode (using GCP as telemetry backend)
export CONFY_ENV=production
export GOOGLE_APPLICATION_CREDENTIALS=../../../Apps/micro-starter-kit.json
make run-play
```

### Test

```bash
### TLS ###
grpcurl -insecure \
-protoset <(buf build -o -) \
-d '{"name": "sumo"}' 0.0.0.0:8084 mkit.service.greeter.v1.GreeterService/Hello

# when server running in `production` mode
grpcurl -insecure \
-protoset <(buf build -o -) \
-d '{"name": "sumo"}' 0.0.0.0:8080 mkit.service.greeter.v1.GreeterService/Hello
```

### Reference 
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/GoogleCloudPlatform/opentelemetry-operations-go
- https://cloud.google.com/trace/docs/setup/go-ot
