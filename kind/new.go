package kind

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/system"
)

// Option modifies a Kind returned from New.
type Option func(*Kind)

// New returns a new [Kind].
func New(opts ...Option) *Kind {
	m := &Kind{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithSystem sets the Kind's Systek.
func WithSystem(system *system.System) Option {
	return func(k *Kind) {
		k.system = system
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
