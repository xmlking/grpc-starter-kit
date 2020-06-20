package handler

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
)

type greeterHandler struct {
}

// NewUserHandler returns an instance of `GreeterServiceServer`.
func NewGreeterHandler() greeterv1.GreeterServiceServer {
	return &greeterHandler{}
}

// Hello method
func (s *greeterHandler) Hello(ctx context.Context, req *greeterv1.HelloRequest) (*greeterv1.HelloResponse, error) {
	log.Info().Msg("Received Greeter.Hello request")
	return &greeterv1.HelloResponse{Msg: "Hello " + req.Name + " from cmux"}, nil
}
