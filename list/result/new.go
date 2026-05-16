package result

import (
	"github.com/relexec/rxp/list/option"
)

// Option modifies a Result returned from New.
type Option[T any] func(*Result[T])

// New returns a new [Result].
func New[T any](opts ...Option[T]) *Result[T] {
	m := &Result[T]{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithItems sets the Result's collection of items.
func WithItems[T any](items []T) Option[T] {
	return func(r *Result[T]) {
		r.items = items
	}
}

// WithOptions sets the Result's Options.
func WithOptions[T any](opts option.Options) Option[T] {
	return func(r *Result[T]) {
		r.options = opts
	}
}

// WithMarker sets the Result's Marker.
func WithMarker[T any](marker string) Option[T] {
	return func(r *Result[T]) {
		r.marker = marker
	}
}
