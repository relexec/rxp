package kind

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/system"
)

// Option modifies a Kind returned from New.
type Option func(*Kind)

// New returns a new [Kind].
func New(opts ...Option) *Kind {
	k := &Kind{}
	for _, opt := range opts {
		opt(k)
	}
	return k
}

// WithSystem sets the Kind's System.
func WithSystem(system *system.System) Option {
	return func(k *Kind) {
		k.system = system
	}
}

// WithUUID sets the Kind's UUID.
func WithUUID(uuid string) Option {
	return func(k *Kind) {
		k.uuid = uuid
	}
}

// WithName sets the Kind's name.
func WithName(name api.KindName) Option {
	return func(k *Kind) {
		k.name = name
	}
}

// WithScope sets the Kind's Scope.
func WithScope(scope api.Scope) Option {
	return func(k *Kind) {
		k.scope = scope
	}
}
