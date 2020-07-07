// e2e, black-box testing
package e2e

import (
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	appendTags "github.com/xmlking/toolkit/middleware/tags/append"

	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
)

func TestGreeter_Hello_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}

	serviceName := constants.GREETER_SERVICE
	cfg := config.GetConfig()

	var ucInterceptors = []grpc.UnaryClientInterceptor{
		appendTags.UnaryClientInterceptor(appendTags.WithTraceID(), appendTags.WithPairs(constants.FromServiceKey, "e2e-greeter-test-client")),
	}
	conn, err := config.GetClientConn(cfg.Services.Greeter, ucInterceptors)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed connect to: %s", cfg.Services.Greeter.Endpoint)
	}
	defer conn.Close()

	greeterClient := greeterv1.NewGreeterServiceClient(conn)
	response, err := greeterClient.Hello(context.Background(), &greeterv1.HelloRequest{Name: "foo"})
	if err != nil {
		t.Fatalf("Error when calling service: (%s), method: (Hello): %s", serviceName, err)
	}
	log.Printf("Response from server: %s", response.Msg)
}
