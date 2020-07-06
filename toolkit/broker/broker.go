package broker

import (
    cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Broker interface {
    CeClient() cloudevents.Client
    Options() Options
}

// NewService creates and returns a new Service based on the packages within.
func NewBroker(opts ...Option) Broker {
    return newBroker(opts...)
}
