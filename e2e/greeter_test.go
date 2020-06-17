// e2e, black-box testing
package e2e

import (
    "context"
    "testing"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/balancer/roundrobin"

    "google.golang.org/grpc"

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

    var conn *grpc.ClientConn

    conn, err := grpc.Dial(cfg.Services.Greeter.Endpoint, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
    if err != nil {
        log.Fatal().Msgf("did not connect: %s", err)
    }
    defer conn.Close()
    println(conn.Target())
    c := greeterv1.NewGreeterServiceClient(conn)
    response, err := c.Hello(context.Background(), &greeterv1.HelloRequest{Name: "foo"})
    if err != nil {
        t.Fatalf("Error when calling service: (%s), method: (Hello): %v", serviceName, err)
    }
    log.Printf("Response from server: %s", response.Msg)
}
