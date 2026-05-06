package selector

// Option controls the returned [Selector] from [New]
type Option func(*Selector)

// WithID is used to look up an Object with a specified globally-unique string
// identifier.
func WithID(id string) Option {
	return func(s *Selector) {
		s.id = id
	}
}

// WithDomain is used to look up an Object with a specified Domain.
func WithDomain(domain string) Option {
	return func(s *Selector) {
		s.domain = domain
	}
}

// WithNamespace is used to look up an Object with a specified Namespace.
func WithNamespace(namespace string) Option {
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

// New returns a new Selector given zero or more Option modifiers.
func New(opts ...Option) *Selector {
	o := &Selector{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}
