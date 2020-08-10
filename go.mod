module github.com/xmlking/grpc-starter-kit

go 1.14

// replace github.com/xmlking/configor => /Users/schintha/Developer/Work/go/configor
// replace github.com/xmlking/toolkit => /Users/schintha/Developer/Work/go/toolkit
replace github.com/xmlking/toolkit => github.com/xmlking/toolkit v0.1.1-0.20200710152301-9029e438457f

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.2.2-0.20200810002127-a39a4657df0b
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.2.2-0.20200810002127-a39a4657df0b
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/facebookincubator/ent v0.3.0
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.0
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.15
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.19.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/sercand/kuberesolver v2.4.0+incompatible
	github.com/soheilhy/cmux v0.1.4
	github.com/stretchr/testify v1.6.1
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/configor v0.2.1
	github.com/xmlking/toolkit v0.1.0
	go.opentelemetry.io/otel v0.10.0
	go.opentelemetry.io/otel/exporters/stdout v0.10.0
	go.opentelemetry.io/otel/sdk v0.10.0
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485
	google.golang.org/grpc v1.31.0
	google.golang.org/grpc/examples v0.0.0-20200810225334-2983360ff4e7 // indirect
	google.golang.org/protobuf v1.25.0
)
