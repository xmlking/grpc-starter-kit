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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var once sync.Once

// InitTracing before ending program, wait for all enqueued spans to be exported
// expected GOOGLE_CLOUD_PROJECT & GOOGLE_APPLICATION_CREDENTIALS Environment Variable set
func InitTracing(ctx context.Context, cfg *config.Features_Tracing) func() {
	var tp *trace.TracerProvider
	var exporter sdktrace.SpanExporter

	once.Do(func() {
		log.Debug().Interface("TracingConfig", cfg).Msg("Initializing Tracing")

		resources, err := resource.New(ctx,
			// Builtin detectors provide default values and support
			// OTEL_RESOURCE_ATTRIBUTES and OTEL_SERVICE_NAME environment variables
			resource.WithProcess(),                                  // This option configures a set of Detectors that discover process information
			resource.WithAttributes(attribute.String("foo", "bar")), // Or specify resource attributes directly
		)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to initialize resources for tracing exporter")
		}

		target, err := config.ParseTarget(cfg.Target)
		if err != nil {
			log.Fatal().Err(err).Msg("telemetry.tracing config error:")
		}

		switch target {
		case config.GCP:
			projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
			exporter, err = cloudtrace.New(cloudtrace.WithProjectID(projectID))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize google tracing exporter")
			}
			tp = sdktrace.NewTracerProvider(
				// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
				// In a production application, use sdktrace.TraceIDRatioBased/ParentBased/NeverSample with a desired probability.
				// fraction >= 1 means AlwaysSample()
				sdktrace.WithSampler(sdktrace.TraceIDRatioBased(cfg.SamplingFraction)),
				sdktrace.WithResource(resources),
				sdktrace.WithBatcher(exporter),
			)

		case config.STDOUT:
			opts := []stdouttrace.Option{
				stdouttrace.WithPrettyPrint(),
			}
			exporter, err = stdouttrace.New(opts...)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to initialize stdout tracing exporter")
			}
			bsp := sdktrace.NewBatchSpanProcessor(exporter)
			tp = sdktrace.NewTracerProvider(
				// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
				// In a production application, use sdktrace.TraceIDRatioBased/ParentBased/NeverSample with a desired probability.
				// fraction >= 1 means AlwaysSample()
				sdktrace.WithSampler(sdktrace.TraceIDRatioBased(cfg.SamplingFraction)),
				sdktrace.WithResource(resources),
				sdktrace.WithSpanProcessor(bsp),
			)

		default:
			log.Fatal().Msgf("unsupported tracing Target: '%s'", target)
		}

		// Registers trace Provider globally.
		otel.SetTracerProvider(tp)
	})

	return func() {
		exporter.Shutdown(ctx)
		tp.Shutdown(ctx)
	}
}
