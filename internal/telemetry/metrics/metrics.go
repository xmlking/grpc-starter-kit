package metrics

// https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/blob/master/example/metric/example.go
// https://github.com/liiling/kernel_metrics_agent/blob/master/otel-pipeline/main.go
import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	gmetrics "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	export "go.opentelemetry.io/otel/sdk/export/metric"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	once sync.Once

	exporter *controller.Controller
)

// InitMetrics expected GOOGLE_CLOUD_PROJECT & GOOGLE_APPLICATION_CREDENTIALS Environment Variable set
func InitMetrics(ctx context.Context, cfg *config.Features_Metrics) func() {
	once.Do(func() {
		log.Debug().Interface("MetricConfig", cfg).Msg("Initializing Metrics")
		var err error
		if config.IsProduction() {
			projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
			opts := []gmetrics.Option{gmetrics.WithProjectID(projectID)}
			pushOpts := []controller.Option{
				controller.WithCollectPeriod(time.Second * 10),
			}
			//resOpt := basic.WithResource(resource.NewWithAttributes(
			//    semconv.SchemaURL,
			//    attribute.String("instance_id", "abc123"),
			//    attribute.String("application", "example-app"),
			//))
			exporter, err = gmetrics.InstallNewPipeline(opts, pushOpts...)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize metrics exporter")
			}
		} else {
			opts := []stdoutmetric.Option{
				stdoutmetric.WithPrettyPrint(),
			}
			var metricExporter *stdoutmetric.Exporter
			metricExporter, err = stdoutmetric.New(opts...)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize metrics exporter")
			}
			exporter = controller.New(
				processor.New(
					simple.NewWithExactDistribution(),
					metricExporter,
				),
				controller.WithExporter(metricExporter),
				controller.WithCollectPeriod(5*time.Second),
			)
			err = exporter.Start(ctx)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize metrics controller")
			}
		}

		// Registers metrics Provider globally.
		global.SetMeterProvider(exporter.MeterProvider())
		propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
		otel.SetTextMapPropagator(propagator)
	})

	return func() {
		exporter.Stop(ctx)
	}
}

// InitPrometheusMetrics Initialize Prometheus Metrics
// Usage: https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go
func InitPrometheusMetrics(ctx context.Context, cfg *config.Features_Metrics) func() {
	pConfig := prometheus.Config{}
	pController := controller.New(
		processor.New(
			simple.NewWithHistogramDistribution(
				histogram.WithExplicitBoundaries(pConfig.DefaultHistogramBoundaries),
			),
			export.CumulativeExportKindSelector(),
			processor.WithMemory(true),
		),
		controller.WithCollectPeriod(0),
		controller.WithResource(resource.Empty()),
	)

	exporter, err := prometheus.New(pConfig, pController)

	if err != nil {
		log.Fatal().Err(err).Msgf("failed to initialize prometheus exporter")
	}

	// Registers metrics Provider globally.
	global.SetMeterProvider(exporter.MeterProvider())
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(propagator)

	http.HandleFunc("/metrics", exporter.ServeHTTP)

	port := 2222
	pSrv := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	go pSrv.ListenAndServe()

	log.Info().Msgf("Prometheus server running on :%d\n", port)

	return func() {
		log.Info().Msgf("Stopping prometheus metrics server...")
		pSrv.Shutdown(ctx)
	}
}
