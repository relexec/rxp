package list

import (
	"context"

	"github.com/relexec/rxp/object/list/option"
	"github.com/relexec/rxp/types"
)

// Result describes a successful ObjectList response.
type Result interface {
	// Objects returns the set of of Objects returned from the single call to
	// ObjectList.
	Objects() []types.Object
	// Options returns the Options that were used in the call to ObjectList.
	// These will have any server-side defaulted values set. For example, if
	// you did not specify a limit for the number of Objects to return, the
	// server will always bound this to a default value. That default will be
	// set in this field.
	Options() option.Options
	// Marker returns the UUID of the last Object on a previous "page" of
	// results returned from a call to ObjectList. This value can be passed in
	// subsequent calls to ObjectList to "continue" the list from that Object.
	Marker() string
	// More returns true if there are more Objects to be retrived.
	More() bool
}

// ObjectLister lists zero or more [types.Object] from persistent storage.
type ObjectLister interface {
	// ObjectList lists zero or more [types.Object] from persistent storage.
	ObjectList(
		ctx context.Context,
		expr types.Expression,
		opts ...option.Option,
	) (Result, error)
}
