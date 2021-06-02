package registry

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/sarulabs/di/v2"
	broker "github.com/xmlking/toolkit/broker/cloudevents"
	"google.golang.org/grpc"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/database"
	greeterv1 "github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/account/handler"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
)

// Container - provide di Container
type Container struct {
	ctn di.Container
}

// NewContainer - create new Container
func NewContainer(appCtx context.Context, cfg config.Configuration) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "config",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return cfg, nil
			},
		},
		{
			Name:  "user-repository",
			Scope: di.App,
			Build: buildUserRepository,
		},
		{
			Name:  "profile-repository",
			Scope: di.App,
			Build: buildProfileRepository,
		},
		{
			Name:  "translog-publisher",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				bkr := broker.NewBroker(appCtx)
				return bkr.NewPublisher(cfg.Services.Recorder.Endpoint)
			},
		},
		{
			Name:  "email-publisher",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				bkr := broker.NewBroker(appCtx)
				return bkr.NewPublisher(cfg.Services.Emailer.Endpoint)
			},
		},
		{
			Name:  "greeter-connection",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.GetClientConn(cfg.Services.Greeter, nil)
			},
			Close: func(obj interface{}) error {
				return obj.(*grpc.ClientConn).Close()
			},
		},
		{
			Name:  "greeter-client",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				greeterConn := ctn.Get("greeter-connection").(*grpc.ClientConn)
				return greeterv1.NewGreeterServiceClient(greeterConn), nil
			},
		},
		{
			Name:  "user-handler",
			Scope: di.App,
			Build: buildUserHandler,
		},
		{
			Name:  "profile-handler",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				repo := ctn.Get("profile-repository").(repository.ProfileRepository)

				subLogger := log.With().Str("component", "ProfileHandler").Logger()
				return handler.NewProfileHandler(repo, subLogger), nil
			},
		},
		{
			Name:  "database",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return database.InitDatabase(*cfg.Database)
			},
			Close: func(obj interface{}) error {
				return obj.(*ent.Client).Close()
			},
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// Resolve object
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean Container
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

// Delete Container
func (c *Container) Delete() error {
	return c.ctn.Delete()
}

func buildUserRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("database").(*ent.Client)
	return repository.NewUserRepository(db), nil
}

func buildProfileRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("database").(*ent.Client)
	return repository.NewProfileRepository(db), nil
}

func buildUserHandler(ctn di.Container) (interface{}, error) {
	repo := ctn.Get("user-repository").(repository.UserRepository)
	emailPublisher := ctn.Get("email-publisher").(broker.Publisher)
	greeterSrvClient := ctn.Get("greeter-client").(greeterv1.GreeterServiceClient)
	return handler.NewUserHandler(repo, emailPublisher, greeterSrvClient), nil
}
