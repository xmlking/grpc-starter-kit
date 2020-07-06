package broker

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/shared/eventing"
)

const (
	DefaultName = "mkit.broker.default"
)

type broker struct {
	opts     Options
	ceClient cloudevents.Client
}

func newBroker(opts ...Option) Broker {
	// Default Options
	options := Options{
		Name: DefaultName,
	}
	b := broker{opts: options}
	b.ApplyOptions(opts...)

	if b.opts.Endpoint != "" {
		b.ceClient = eventing.NewSinkClient(b.opts.Endpoint)
	} else {
		log.Warn().Msg("no broker Endpoint provided. creating default broker")
		if ceClient, err := cloudevents.NewDefaultClient(); err != nil {
			b.ceClient = ceClient
		} else {
			log.Fatal().Err(err).Msg("unable to create default broker")
		}
	}
	return &b
}

func (b *broker) ApplyOptions(opts ...Option) {
	// process options
	for _, o := range opts {
		o(&b.opts)
	}
}

func (b *broker) CeClient() cloudevents.Client {
	return b.ceClient
}

func (b *broker) Options() Options {
	return b.opts
}
