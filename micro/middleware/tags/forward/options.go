package forward

import (
	"github.com/xmlking/grpc-starter-kit/shared/constants"
)

type Option func(*Options)

type Options struct {
	tags []string
}

// Default: TraceIDKey
func defaultOptions() Options {
	return Options{
		tags: []string{constants.TraceIDKey},
	}
}

func WithForwardTags(tags ...string) Option {
	return func(args *Options) {
		args.tags = tags
	}
}
