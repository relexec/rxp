package run

import (
	"time"

	"github.com/relexec/rxp/run/request"
)

// Option modifies a Run returned from New.
type Option func(*Run)

// New returns a new [Run].
func New(opts ...Option) *Run {
	r := &Run{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// WithRequest sets the Run's Request.
func WithRequest(req request.Request) Option {
	return func(r *Run) {
		r.req = req
	}
}

// WithRoot sets the Run's root Run UUID. A root Run is the outermost execution
// of a piece of work.
func WithRoot(root string) Option {
	return func(r *Run) {
		r.root = root
	}
}

// WithParent sets the Run that spawned this Run.
func WithParent(parent *Run) Option {
	return func(r *Run) {
		r.parent = parent
	}
}

// WithScheduledOn sets the Run's scheduled on nano timestamp.
func WithScheduledOn(ts time.Time) Option {
	return func(r *Run) {
		r.scheduledOn = ts
	}
}

// WithCompletedOn sets the Run's completed on nano timestamp.
func WithCompletedOn(ts time.Time) Option {
	return func(r *Run) {
		r.completedOn = ts
	}
}

// WithFailedOn sets the Run's failed on nano timestamp.
func WithFailedOn(ts time.Time) Option {
	return func(r *Run) {
		r.completedOn = ts
	}
}

// WithCanceledOn sets the Run's canceled on nano timestamp.
func WithCanceledOn(ts time.Time) Option {
	return func(r *Run) {
		r.canceledOn = ts
	}
}

// WithPausedOn sets the Run's paused on nano timestamp.
func WithPausedOn(ts time.Time) Option {
	return func(r *Run) {
		r.pausedOn = ts
	}
}

// WithResumedOn sets the Run's resumed on nano timestamp.
func WithResumedOn(ts time.Time) Option {
	return func(r *Run) {
		r.resumedOn = ts
	}
}
