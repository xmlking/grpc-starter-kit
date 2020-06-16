package eventing

import (
    cloudevents "github.com/cloudevents/sdk-go/v2"

    "github.com/rs/zerolog/log"

    _ "github.com/xmlking/grpc-starter-kit/shared/constants"
)

func NewSourceClient() cloudevents.Client {
    //cfg := config.GetConfig()
    //transport, err := cepubsub.New(context.Background(),
    //    cepubsub.WithProjectID(cfg.gcp.ProjectID),
    //    cepubsub.WithTopicID(cfg.gcp.TopicID),
    //)
    //client, err := cloudevents.NewClient(transport, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())

    // The default client is HTTP.
    client, err := cloudevents.NewDefaultClient()
    if err != nil {
        log.Fatal().Err(err).Msgf("failed to create client")
    }
    return client
}

func NewSinkClient() cloudevents.Client {
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
    client, err := cloudevents.NewDefaultClient()
    if err != nil {
        log.Fatal().Err(err).Msgf("failed to create client")
    }
    return client
}
