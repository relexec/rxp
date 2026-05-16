package domain

import (
	"github.com/relexec/rxp/types"
)

// Option modifies a Domain returned from New.
type Option func(*Domain)

// New returns a new [Domain].
func New(opts ...Option) *Domain {
	d := &Domain{}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

// WithSystem sets the Domain's System.
func WithSystem(system types.System) Option {
	return func(d *Domain) {
		d.system = system
	}
}

// WithUUID sets the Domain's UUID.
func WithUUID(uuid string) Option {
	return func(d *Domain) {
		d.uuid = uuid
	}
}

// WithName sets the Domain's Name.
func WithName(name types.DomainName) Option {
	return func(d *Domain) {
		d.name = name
	}
}
