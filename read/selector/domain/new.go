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

// WithName is used to look up an target with a specified DomainName.
func WithName(domainName types.DomainName) Option {
	return func(s *Selector) {
		s.domainName = domainName
	}
}

// WithSystem is used to look up an target with a specified System.
func WithSystem(system types.System) Option {
	return func(s *Selector) {
		s.system = system
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
