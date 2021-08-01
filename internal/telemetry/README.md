# telemetry

https://cloud.google.com/trace/docs/setup/go-ot
https://lightstep.com/blog/opentelemetry-101-what-are-metrics/
 https://medium.com/google-cloud/integrating-tracing-and-logging-with-opentelemetry-and-stackdriver-a5396fbc3e78

// https://github.com/cds-snc/covid-alert-server/blob/master/pkg/telemetry/telemetry.go
// https://github.com/liiling/kernel_metrics_agent/blob/master/otel-pipeline/main.go
// https://github.com/CovidShield/server/blob/master/pkg/telemetry/telemetry.go
// https://github.com/liiling/kernel_metrics_agent/blob/master/otel-pipeline/main.go

https://github.com/liiling/kernel_metrics_agent/tree/master/otel-pipeline
 
Adding new snap https://github.com/open-telemetry/opentelemetry-go


gRPC tracking

https://github.com/open-telemetry/opentelemetry-go/blob/master/example/grpc/server/main.go


### Examples

https://github.com/open-telemetry/opentelemetry-go/blob/master/exporters/stdout/example_test.go

```bash
export OTEL_RESOURCE_ATTRIBUTES=key=value,rk5=7
export OTEL_SERVICE_NAME=play-service
```
```go
import (
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/instrumentation/grpctrace"
)

func main() {
	config.Init()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpctrace.UnaryServerInterceptor(global.Tracer(""))),
		grpc.StreamInterceptor(grpctrace.StreamServerInterceptor(global.Tracer(""))),
	)

	api.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```


Metrics

https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/blob/master/example/metric/example.go

