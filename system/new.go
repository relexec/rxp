package system

// Option modifies a System returned from New.
type Option func(*System)

// New returns a new [System].
func New(opts ...Option) *System {
	s := &System{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithUUID sets the System's globally-unique identifier.
func WithUUID(uuid string) Option {
	return func(s *System) {
		s.uuid = uuid
	}
}

// WithName sets the System's optional human-readable name.
func WithName(name string) Option {
	return func(s *System) {
		s.name = name
	}
}
