package registry

import (
	"github.com/sarulabs/di/v2"

	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/service/recorder/subscriber"
	configPB "github.com/xmlking/grpc-starter-kit/shared/proto/config/v1"
)

// Container - provide di Container
type Container struct {
	ctn di.Container
}

// NewContainer - create new Container
func NewContainer(cfg configPB.Configuration) (*Container, error) {
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
			Name:  "transaction-subscriber",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return subscriber.NewTransactionSubscriber(), nil
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
