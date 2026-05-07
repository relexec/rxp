package selector

import "github.com/relexec/rxp/types"

// Option controls the returned [Selector] from [New]
type Option func(*Selector)

// WithSystem is used to look up an Domain in a specific System. If no system
// identifier is specified, the host system responding to the DomainRead
// request is used.
func WithSystem(system types.System) Option {
	return func(s *Selector) {
		s.system = system
	}
}

// WithName is used to look up an Domain with a specified Name.
func WithName(name string) Option {
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
