module github.com/xmlking/grpc-starter-kit

go 1.14

// replace github.com/xmlking/configor => /Users/schintha/Developer/Work/go/configor
// replace github.com/xmlking/toolkit => /Users/schintha/Developer/Work/go/toolkit
replace github.com/xmlking/toolkit => github.com/xmlking/toolkit v0.1.1-0.20200710152301-9029e438457f

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/cloudevents/sdk-go/v2 v2.1.0
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.0
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.14
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.19.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/soheilhy/cmux v0.1.4
	github.com/stretchr/testify v1.6.1
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/configor v0.2.1
	github.com/xmlking/toolkit v0.1.0
	google.golang.org/genproto v0.0.0-20200624020401-64a14ca9d1ad
	google.golang.org/grpc v1.30.0
	google.golang.org/grpc/examples v0.0.0-20200630190442-3de8449f8555 // indirect
	google.golang.org/protobuf v1.25.0
)
