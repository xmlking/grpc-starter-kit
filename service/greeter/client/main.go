package main

import (
    "context"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/balancer/roundrobin"

    "google.golang.org/grpc"

    "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
    "github.com/xmlking/grpc-starter-kit/shared/config"
    "github.com/xmlking/grpc-starter-kit/shared/constants"
)

func main() {
    serviceName := constants.GREETER_SERVICE
    cfg := config.GetConfig()

    println(serviceName)
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
        log.Fatal().Msgf("Error when calling SayHello: %s", err)
    }
    log.Printf("Response from server: %s", response.Msg)
}
