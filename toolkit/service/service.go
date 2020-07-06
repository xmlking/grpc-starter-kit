package service

import (
    "google.golang.org/grpc"
)

type Service interface {
    Server() *grpc.Server
    Options() Options
    ApplyOptions(opts ...Option) // TODO: no use, make private ?
    AddSubscriber(fn interface{})
    // Run the service
    Shutdown() error
    // Run the service
    Start() error
    //Config Interface
}

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...Option) Service {
    return newService(opts...)
}
