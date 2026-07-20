package kind

import (
	"github.com/relexec/rxp/api"
)

// Option modifies a api.Kind returned from New.
type Option func(*api.Kind)

// New returns a new [api.Kind].
func New(opts ...Option) *api.Kind {
	k := &api.Kind{}
	for _, opt := range opts {
		opt(k)
	}
	return k
}

// WithSystem sets the api.Kind's System.
func WithSystem(system *api.System) Option {
	return func(k *api.Kind) {
		k.SetSystem(system)
	}
}

// WithUUID sets the api.Kind's UUID.
func WithUUID(uuid string) Option {
	return func(k *api.Kind) {
		k.SetUUID(uuid)
	}
}

// WithName sets the api.Kind's name.
func WithName(name api.KindName) Option {
	return func(k *api.Kind) {
		k.SetName(name)
	}
}

// WithScope sets the api.Kind's Scope.
func WithScope(scope api.Scope) Option {
	return func(k *api.Kind) {
		k.SetScope(scope)
	}
}
