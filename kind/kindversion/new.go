package kindversion

import (
	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/kind/kindversion/schema"
)

// Option modifies a api.KindVersion returned from New.
type Option func(*api.KindVersion)

// New returns a new [api.KindVersion].
func New(opts ...Option) *api.KindVersion {
	kv := &api.KindVersion{}
	for _, opt := range opts {
		opt(kv)
	}
	return kv
}

// WithSystem sets the api.KindVersion's System.
func WithSystem(system *api.System) Option {
	return func(kv *api.KindVersion) {
		kv.SetSystem(system)
	}
}

// WithKind sets the api.KindVersion's Kind.
func WithKind(k *api.Kind) Option {
	return func(kv *api.KindVersion) {
		kv.SetKind(k)
	}
}

// WithVersion sets the api.KindVersion's Version.
func WithVersion(ver semver.Version) Option {
	return func(kv *api.KindVersion) {
		kv.SetVersion(ver)
	}
}

// WithSchema sets the api.KindVersion's Schema.
func WithSchema(schema *schema.Schema) Option {
	return func(kv *api.KindVersion) {
		kv.SetSchema(schema)
	}
}
