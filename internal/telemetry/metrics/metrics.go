package metrics

// https://github.com/GoogleCloudPlatform/opentelemetry-operations-go/blob/master/example/metric/example.go
// https://github.com/liiling/kernel_metrics_agent/blob/master/otel-pipeline/main.go
import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	smetrics "go.opentelemetry.io/otel/exporters/stdout"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	//pmetrics "go.opentelemetry.io/otel/exporters/metric/prometheus"
	gmetrics "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"

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
func InitMetrics(cfg *config.Features_Metrics) *controller.Controller {
	once.Do(func() {
		log.Debug().Interface("MetricConfig", cfg).Msg("Initializing Metrics")
		var err error
		pushOpts := []controller.Option{
			controller.WithCollectPeriod(time.Second * 10),
		}
		if config.IsProduction() {
			projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
			opts := []gmetrics.Option{gmetrics.WithProjectID(projectID)}
			exporter, err = gmetrics.InstallNewPipeline(opts, pushOpts...)
		} else {
			opts := []smetrics.Option{
				smetrics.WithPrettyPrint(),
			}
			_, exporter, err = smetrics.InstallNewPipeline(opts, pushOpts)
		}
		if err != nil {
			log.Fatal().Err(err).Msg("failed to initialize metrics exporter")
		}
	})
	return exporter
}

//func initPrometheusMetrics() {
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
