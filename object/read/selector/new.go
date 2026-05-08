package selector

import "github.com/relexec/rxp/types"

// Option controls the returned [Selector] from [New]
type Option func(*Selector)

// WithSystem is used to look up an Object in a specific System. If no system
// identifier is specified, the host system responding to the ObjectRead
// request is used.
func WithSystem(system types.System) Option {
	return func(s *Selector) {
		s.system = system
	}
}

// WithKindVersion specifies the type and version of Object to look up.
func WithKindVersion(kindVersion types.KindVersion) Option {
	return func(s *Selector) {
		s.kindVersion = kindVersion
	}
}

// WithUUID is used to look up an Object with a specified globally-unique
// string identifier.
func WithUUID(uuid string) Option {
	return func(s *Selector) {
		s.uuid = uuid
	}
}

// WithDomain is used to look up an Object with a specified Domain.
func WithDomain(domain types.Domain) Option {
	return func(s *Selector) {
		s.domain = domain
	}
}

// WithNamespace is used to look up an Object with a specified Namespace.
func WithNamespace(namespace types.Namespace) Option {
	return func(s *Selector) {
		s.namespace = namespace
	}
}

// WithName is used to look up an Object with a specified Name.
func WithName(name string) Option {
	return func(s *Selector) {
		s.name = name
	}
}

// WithGeneration is used to look up an Object with a Spec having the supplied
// Generation. If WithGeneration is not used or a value of 0 is supplied to
// WithGeneration, the Object Spec's latest generation is read.
func WithGeneration(generation types.Generation) Option {
	return func(s *Selector) {
		s.generation = generation
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
