package system

import "github.com/relexec/rxp/api"

// Option modifies a System returned from New.
type Option func(*api.System)

// New returns a new [System].
func New(opts ...Option) *api.System {
	s := &api.System{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithUUID sets the System's globally-unique identifier.
func WithUUID(uuid string) Option {
	return func(s *api.System) {
		s.SetUUID(uuid)
	}
}

// WithTag sets an optional string tag for the System. Note this is not called
// "name" because a Name in rxp has a specific semantic meaning that reflects
// the uniqueness constraint its value. Tags have no such uniqueness
// constraint.
func WithTag(tag string) Option {
	return func(s *api.System) {
		s.SetTag(tag)
	}
}
