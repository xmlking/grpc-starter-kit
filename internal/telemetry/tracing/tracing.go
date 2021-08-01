package tracing

// https://cloud.google.com/trace/docs/setup/go-ot#gke
// https://medium.com/google-cloud/integrating-tracing-and-logging-with-opentelemetry-and-stackdriver-a5396fbc3e78

// Adding new snap https://github.com/open-telemetry/opentelemetry-go
import (
	"context"
	"os"
	"sync"

	cloudtrace "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/rs/zerolog/log"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	once      sync.Once
	tp        *trace.TracerProvider
	closeFunc func()
)

// expected GOOGLE_CLOUD_PROJECT & GOOGLE_APPLICATION_CREDENTIALS Environment Variable set

// InitTracing before ending program, wait for all enqueued spans to be exported
func InitTracing(ctx context.Context, cfg *config.Features_Tracing) func() {
	once.Do(func() {
		log.Debug().Interface("TracingConfig", cfg).Msg("Initializing Tracing")
		if config.IsProduction() {
			println("---------")
			projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
			exporter, err := cloudtrace.New(cloudtrace.WithProjectID(projectID))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize google tracing exporter")
			}
			tp = sdktrace.NewTracerProvider(
				// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
				// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
				sdktrace.WithSampler(sdktrace.AlwaysSample()),
				sdktrace.WithBatcher(exporter),
			)

			closeFunc = func() {
				exporter.Shutdown(ctx)
				tp.Shutdown(ctx)
			}
		} else {
			opts := []stdouttrace.Option{
				stdouttrace.WithPrettyPrint(),
			}
			exporter, err := stdouttrace.New(opts...)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize stdout tracing exporter")
			}
			bsp := sdktrace.NewBatchSpanProcessor(exporter)
			tp = sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))

			closeFunc = func() {
				exporter.Shutdown(ctx)
				tp.Shutdown(ctx)
			}
		}
	})

	// Registers trace Provider globally.
	otel.SetTracerProvider(tp)
	//propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	//otel.SetTextMapPropagator(propagator)
	return closeFunc
}
