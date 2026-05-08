package selector

import "github.com/relexec/rxp/types"

// Option controls the returned [Selector] from [New]
type Option func(*Selector)

// WithDomain is used to look up an Namespace with a specified Domain.
func WithDomain(domain types.Domain) Option {
	return func(s *Selector) {
		s.domain = domain
	}
}

// WithName is used to look up an Namespace with a specified Name.
func WithName(name types.NamespaceName) Option {
	return func(s *Selector) {
		s.name = name
	}
}

// New returns a new Selector given zero or more Option modifiers.
func New(opts ...Option) Selector {
	s := Selector{}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}
