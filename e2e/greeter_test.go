// e2e, black-box testing
package e2e

import (
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"

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

	conn, err := config.GetClientConn(cfg.Services.Greeter)
	defer conn.Close()

	greeterClient := greeterv1.NewGreeterServiceClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), constants.TraceIDKey, "23456", constants.FromServiceKey, "e2e-greeter-test-client")
	response, err := greeterClient.Hello(ctx, &greeterv1.HelloRequest{Name: "foo"})
	if err != nil {
		t.Fatalf("Error when calling service: (%s), method: (Hello): %s", serviceName, err)
	}
	log.Printf("Response from server: %s", response.Msg)
}
