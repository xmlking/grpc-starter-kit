package rpclog

import (
	"regexp"

	"github.com/rs/zerolog/log"
)

type Option func(*Options)

type Options struct {
	ExcludePatterns []*regexp.Regexp
}

// Default: exclude health check method
func defaultOptions() Options {
	excludePatterns, _ := regexp.Compile("/grpc.health.v1.Health/Check")
	return Options{
		ExcludePatterns: []*regexp.Regexp{excludePatterns},
	}
}

func WithExcludeMethods(excludeMethods ...string) Option {
	return func(args *Options) {
		excludePatterns := make([]*regexp.Regexp, len(excludeMethods))
		for i, m := range excludeMethods {
			r, err := regexp.Compile(m)
			if err != nil {
				log.Fatal().Msgf("error: invalid method pattern: %v", err)
			}
			excludePatterns[i] = r
		}
	}
}

func (o *Options) matchMethod(method string) bool {
	for _, r := range o.ExcludePatterns {
		if r.MatchString(method) {
			return true
		}
	}
	return false
}
