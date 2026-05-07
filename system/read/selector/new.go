package selector

// Option controls the returned [Selector] from [New]
type Option func(*Selector)

// WithUUID is used to look up the System by its globally-unique identifier.
func WithUUID(uuid string) Option {
	return func(s *Selector) {
		s.uuid = uuid
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
