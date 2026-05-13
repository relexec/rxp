package kind

import (
	"github.com/relexec/rxp/types"
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
func WithSystem(system types.System) Option {
	return func(k *Kind) {
		k.system = system
	}
}

// WithName sets the Kind's name.
func WithName(name types.KindName) Option {
	return func(k *Kind) {
		k.name = name
	}
}

// WithNamescope sets the Kind's Namescope.
func WithNamescope(namescope types.Namescope) Option {
	return func(k *Kind) {
		k.namescope = namescope
	}
}
