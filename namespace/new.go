package namespace

import (
	"github.com/relexec/rxp/types"
)

// Option modifies a Namespace returned from New.
type Option func(*Namespace)

// New returns a new [Namespace].
func New(opts ...Option) *Namespace {
	n := &Namespace{}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// WithDomain sets the Namespace's Domain.
func WithDomain(domain types.Domain) Option {
	return func(n *Namespace) {
		n.domain = domain
	}
}

// WithName sets the Namespace's Name.
func WithName(name types.NamespaceName) Option {
	return func(n *Namespace) {
		n.name = name
	}
}
