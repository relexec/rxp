package domain

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/system"
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
func WithSystem(system *system.System) Option {
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
func WithName(name api.DomainName) Option {
	return func(d *Domain) {
		d.name = name
	}
}

// WithRoot sets the Domain's Root.
func WithRoot(root *Domain) Option {
	return func(d *Domain) {
		d.root = root
	}
}

// WithParent sets the Domain's Parent.
func WithParent(parent *Domain) Option {
	return func(d *Domain) {
		d.parent = parent
	}
}
