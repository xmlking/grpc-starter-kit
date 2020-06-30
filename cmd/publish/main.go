package main

import (
	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/shared/broker"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

var (
	cfg = config.GetConfig()
)

func main() {
	// broker.DefaultBroker = broker.NewBroker(broker.ProjectID("my-project-id")); // use cfg.pubsub.ProjectID
	broker.DefaultBroker = broker.NewBroker()

	msg := pubsub.Message{
		ID:         uuid.New().String(),
		Data:       []byte("ABC€"),
		Attributes: map[string]string{"sumo": "demo"},
	}

	if err := broker.Publish("ingestion-in-dev", &msg); err != nil {
		log.Error().Err(err).Send()
	}
}
