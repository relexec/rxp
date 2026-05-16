package list

import (
	"context"

	"github.com/relexec/rxp/list/option"
	"github.com/relexec/rxp/types"
)

type Options interface {
	// Limit returns the max number of items to return.
	Limit() int
	// Marker returns the UUID of the last item on the page of items previously
	// read.
	Marker() string
}

type Result[T any] interface {
	// Items returns the set of things returned from the single call to List.
	Items() []T
	// Options returns the Options that were used in the call to ObjectList.
	// These will have any server-side defaulted values set. For example, if
	// you did not specify a limit for the number of Objects to return, the
	// server will always bound this to a default value. That default will be
	// set in this field.
	Options() Options
	// Marker returns the UUID of the last item on a previous "page" of
	// results returned from a call to List. This value can be passed in
	// subsequent calls to List to "continue" the list from that item.
	Marker() string
	// More returns true if there are more items to be retrieved.
	More() bool
}

// Lister lists zero or more items.
type Lister[T any] interface {
	// List lists zero or more items.
	List(
		ctx context.Context,
		expr types.Expression,
		opts ...option.Option,
	) (Result[T], error)
}
