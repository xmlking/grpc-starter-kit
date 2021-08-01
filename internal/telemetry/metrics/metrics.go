package metrics

// https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/blob/master/example/metric/example.go
// https://github.com/liiling/kernel_metrics_agent/blob/master/otel-pipeline/main.go
import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	//pmetrics "go.opentelemetry.io/otel/exporters/metric/prometheus"
	gmetrics "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"

	"github.com/xmlking/grpc-starter-kit/internal/config"
)

// https://github.com/cds-snc/covid-alert-server/blob/master/pkg/telemetry/telemetry.go
// https://github.com/liiling/kernel_metrics_agent/blob/master/otel-pipeline/main.go
// https://github.com/CovidShield/server/blob/master/pkg/telemetry/telemetry.go
// https://github.com/liiling/kernel_metrics_agent/blob/master/otel-pipeline/main.go

//exporter := metrics.InitMetrics(cfg.Features.Metrics)
//defer exporter.Stop()

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

	})
	// Registers metrics Provider globally.
	global.SetMeterProvider(exporter.MeterProvider())
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(propagator)

	return func() {
		exporter.Stop(ctx)
	}
}

//func initPrometheusMetrics(ctx context.Context, cfg *config.Features_Metrics) func() {
//	once.Do(func() {
//		exporter, err := pmetrics.InstallNewPipeline(
//			pmetrics.Config{
//				DefaultHistogramBoundaries: []float64{-0.5, 1},
//			},
//			pull.WithCachePeriod(time.Second*10),
//		)
//		if err != nil {
//			log.Fatal().Err(err).Msgf("failed to initialize prometheus exporter")
//		}
//
//		port := 2112
//		http.HandleFunc("/metrics", exporter.ServeHTTP)
//		go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
//		log.Info().Msgf("Prometheus server running on :%d\n", port)
//	})
//}
