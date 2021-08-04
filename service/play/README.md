# Play

Experimenting with [OpenTelemetry](https://opentelemetry.io/) using [Google Cloud Trace](https://cloud.google.com/trace) and [Google Cloud Monitoring](https://cloud.google.com/monitoring).


## Usge

### Run

```bash
make run-play
# or
go run service/play/main.go

# enable metrics and tracking
export CONFY_FEATURES_METRICS_ENABLED=true
export CONFY_FEATURES_TRACING_ENABLED=true
# enable metrics target: `prometheus` and tracing target: `stdout`
export CONFY_FEATURES_METRICS_TARGET=prometheus
export CONFY_FEATURES_TRACING_TARGET=stdout
# when using with target: `gcp`
export GOOGLE_CLOUD_PROJECT=xyz
export GOOGLE_APPLICATION_CREDENTIALS=../../../Apps/micro-starter-kit.json

make run-play
```

#### prometheus exporter 

http://localhost:9213/metrics 

### Test

```bash
### TLS ###
grpcurl -insecure \
-protoset <(buf build -o -) \
-d '{"name": "sumo"}' 0.0.0.0:8084 mkit.service.greeter.v1.GreeterService/Hello
```

### Reference 
- https://github.com/open-telemetry/opentelemetry-go
- https://github.com/GoogleCloudPlatform/opentelemetry-operations-go
- https://cloud.google.com/trace/docs/setup/go-ot
