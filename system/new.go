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

// WithTag sets an optional string tag for the System. Note this is not called
// "name" because a Name in rxp has a specific semantic meaning that reflects
// the uniqueness constraint its value. Tags have no such uniqueness
// constraint.
func WithTag(tag string) Option {
	return func(s *System) {
		s.tag = tag
	}
}
