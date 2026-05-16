package result

import (
	"github.com/relexec/rxp/list"
	"github.com/relexec/rxp/list/option"
	"github.com/relexec/rxp/types"
)

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
	options option.Options
	// marker contains the UUID of the last item on a previous "page" of
	// results returned from a call to List. This value can be passed in
	// subsequent calls to List to "continue" the list from that item.
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
func (r Result[T]) Options() list.Options {
	return r.options
}

// Marker returns the UUID of the last Object on a previous "page" of
// results returned from a call to ObjectList. This value can be passed in
// subsequent calls to ObjectList to "continue" the list from that Object.
func (r Result[T]) Marker() string {
	return r.marker
}

// More returns true if there are more Objects to be retrived.
func (r Result[T]) More() bool {
	return r.marker != ""
}

var _ list.Result[types.Object] = Result[types.Object]{}
