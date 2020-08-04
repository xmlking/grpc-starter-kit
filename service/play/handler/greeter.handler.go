package handler

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/api/trace"

	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
)

type greeterHandler struct {
	visitsCounter metric.BoundInt64Counter
	tracer        trace.Tracer
}

// NewUserHandler returns an instance of `GreeterServiceServer`.
func NewGreeterHandler() greeterv1.GreeterServiceServer {
	visitsCounter := createIntCounter("visit-counter",
		"A counter representing number of times a website is visited.")
	tracer := global.Tracer("ex.com/basic")
	return &greeterHandler{
		visitsCounter: visitsCounter,
		tracer:        tracer,
	}
}

// Hello method
func (s *greeterHandler) Hello(ctx context.Context, req *greeterv1.HelloRequest) (*greeterv1.HelloResponse, error) {
	log.Info().Msg("Received Greeter.Hello request")
	// metrics
	s.visitsCounter.Add(ctx, 1)
	// trace
	s.tracer.WithSpan(context.Background(), "foo",
		func(ctx context.Context) error {
			s.tracer.WithSpan(ctx, "bar",
				func(ctx context.Context) error {
					s.tracer.WithSpan(ctx, "baz",
						func(ctx context.Context) error {
							return nil
						},
					)
					return nil
				},
			)
			return nil
		},
	)

	log.Info().Msgf("visitsCounter: %v", s.visitsCounter)

	return &greeterv1.HelloResponse{Msg: "Hello " + req.Name + " from cmux"}, nil
}

func createIntCounter(name string, desc string) metric.BoundInt64Counter {
	meter := global.Meter("otel-switch-backend")
	counter := metric.Must(meter).NewInt64Counter(name,
		metric.WithDescription(desc),
	).Bind(kv.String("label", "test"))
	return counter
}
