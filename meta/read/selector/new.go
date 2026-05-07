package selector

import "github.com/relexec/rxp/types"

// Option controls the returned [Selector] from [New]
type Option func(*Selector)

// WithSystem is used to look up an Meta in a specific rxp system. If no system
// identifier is specified, the host system responding to the MetaRead request
// is used.
func WithSystem(system types.System) Option {
	return func(s *Selector) {
		s.system = system
	}
}

// WithKindVersion is used to look up a Meta with a specified KindVersion.
func WithKindVersion(kv types.KindVersion) Option {
	return func(s *Selector) {
		s.kindVersion = kv
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
