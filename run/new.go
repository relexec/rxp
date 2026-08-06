package run

import (
	"time"

	"github.com/relexec/rxp/api"
)

// Option modifies a Run returned from New.
type Option func(*api.Run)

// New returns a new [Run].
func New(opts ...Option) *api.Run {
	r := &api.Run{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// WithRequest sets the Run's Request.
func WithRequest(req api.RunRequest) Option {
	return func(r *api.Run) {
		r.SetRequest(req)
	}
}

// WithRoot sets the Run's root Run UUID. A root Run is the outermost execution
// of a piece of work.
func WithRoot(root string) Option {
	return func(r *api.Run) {
		r.SetRoot(root)
	}
}

// WithParent sets the Run that spawned this Run.
func WithParent(parent *api.Run) Option {
	return func(r *api.Run) {
		r.SetParent(parent)
	}
}

// WithScheduledOn sets the Run's scheduled on nano timestamp.
func WithScheduledOn(ts time.Time) Option {
	return func(r *api.Run) {
		r.SetScheduledOn(ts)
	}
}
