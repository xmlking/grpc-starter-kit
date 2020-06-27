package append

import (
	"google.golang.org/grpc/metadata"
)

type Option func(*Options)

type Options struct {
	metadata       metadata.MD
	metadataFunc   MetadataFunc
	overwrite      bool
	traceIdEnabled bool
}

type MetadataFunc func() metadata.MD

// Default: traceIdEnabled = false and overwrite = true
func defaultOptions() Options {
	return Options{
		overwrite:      true,
		traceIdEnabled: false,
		metadataFunc:   func() metadata.MD { return metadata.MD{} },
	}
}

// WithPairs pass kv: "k1", "v1", "k1", "v2", "k2", "v3"
func WithPairs(kv ...string) Option {
	return func(args *Options) {
		args.metadata = metadata.Pairs(kv...)
	}
}

func WithMetadataFunc(f MetadataFunc) Option {
	return func(args *Options) {
		args.metadataFunc = f
	}
}

func WithTraceID() Option {
	return func(args *Options) {
		args.traceIdEnabled = true
	}
}

func WithOverwrite(overwrite bool) Option {
	return func(args *Options) {
		args.overwrite = overwrite
	}
}

// getMetadata return joined metadata
func (o *Options) getMetadata() metadata.MD {
	return metadata.Join(o.metadata, o.metadataFunc())
}
