module github.com/xmlking/grpc-starter-kit

go 1.16

replace github.com/xmlking/grpc-starter-kit => ./

//replace github.com/xmlking/toolkit => ../toolkit

//replace github.com/xmlking/toolkit/confy => ../toolkit/confy

//replace github.com/xmlking/toolkit/logger => ../toolkit/logger

//replace github.com/xmlking/toolkit => github.com/xmlking/toolkit v0.1.2-0.20210125025404-51fc2d71fb2d

require (
	entgo.io/contrib v0.0.0-20210422070405-6ca335ec4de1
	entgo.io/ent v0.8.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.20.1
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.20.1
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/cockroachdb/errors v1.8.4
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/rs/zerolog v1.22.0
	github.com/sarulabs/di/v2 v2.4.2
	github.com/sercand/kuberesolver v2.4.0+incompatible
	github.com/soheilhy/cmux v0.1.5
	github.com/srikrsna/protoc-gen-gotag v0.5.0
	github.com/stretchr/testify v1.7.0
	github.com/tcfw/go-grpc-k8s-resolver v0.0.1
	github.com/thoas/go-funk v0.8.0
	github.com/xmlking/toolkit v0.1.7-0.20210602041816-47cebf71982e
	github.com/xmlking/toolkit/confy v0.1.7-0.20210602041816-47cebf71982e
	github.com/xmlking/toolkit/logger v0.1.7-0.20210602041816-47cebf71982e
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/stdout v0.20.0
	go.opentelemetry.io/otel/metric v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	go.opentelemetry.io/otel/sdk/metric v0.20.0
	go.opentelemetry.io/otel/trace v0.20.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)
