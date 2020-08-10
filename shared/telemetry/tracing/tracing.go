package tracing

// https://cloud.google.com/trace/docs/setup/go-ot#gke
// https://medium.com/google-cloud/integrating-tracing-and-logging-with-opentelemetry-and-stackdriver-a5396fbc3e78

// Adding new snap https://github.com/open-telemetry/opentelemetry-go
import (
	"os"
	"sync"
	"time"

	gtrace "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/rs/zerolog/log"
	strace "go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/xmlking/grpc-starter-kit/shared/config"
)

var (
	once      sync.Once
	closeFunc func()
)

// expected GOOGLE_CLOUD_PROJECT & GOOGLE_APPLICATION_CREDENTIALS Environment Variable set

// // before ending program, wait for all enqueued spans to be exported
func InitTracing(cfg *config.Features_Tracing) func() {
	once.Do(func() {
		log.Debug().Interface("TracingConfig", cfg).Msg("Initializing Tracing")
		if config.IsProduction() {
			sampling := cfg.Sampling
			projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
			_, flush, err := gtrace.InstallNewPipeline(
				[]gtrace.Option{gtrace.WithProjectID(projectID)},
				// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
				// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
				// sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
				sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.ProbabilitySampler(sampling)}),
			)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize google tracing exporter")
			}
			closeFunc = func() {
				flush()
			}
		} else {
			opts := []strace.Option{
				strace.WithQuantiles([]float64{0.5}),
				strace.WithPrettyPrint(),
			}
			pushOpts := []push.Option{
				push.WithPeriod(time.Second * 10),
			}
			// Registers both a trace and meter Provider globally.
			pipeline, err := strace.InstallNewPipeline(opts, pushOpts)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize stdout tracing exporter")
			}
			closeFunc = func() {
				pipeline.Stop()
			}
		}
	})
	return closeFunc
}

//func InitTracing_old(cfg *config.Features_Tracing) {
//    once.Do(func() {
//        log.Debug().Interface("TracingConfig", cfg).Msg("Initializing Tracing")
//        var exporter trace.SpanSyncer
//        var err error
//        if config.IsProduction() {
//            projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
//            exporter, err = gtrace.NewExporter(gtrace.WithProjectID(projectID))
//            if err != nil {
//                log.Fatal().Err(err).Msg("failed to initialize google tracing exporter")
//            }
//        } else {
//            exporter, err = strace.NewExporter(strace.WithPrettyPrint())
//            if err != nil {
//                log.Fatal().Err(err).Msg("failed to initialize stdout tracing exporter")
//            }
//        }
//
//        // Create trace provider with the exporter.
//        //
//        // By default it uses AlwaysSample() which samples all traces.
//        // In a production environment or high QPS setup please use
//        // ProbabilitySampler set at the desired probability.
//        sampling := cfg.Sampling
//        tp, err := sdktrace.NewProvider(
//            // sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
//            sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.ProbabilitySampler(sampling)}),
//            sdktrace.WithSyncer(exporter),
//        )
//        if err != nil {
//            log.Fatal().Err(err).Send()
//        }
//        global.SetTraceProvider(tp)
//    })
//}
