package broker

import (
    //cepubsub "github.com/cloudevents/sdk-go/protocol/pubsub/v2"
    cloudevents "github.com/cloudevents/sdk-go/v2"

    "github.com/rs/zerolog/log"

    "github.com/xmlking/grpc-starter-kit/shared/config"
    _ "github.com/xmlking/grpc-starter-kit/shared/constants"
)

var (
    DefaultClient cloudevents.Client
)

func init() {
    DefaultClient = newClient()
}

func newClient() cloudevents.Client {
    cfg := config.GetConfig()
    println(cfg.String())
    //t, err := cepubsub.New(context.Background(),
    //    cepubsub.WithProjectID("env.ProjectID"),
    //    cepubsub.WithTopicID("env.TopicID"),
    //    cepubsub.WithSubscriptionID("env.SubscriptionID"))
    //c, err := cloudevents.NewClient(t, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())

    // The default client is HTTP.
    client, err := cloudevents.NewDefaultClient()
    if err != nil {
        log.Fatal().Err(err).Msgf("failed to create client")
    }
    return client
}
