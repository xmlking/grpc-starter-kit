package handler

import (
    "context"

    "github.com/rs/zerolog/log"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/baggage"
    "go.opentelemetry.io/otel/metric"
    "go.opentelemetry.io/otel/trace"
    "google.golang.org/genproto/googleapis/api/label"

    "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
)

var (
	fooKey     = attribute.Key("ex.com/foo")
	barKey     = attribute.Key("ex.com/bar")
	lemonsKey  = attribute.Key("ex.com/lemons")
	anotherKey = attribute.Key("ex.com/another")
)

type greeterHandler struct {
	visitsCounter metric.BoundInt64Counter
	tracer        trace.Tracer
}

// NewUserHandler returns an instance of `GreeterServiceServer`.
func NewGreeterHandler() greeterv1.GreeterServiceServer {
	visitsCounter := createIntCounter("visit-counter",
		"A counter representing number of times a website is visited.")
	tracer := otel.Tracer("ex.com/basic")
	return &greeterHandler{
		visitsCounter: visitsCounter,
		tracer:        tracer,
	}
}

// Hello method
func (s *greeterHandler) Hello(ctx context.Context, req *greeterv1.HelloRequest) (*greeterv1.HelloResponse, error) {
	log.Info().Msg("Received Greeter.Hello request")

	ctx = baggage.ContextWithValues(ctx, fooKey.String("foo1"), barKey.String("bar1"))

	// metrics
	s.visitsCounter.Add(ctx, 1)
	// trace
	var span trace.Span
	ctx, span = s.tracer.Start(ctx, "operation")
	defer span.End()

	span.AddEvent("Nice operation!", trace.WithAttributes(label.Int("bogons", 100)))
	span.SetAttributes(anotherKey.String("yes"))

	_ = func(ctx context.Context) error {
		var span trace.Span
		ctx, span = s.tracer.Start(ctx, "operation")
		defer span.End()

		span.AddEvent("Nice operation!", trace.WithAttributes(label.Int("bogons", 100)))
		span.SetAttributes(anotherKey.String("yes"))

		return func(ctx context.Context) error {
			var span trace.Span
			ctx, span = s.tracer.Start(ctx, "Sub operation...")
			defer span.End()

			span.SetAttributes(lemonsKey.String("five"))
			span.AddEvent("Sub span event")

			return nil
		}(ctx)
	}(ctx)

	log.Info().Msgf("visitsCounter: %v", s.visitsCounter)

	return &greeterv1.HelloResponse{Msg: "Hello " + req.Name + " from cmux"}, nil
}

func createIntCounter(name string, desc string) metric.BoundInt64Counter {
	meter := otel.Meter("otel-switch-backend")
	counter := metric.Must(meter).NewInt64Counter(name,
		metric.WithDescription(desc),
	).Bind(label.String("label", "test"))
	return counter
}
