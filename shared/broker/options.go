package broker

import (
	"context"
	"time"

	"google.golang.org/api/option"
)

// https://github.com/cloudevents/sdk-go/blob/master/protocol/pubsub/v2/options.go

type Options struct {

	// Handler executed when error happens in broker message
	// processing
	ErrorHandler Handler

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type PublishOptions struct {
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type SubscribeOptions struct {
	// AutoAck defaults to true. When a handler returns
	// with a nil error the message is acked.
	AutoAck bool
	// Subscribers with the same queue name
	// will create a shared subscription where each
	// receives a subset of messages.
	Queue string

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type Option func(*Options)

type PublishOption func(*PublishOptions)

// PublishContext set context
func PublishContext(ctx context.Context) PublishOption {
	return func(o *PublishOptions) {
		o.Context = ctx
	}
}

type SubscribeOption func(*SubscribeOptions)

func NewSubscribeOptions(opts ...SubscribeOption) SubscribeOptions {
	opt := SubscribeOptions{
		AutoAck: true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// DisableAutoAck will disable auto acking of messages
// after they have been handled.
func DisableAutoAck() SubscribeOption {
	return func(o *SubscribeOptions) {
		o.AutoAck = false
	}
}

// ErrorHandler will catch all broker errors that cant be handled
// in normal way, for example Codec errors
func ErrorHandler(h Handler) Option {
	return func(o *Options) {
		o.ErrorHandler = h
	}
}

// Queue sets the name of the queue to share messages on
func Queue(name string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Queue = name
	}
}

// SubscribeContext set context
func SubscribeContext(ctx context.Context) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Context = ctx
	}
}

type clientOptionKey struct{}

type projectIDKey struct{}

type maxOutstandingMessagesKey struct{}

type maxExtensionKey struct{}

type createSubscription struct{}

type deleteSubscription struct{}

// ClientOption is a broker Option which allows google pubsub client options to be
// set for the client
func ClientOption(c ...option.ClientOption) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, clientOptionKey{}, c)
	}
}

// ProjectID provides an option which sets the google project id
func ProjectID(id string) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, projectIDKey{}, id)
	}
}

// CreateSubscription prevents the creation of the subscription if it not exists
func CreateSubscription(b bool) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, createSubscription{}, b)
	}
}

// DeleteSubscription prevents the deletion of the subscription if it not exists
func DeleteSubscription(b bool) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, deleteSubscription{}, b)
	}
}

// MaxOutstandingMessages sets the maximum number of unprocessed messages
// (unacknowledged but not yet expired) to receive.
func MaxOutstandingMessages(max int) SubscribeOption {
	return func(o *SubscribeOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, maxOutstandingMessagesKey{}, max)
	}
}

// MaxExtension is the maximum period for which the Subscription should
// automatically extend the ack deadline for each message.
func MaxExtension(d time.Duration) SubscribeOption {
	return func(o *SubscribeOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, maxExtensionKey{}, d)
	}
}
