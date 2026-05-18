package selector

import "github.com/relexec/rxp/types"

// Option controls the returned [Selector] from [New]
type Option func(*Selector)

// WithUUID is used to look up an target with a specified globally-unique
// string identifier.
func WithUUID(uuid string) Option {
	return func(s *Selector) {
		s.uuid = uuid
	}
}

// WithName is used to look up an target with a specified NamespaceName.
func WithName(namespaceName types.NamespaceName) Option {
	return func(s *Selector) {
		s.namespaceName = namespaceName
	}
}

// WithDomain is used to look up an target with a specified Domain.
func WithDomain(domain types.Domain) Option {
	return func(s *Selector) {
		s.domain = domain
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
