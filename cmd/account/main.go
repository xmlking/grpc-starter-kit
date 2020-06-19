package main

import (
    "context"
    "flag"

    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
    "google.golang.org/grpc/metadata"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/balancer/roundrobin"

    "github.com/golang/protobuf/ptypes/wrappers"

    "github.com/xmlking/grpc-starter-kit/micro/middleware/rpclog"
    userv1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
    "github.com/xmlking/grpc-starter-kit/shared/config"
    _ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

var (
    cfg = config.GetConfig()
)

func main() {
    log.Debug().Msgf("IsProduction? %v", config.IsProduction())
    //log.Debug().Interface("Dialect", cfg.Database.Dialect).Send()
    //log.Debug().Msg(cfg.Database.Host)
    //log.Debug().Uint32("Port", cfg.Database.Port).Send()
    //log.Debug().Uint64("FlushInterval", cfg.Features.Tracing.FlushInterval).Send()
    //log.Debug().Msgf("cfg is %v", cfg)

    username := flag.String("username", "sumo", "username of user to be create")
    email := flag.String("email", "sumo@demo.com", "email of user to be create")
    limit := flag.Uint64("limit", 10, "Limit number of results")
    flag.Parse()

    log.Debug().Str("username", *username).Str("email", *email).Uint64("limit", *limit).Msg("Flags Using:")

    conn, err := grpc.Dial(
        cfg.Services.Account.Endpoint, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name),
        grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
            rpclog.UnaryClientInterceptor(),
        )),
    )
    if err != nil {
        log.Fatal().Msgf("did not connect: %s", err)
    }

    userClient := userv1.NewUserServiceClient(conn)

    // Sending metadata - client side
    //md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
    //ctx := metadata.NewOutgoingContext(context.Background(), md)
    // create a new context with some metadata - (Optional) Just for demonstration
    ctx := metadata.AppendToOutgoingContext(context.Background(), "X-User-Id", "john", "X-From-Id", "script")

    rsp, err := userClient.Create(ctx, &userv1.CreateRequest{
        Username:  &wrappers.StringValue{Value: "sumo"},
        FirstName: &wrappers.StringValue{Value: "sumo"},
        LastName:  &wrappers.StringValue{Value: "demo"},
        Email:     &wrappers.StringValue{Value: "sumo@demo.com"},
    })

    log.Info().Interface("createRsp", rsp).Send()

    getUserList(userClient, uint32(*limit))
}

func getUserList(us userv1.UserServiceClient, limit uint32) {
    if rsp, err := us.List(context.Background(), &userv1.ListRequest{Limit: &wrappers.UInt32Value{Value: limit}}); err != nil {
        log.Fatal().Err(err).Msg("Unable to List Users")
    } else {
        log.Info().Interface("listRsp", rsp).Send()
    }
}
