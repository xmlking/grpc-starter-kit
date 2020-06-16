// e2e, black-box testing
package e2e

import (
    "context"
    "testing"
    "time"

    cloudevents "github.com/cloudevents/sdk-go/v2"
    // cecontext "github.com/cloudevents/sdk-go/v2/context"

    "github.com/rs/zerolog/log"

    emailerPB "github.com/xmlking/grpc-starter-kit/service/emailer/proto/emailer"
    _ "github.com/xmlking/grpc-starter-kit/shared/constants"
    "github.com/xmlking/grpc-starter-kit/shared/eventing"
    _ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func TestEmailSubscriber_Handle_Send_E2E(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping e2e test")
    }

    client  := eventing.NewSourceClient()
    topic   := "mkit.service.emailer"

    // Create an Event.
    event := cloudevents.NewEvent()
    event.SetSource("github.com/xmlking/grpc-starter-kit/service/emailer")
    event.SetType("account.welcome.email")
    event.SetData(cloudevents.ApplicationJSON, &emailerPB.Message{Subject: "Sumo", To: "sumo@demo.com"})

    // Set a target.
    // ctx := cecontext.WithTopic(context.Background(), topic) // for GCP PubSub
    ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")
    ctxWithRetries := cloudevents.ContextWithRetriesLinearBackoff(ctx, 10*time.Millisecond, 3)
    // if you want to send raw like Avro or protobuf
    // ctx = cloudevents.WithEncodingBinary(ctx)

    // Send that Event.
    if result := client.Send(ctxWithRetries, event); !cloudevents.IsACK(result) {
        log.Fatal().Msgf("failed to send, %v", result)
    }

    t.Logf("Successfully published to: %s", topic)
}

func TestEmailSubscriber_Handle_Request_E2E(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping e2e test")
    }

    client  := eventing.NewSourceClient()
    topic   := "mkit.service.emailer"

    // Create an Event.
    event := cloudevents.NewEvent()
    event.SetSource("github.com/xmlking/grpc-starter-kit/service/emailer")
    event.SetType("account.welcome.email")
    event.SetData(cloudevents.ApplicationJSON, &emailerPB.Message{Subject: "Sumo", To: "sumo@demo.com"})

    // Set a target.
    // ctx := cecontext.WithTopic(context.Background(), topic) // for GCP PubSub
    ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")
    ctxWithRetries := cloudevents.ContextWithRetriesLinearBackoff(ctx, 10*time.Millisecond, 3)
    // if you want to send raw like Avro or protobuf
    // ctx = cloudevents.WithEncodingBinary(ctx)

    // Request that Event.
    if resp, res := client.Request(ctxWithRetries, event); !cloudevents.IsACK(res) {
        log.Fatal().Msgf("failed to send, %v", res)
    } else if resp != nil {
        log.Debug().Msg(resp.String())
        log.Debug().Msgf("Got Event Response Context: %+v\n", resp.Context)
    }

    t.Logf("Successfully published to: %s", topic)
}
