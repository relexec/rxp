package namespace

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
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
func WithDomain(domain *domain.Domain) Option {
	return func(n *Namespace) {
		n.domain = domain
	}
}

// WithUUID sets the Namespace's UUID.
func WithUUID(uuid string) Option {
	return func(n *Namespace) {
		n.uuid = uuid
	}
}

// WithName sets the Namespace's Name.
func WithName(name api.NamespaceName) Option {
	return func(n *Namespace) {
		n.name = name
	}
}
