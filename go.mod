module github.com/xmlking/grpc-starter-kit

go 1.16

replace github.com/xmlking/grpc-starter-kit => ./

//replace github.com/xmlking/toolkit => /Users/schintha/Developer/Work/go/toolkit

//replace github.com/xmlking/toolkit => github.com/xmlking/toolkit v0.1.2-0.20210125025404-51fc2d71fb2d

//replace github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace => github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.15.1-0.20210128222945-8d20ac5c4dd4

//replace github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric => github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.15.1-0.20210128222945-8d20ac5c4dd4

require (
	entgo.io/ent v0.7.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.19.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.19.0
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/cockroachdb/errors v1.8.3
	github.com/envoyproxy/protoc-gen-validate v0.5.1
	github.com/golang/protobuf v1.5.1
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/rs/zerolog v1.21.0
	github.com/sarulabs/di/v2 v2.4.2
	github.com/sercand/kuberesolver v2.4.0+incompatible
	github.com/soheilhy/cmux v0.1.4
	github.com/srikrsna/protoc-gen-gotag v0.5.0
	github.com/stretchr/testify v1.7.0
	github.com/tcfw/go-grpc-k8s-resolver v0.0.0-20201027075059-d3a2d14aa08f
	github.com/thoas/go-funk v0.8.0
	github.com/xmlking/toolkit v0.1.3-0.20210326062219-686e14c11b53
	github.com/xmlking/toolkit/confy v0.1.3-0.20210326062219-686e14c11b53
	github.com/xmlking/toolkit/logger v0.0.0-20210326062219-686e14c11b53
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.19.0
	go.opentelemetry.io/otel v0.19.0
	go.opentelemetry.io/otel/exporters/stdout v0.19.0
	go.opentelemetry.io/otel/metric v0.19.0
	go.opentelemetry.io/otel/sdk v0.19.0
	go.opentelemetry.io/otel/sdk/metric v0.19.0
	go.opentelemetry.io/otel/trace v0.19.0
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/examples v0.0.0-20210318230139-bce1cded4b05 // indirect
	google.golang.org/protobuf v1.26.0
)
