package broker

import (
	"context"

	"cloud.google.com/go/pubsub"
)

// Broker is an interface used for asynchronous messaging.
type Broker interface {
	Options() Options
	Shutdown() error
	Publish(topic string, m *pubsub.Message, opts ...PublishOption) error
	Subscribe(topic string, h Handler, opts ...SubscribeOption) error
	String() string
}

// Handler is used to process messages via a subscription of a topic.
// The handler is passed a publication interface which contains the
// message and optional Ack method to acknowledge receipt of the message.
type Handler func(context.Context, *pubsub.Message) error

type Message struct {
	Header map[string]string
	Body   []byte
}

// Event is given to a subscription handler for processing
type Event interface {
	Topic() string
	Message() *Message
	Ack() error
	Error() error
}

// Subscriber is a convenience return type for the Subscribe method
type Subscriber interface {
	Options() SubscribeOptions
	Topic() string
	Unsubscribe() error
}

var DefaultBroker Broker

func Shutdown() error {
	return DefaultBroker.Shutdown()
}

func Publish(topic string, msg *pubsub.Message, opts ...PublishOption) error {
	return DefaultBroker.Publish(topic, msg, opts...)
}

func Subscribe(topic string, handler Handler, opts ...SubscribeOption) error {
	return DefaultBroker.Subscribe(topic, handler, opts...)
}

func String() string {
	return DefaultBroker.String()
}
