package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
)

const (
	testSub      = "in-sub"
	testSubTopic = "in-topic"
	resultTopic  = "out-topic"
	resultSub    = "out-sub"
)

func TestPubSubProcessing(t *testing.T) {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	app, err := newApp(Config{context: ctx, gcpProjectName: "test", subscriptionName: testSub, topicName: resultTopic, options: []option.ClientOption{option.WithoutAuthentication()}})
	assert.Nil(t, err, "app creation is successfull")
	prepare(app)

	//subscribe first
	go func() {
		app.client.Topic(testSubTopic).Publish(ctx, &pubsub.Message{Data: []byte("{\"greeting\" : \"hello\"}")})
		app.client.Subscription(resultSub).Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			var jsonMessage map[string]interface{}
			json.Unmarshal(msg.Data, &jsonMessage)
			assert.Equal(t, "hello", jsonMessage["greeting"], "greeting field is kept as is")
			assert.NotEmpty(t, jsonMessage["processed_time"], "processed time field is added")
			fmt.Println("finished assertions")
		})
	}()

	app.run()

}

func prepare(app *App) {
	//no error checking, since this is just a demo
	topic, _ := app.client.CreateTopic(app.config.context, testSubTopic)
	app.client.CreateSubscription(app.config.context, testSub, pubsub.SubscriptionConfig{Topic: topic})
	res, _ := app.client.CreateTopic(app.config.context, resultTopic)
	app.client.CreateSubscription(app.config.context, resultSub, pubsub.SubscriptionConfig{Topic: res})
}
