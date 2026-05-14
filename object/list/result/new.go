package result

import (
	"github.com/relexec/rxp/object/list/option"
	"github.com/relexec/rxp/types"
)

// Option modifies a Result returned from New.
type Option func(*Result)

// New returns a new [Result].
func New(opts ...Option) *Result {
	m := &Result{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithObjects sets the Result's collection of Objects.
func WithObjects(objects []types.Object) Option {
	return func(r *Result) {
		r.objects = objects
	}
}

// WithOptions sets the Result's Options.
func WithOptions(opts option.Options) Option {
	return func(r *Result) {
		r.options = opts
	}
}

// WithMarker sets the Result's Marker.
func WithMarker(marker string) Option {
	return func(r *Result) {
		r.marker = marker
	}
}
