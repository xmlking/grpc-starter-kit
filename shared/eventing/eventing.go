package eventing

import (
    "encoding/json"
    "net/http"

    //cepubsub "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
    cloudevents "github.com/cloudevents/sdk-go/v2"
    "github.com/rs/zerolog/log"

    _ "github.com/xmlking/grpc-starter-kit/shared/constants"
    "github.com/xmlking/grpc-starter-kit/toolkit/util/endpoint"
)

func NewSourceClient(target string) cloudevents.Client {
    //cfg := config.GetConfig()
    //ctx := context.Background()
    //tOpts := []cepubsub.Option{
    //    cepubsub.WithProjectID(cfg.gcp.ProjectID),
    //    cepubsub.WithTopicID(cfg.gcp.TopicID),
    //    cepubsub.WithSubscriptionAndTopicID(cfg.gcp.SubscriptionID, cfg.gcp.TopicID),
    //}
    //transport, err := cepubsub.New(ctx, tOpts...)
    //if err != nil {
    //    // TODO
    //}
    //client, err := cloudevents.NewClient(transport,  cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
    //if err != nil {
    //    // TODO
    //}

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

    lis, err := endpoint.GetListener(target)
    if err != nil {
        log.Fatal().Err(err).Msgf("failed to create listener for target: %v", target)
    }

    p, err := cloudevents.NewHTTP(cloudevents.WithListener(lis),
        cloudevents.WithGetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // health check handler
            json.NewEncoder(w).Encode(map[string]bool{"ok": true})
        }))
    if err != nil {
        log.Fatal().Err(err).Msg("failed to create protocol")
    }

    client, err := cloudevents.NewClient(p)
    if err != nil {
        log.Fatal().Err(err).Msgf("failed to create client")
    }
    return client
}
