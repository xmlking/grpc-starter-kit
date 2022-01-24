package handler

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/trace"

	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
)

var (
	lemonsKey  = attribute.Key("ex.com/lemons")
	anotherKey = attribute.Key("ex.com/another")
)

type greeterHandler struct {
	visitsCounter metric.Int64Counter
	tracer        trace.Tracer
	meter         metric.Meter
}

// NewGreeterHandler returns an instance of `GreeterServiceServer`.
func NewGreeterHandler() greeterv1.GreeterServiceServer {
	visitsCounter := createIntCounter("visit-counter",
		"A counter representing number of times a website is visited.")
	tracer := otel.Tracer("ex.com/basic")
	meter := global.Meter("ex.com/basic")

	return &greeterHandler{
		visitsCounter: visitsCounter,
		tracer:        tracer,
		meter:         meter,
	}
}

// Hello method
func (s *greeterHandler) Hello(ctx context.Context, req *greeterv1.HelloRequest) (*greeterv1.HelloResponse, error) {
	log.Info().Msg("Received Greeter.Hello request")

	// TODO: handle errors
	foo, _ := baggage.NewMember("ex.com.foo", "foo1")
	bar, _ := baggage.NewMember("ex.com.bar", "bar1")
	bag, _ := baggage.New(foo, bar)
	ctx = baggage.ContextWithBaggage(ctx, bag)

	// metrics
	s.visitsCounter.Add(ctx, 1)
	// trace
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("server", "handling this..."))
	defer span.End()

	span.AddEvent("Nice operation!", trace.WithAttributes(attribute.Int("bogons", 100)))
	span.SetAttributes(anotherKey.String("yes"))

	_ = func(ctx context.Context) error {
		var span trace.Span
		ctx, span = s.tracer.Start(ctx, "grpc-starter-kit/play")
		defer span.End()

		span.AddEvent("Nice operation!", trace.WithAttributes(attribute.Int("bogons", 100)))
		span.SetAttributes(anotherKey.String("yes"))

		commonAttributes := []attribute.KeyValue{
			lemonsKey.Int(10),
			attribute.String("A", "1"),
			attribute.String("B", "2"),
			attribute.String("C", "3"),
		}
		histogram := metric.Must(s.meter).NewFloat64Histogram("ex.com.two")

		s.meter.RecordBatch(
			ctx,
			commonAttributes,
			histogram.Measurement(2.0),
		)

		return func(ctx context.Context) error {
			span := trace.SpanFromContext(ctx)
			span.SetAttributes(attribute.String("operation", "Sub operation.."))
			defer span.End()

			span.SetAttributes(lemonsKey.String("five"))
			span.AddEvent("Sub span event")
			histogram.Record(ctx, 1.3, commonAttributes...)
			return nil
		}(ctx)
	}(ctx)

	log.Info().Msgf("visitsCounter: %v", s.visitsCounter)

	return &greeterv1.HelloResponse{Msg: "Hello " + req.Name + " from play"}, nil
}

func createIntCounter(name string, desc string) metric.Int64Counter {
	meter := global.Meter("otel-switch-backend")
	counter := metric.Must(meter).NewInt64Counter(name,
		metric.WithDescription(desc),
	)
	//.Bind(attribute.String("label", "test"))
	return counter
}
