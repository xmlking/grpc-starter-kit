// e2e, black-box testing
package e2e

import (
	"context"
	"testing"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	// cecontext "github.com/cloudevents/sdk-go/v2/context"

	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/mkit/service/emailer/v1"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	_ "github.com/xmlking/grpc-starter-kit/shared/constants"
	"github.com/xmlking/grpc-starter-kit/shared/eventing"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

func TestEmailSubscriber_Handle_Send_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}

	cfg := config.GetConfig()
	topic := cfg.Services.Emailer.Endpoint
	client := eventing.NewSourceClient(topic)

	// Create an Event.
	event := cloudevents.NewEvent()
	event.SetSource("github.com/xmlking/grpc-starter-kit/service/emailer")
	event.SetType("account.welcome.email")
	_ = event.SetData(cloudevents.ApplicationJSON, &emailerv1.Message{Subject: "Sumo", To: "sumo@demo.com"})

	// Set a target.
	// ctx := cecontext.WithTopic(context.Background(), topic) // for GCP PubSub
	//ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8082/")
	ctxWithRetries := cloudevents.ContextWithRetriesLinearBackoff(context.Background(), 10*time.Millisecond, 3)
	// if you want to send raw like Avro or protobuf
	// ctx = cloudevents.WithEncodingBinary(ctx)

	// Send that Event.
	if result := client.Send(ctxWithRetries, event); !cloudevents.IsACK(result) {
		log.Fatal().Msgf("failed to send, %+v", result)
	}

	t.Logf("Successfully published to: %s", topic)
}

func TestEmailSubscriber_Handle_Request_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}

	cfg := config.GetConfig()
	topic := cfg.Services.Emailer.Endpoint
	client := eventing.NewSourceClient(topic)

	// Create an Event.
	event := cloudevents.NewEvent()
	event.SetSource("github.com/xmlking/grpc-starter-kit/service/emailer")
	event.SetType("account.welcome.email")
	_ = event.SetData(cloudevents.ApplicationJSON, &emailerv1.Message{Subject: "Sumo", To: "sumo@demo.com"})

	// Set a target.
	// ctx := cecontext.WithTopic(context.Background(), topic) // for GCP PubSub
	ctxWithRetries := cloudevents.ContextWithRetriesLinearBackoff(context.Background(), 10*time.Millisecond, 3)
	// if you want to send raw like Avro or protobuf
	// ctx = cloudevents.WithEncodingBinary(ctx)

	// Request that Event.
	if resp, res := client.Request(ctxWithRetries, event); !cloudevents.IsACK(res) {
		log.Fatal().Msgf("failed to send, %+v", res)
	} else if resp != nil {
		log.Debug().Msg(resp.String())
		log.Debug().Msgf("Got Event Response Context: %+v\n", resp.Context)
	}

	t.Logf("Successfully published to: %s", topic)
}
