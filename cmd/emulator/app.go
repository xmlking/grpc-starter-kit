package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type App struct {
	context context.Context
	client  *pubsub.Client
	config  Config
}

type Config struct {
	context          context.Context
	gcpProjectName   string
	subscriptionName string
	topicName        string
	options          []option.ClientOption
}

func newApp(config Config) (*App, error) {
	client, err := pubsub.NewClient(config.context, config.gcpProjectName, config.options...)
	if err != nil {
		return nil, err
	}

	return &App{context: config.context, client: client, config: config}, nil
}

func (app *App) run() {

	log.Println("waiting for messages")

	app.client.Subscription(app.config.subscriptionName).Receive(app.config.context, func(ctx context.Context, message *pubsub.Message) {

		var messageJson map[string]interface{}

		json.Unmarshal(message.Data, &messageJson)

		log.Printf("received message with id: %s and content %v", message.ID, messageJson)

		messageJson["processed_time"] = time.Now()

		result, _ := json.Marshal(messageJson)

		app.client.Topic(app.config.topicName).Publish(ctx, &pubsub.Message{Data: result})

		message.Ack()
	})

	log.Println("stopped waiting for messages")
}
