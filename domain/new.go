package domain

import (
	"github.com/relexec/rxp/api"
)

// Option modifies a Domain returned from New.
type Option func(*api.Domain)

// New returns a new [Domain].
func New(opts ...Option) *api.Domain {
	d := &api.Domain{}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

// WithSystem sets the Domain's System.
func WithSystem(system *api.System) Option {
	return func(d *api.Domain) {
		d.SetSystem(system)
	}
}

// WithUUID sets the Domain's UUID.
func WithUUID(uuid string) Option {
	return func(d *api.Domain) {
		d.SetUUID(uuid)
	}
}

// WithName sets the Domain's Name.
func WithName(name api.DomainName) Option {
	return func(d *api.Domain) {
		d.SetName(name)
	}
}

// WithRoot sets the Domain's Root.
func WithRoot(root *api.Domain) Option {
	return func(d *api.Domain) {
		d.SetRoot(root)
	}
}

// WithParent sets the Domain's Parent.
func WithParent(parent *api.Domain) Option {
	return func(d *api.Domain) {
		d.SetParent(parent)
	}
}
