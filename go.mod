module github.com/xmlking/grpc-starter-kit

go 1.17

//replace github.com/xmlking/toolkit => ../toolkit

//replace github.com/xmlking/toolkit/confy => ../toolkit/confy

//replace github.com/xmlking/toolkit/logger => ../toolkit/logger

//replace github.com/xmlking/toolkit/telemetry => ../toolkit/telemetry

//replace github.com/xmlking/toolkit => github.com/xmlking/toolkit v0.1.2-0.20210125025404-51fc2d71fb2d

require (
	entgo.io/contrib v0.2.0
	entgo.io/ent v0.10.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/cloudevents/sdk-go/v2 v2.8.0
	github.com/cockroachdb/errors v1.8.6
	github.com/envoyproxy/go-control-plane v0.10.1
	github.com/envoyproxy/protoc-gen-validate v0.6.3
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/mattn/go-sqlite3 v1.14.10
	github.com/rs/zerolog v1.26.1
	github.com/sarulabs/di/v2 v2.4.2
	github.com/sercand/kuberesolver v2.4.0+incompatible
	github.com/soheilhy/cmux v0.1.5
	github.com/srikrsna/protoc-gen-gotag v0.6.2
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/thoas/go-funk v0.9.1
	github.com/xmlking/toolkit v0.2.4
	github.com/xmlking/toolkit/broker/cloudevents v0.2.4
	github.com/xmlking/toolkit/confy v0.2.4
	github.com/xmlking/toolkit/logger v0.2.4
	github.com/xmlking/toolkit/telemetry v0.2.4
	github.com/xmlking/toolkit/xds v0.2.4
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.28.0
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/metric v0.26.0
	go.opentelemetry.io/otel/trace v1.3.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

require (
	ariga.io/atlas v0.3.2-0.20220120225051-c3fac7d636dd // indirect
	cloud.google.com/go/compute v0.1.0 // indirect
	cloud.google.com/go/monitoring v0.1.0 // indirect
	cloud.google.com/go/trace v1.0.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v1.0.0-RC1.0.20220121174109-24420159f1b4 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v1.0.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cncf/xds/go v0.0.0-20211011173535-cb28da3451f1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20190617123548-eb05cc24525f // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/cockroachdb/sentry-go v0.6.1-cockroachdb.2 // indirect
	github.com/creasty/defaults v1.5.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/hashicorp/hcl/v2 v2.11.1 // indirect
	github.com/jhump/protoreflect v1.10.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.12.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.3.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/sdk v1.3.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.26.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.20.0 // indirect
	golang.org/x/crypto v0.0.0-20211215165025-cf75a172585e // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.9-0.20211216111533-8d383106f7e7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/api v0.65.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
