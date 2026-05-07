package domain

import (
	"github.com/relexec/rxp/types"
)

// Option modifies a Domain returned from New.
type Option func(*Domain)

// New returns a new [Domain].
func New(opts ...Option) *Domain {
	m := &Domain{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithSystem sets the Domain's System.
func WithSystem(system string) Option {
	return func(d *Domain) {
		d.system = system
	}
}

// WithName sets the Domain's Name.
func WithName(name types.DomainName) Option {
	return func(d *Domain) {
		d.name = name
	}
}
