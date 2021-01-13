package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	_ "github.com/xmlking/grpc-starter-kit/internal/logger"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/rs/zerolog/log"
	appendTags "github.com/xmlking/toolkit/middleware/tags/append"
	"github.com/xmlking/toolkit/util"
	"google.golang.org/grpc"

	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	userv1 "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
)

var (
	cfg = config.GetConfig()
)

func main() {
	log.Debug().Msgf("IsProduction? %t", config.IsProduction())
	//log.Debug().Interface("Dialect", cfg.Database.Dialect).Send()
	//log.Debug().Msg(cfg.Database.Host)
	//log.Debug().Uint32("Port", cfg.Database.Port).Send()
	//log.Debug().Uint64("FlushInterval", cfg.Features.Tracing.FlushInterval).Send()
	//log.Debug().Msgf("cfg is %+v", cfg)

	username := flag.String("username", "sumo", "username of user to be create")
	email := flag.String("email", "sumo@demo.com", "email of user to be create")
	limit := flag.Uint64("limit", 10, "Limit number of results")
	flag.Parse()

	log.Debug().Str("username", *username).Str("email", *email).Uint64("limit", *limit).Msg("Flags Using:")

	pairs := []string{constants.FromServiceKey, constants.ACCOUNT_CLIENT}
	for key, val := range cfg.Services.Account.Metadata {
		pairs = append(pairs, key, val)
	}
	var ucInterceptors = []grpc.UnaryClientInterceptor{
		appendTags.UnaryClientInterceptor(appendTags.WithTraceID(), appendTags.WithPairs(pairs...)),
	}
	conn, err := config.GetClientConn(cfg.Services.Account, ucInterceptors)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed connect to: %s", cfg.Services.Account.Endpoint)
	}

	userClient := userv1.NewUserServiceClient(conn)

	suffix := util.RandomStringLower(5)
	if rsp, err := userClient.Create(context.Background(), &userv1.CreateRequest{
		Username:  &wrappers.StringValue{Value: "u_" + suffix},
		FirstName: &wrappers.StringValue{Value: "f_" + suffix},
		LastName:  &wrappers.StringValue{Value: "l_" + suffix},
		Email:     &wrappers.StringValue{Value: fmt.Sprintf("e_%s@demo.com", suffix)},
	}); err != nil {
		log.Error().Err(err).Send()
		os.Exit(1)
	} else {
		log.Info().Interface("createRsp", rsp).Send()
	}

	getUserList(userClient, uint32(*limit))
}

func getUserList(us userv1.UserServiceClient, limit uint32) {
	if rsp, err := us.List(context.Background(), &userv1.ListRequest{Limit: &wrappers.UInt32Value{Value: limit}}); err != nil {
		log.Fatal().Err(err).Msg("Unable to List Users")
	} else {
		log.Info().Interface("listRsp", rsp).Send()
	}
}
