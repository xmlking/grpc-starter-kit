package eventing

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"

	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/shared/config"
	_ "github.com/xmlking/grpc-starter-kit/shared/constants"
)

func NewSourceClient(target string) cloudevents.Client {
	//cfg := config.GetConfig()
	//transport, err := cepubsub.New(context.Background(),
	//    cepubsub.WithProjectID(cfg.gcp.ProjectID),
	//    cepubsub.WithTopicID(cfg.gcp.TopicID),
	//)
	//client, err := cloudevents.NewClient(transport, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())

	// The default client is HTTP.
	// client, err := cloudevents.NewDefaultClient()

	p, err := cloudevents.NewHTTP(cloudevents.WithTarget(target))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create protocol")
	}

	client, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to create client")
	}
	return client
}

func NewSinkClient(target string) cloudevents.Client {
	//cfg := config.GetConfig()
	//transport, err := cepubsub.New(context.Background(),
	//    cepubsub.WithProjectID(cfg.gcp.ProjectID),
	//    cepubsub.WithTopicID(cfg.gcp.TopicID),
	//    cepubsub.WithSubscriptionID(cfg.gcp.SubscriptionID),
	//    you can add more Subscriptions, it the case of multi-receiver
	//    cepubsub.WithSubscriptionID(cfg.gcp.SubscriptionID_2),
	//)
	//client, err := cloudevents.NewClient(transport, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())

	// The default client is HTTP.
	//client, err := cloudevents.NewDefaultClient()

	lis, err := config.GetListener(target)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to create listener for target: %v", target)
	}

	p, err := cloudevents.NewHTTP(cloudevents.WithListener(lis))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create protocol")
	}

	client, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to create client")
	}
	return client
}
