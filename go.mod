module github.com/xmlking/grpc-starter-kit

go 1.15

// replace github.com/xmlking/configor => /Users/schintha/Developer/Work/go/configor
// replace github.com/xmlking/toolkit => /Users/schintha/Developer/Work/go/toolkit
replace github.com/xmlking/toolkit => github.com/xmlking/toolkit v0.1.1-0.20200710152301-9029e438457f

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.13.0
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.13.0
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/cockroachdb/errors v1.8.2
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/facebook/ent v0.5.3
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/sercand/kuberesolver v2.4.0+incompatible
	github.com/soheilhy/cmux v0.1.4
	github.com/stretchr/testify v1.6.1
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/configor v0.2.2
	github.com/xmlking/toolkit v0.1.1
	go.opentelemetry.io/otel v0.15.0
	go.opentelemetry.io/otel/exporters/stdout v0.15.0
	go.opentelemetry.io/otel/sdk v0.15.0
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
