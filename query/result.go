package query

// Result wraps the slice of things returned from a successful call to
// a List call.
type Result[T any] struct {
	// items is the set of things returned from the single call to List.
	items []T
	// options is the set of Options that were used in the call to List.  These
	// will have any server-side defaulted values set. For example, if you did
	// not specify a limit for the number of items to return, the server will
	// always bound this to a default value. That default will be set in this
	// field.
	options Options
	// marker contains the UUID of the last item on a previous "page" of
	// results returned from a call to List. This value can be passed in
	// subsequent calls to List to "continue" the query from that item.
	marker string
}

// Items returns the set of things returned from the single call to List.
func (r Result[T]) Items() []T {
	return r.items
}

// Options returns the Options that were used in the call to ObjectList.
// These will have any server-side defaulted values set. For example, if
// you did not specify a limit for the number of Objects to return, the
// server will always bound this to a default value. That default will be
// set in this field.
func (r Result[T]) Options() Options {
	return r.options
}

// Marker returns the UUID of the last Object on a previous "page" of results
// returned from a call to ObjectList. This value can be passed in subsequent
// calls to Query to "continue" the query from that Object.
func (r Result[T]) Marker() string {
	return r.marker
}

// More returns true if there are more Objects to be retrived.
func (r Result[T]) More() bool {
	return r.marker != ""
}

// ResultModifier modifies a Result returned from NewResult.
type ResultModifier[T any] func(*Result[T])

// NewResult returns a new [Result].
func NewResult[T any](opts ...ResultModifier[T]) *Result[T] {
	m := &Result[T]{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// ResultWithItems sets the Result's collection of items.
func ResultWithItems[T any](items []T) ResultModifier[T] {
	return func(r *Result[T]) {
		r.items = items
	}
}

// ResultWithOptions sets the Result's Options.
func ResultWithOptions[T any](opts Options) ResultModifier[T] {
	return func(r *Result[T]) {
		r.options = opts
	}
}

// ResultWithMarker sets the Result's Marker.
func ResultWithMarker[T any](marker string) ResultModifier[T] {
	return func(r *Result[T]) {
		r.marker = marker
	}
}
