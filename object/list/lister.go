package list

import (
	"context"

	"github.com/relexec/rxp/object/list/option"
	"github.com/relexec/rxp/types"
)

// ObjectListResult wraps the [types.Object] returned from the call to
// ObjectList.
type ObjectListResult struct {
	// Objects is the set of of Objects returned from the single call to
	// ObjectList.
	Objects []types.Object
	// Options is the set of Options that were used in the call to ObjectList.
	// These will have any server-side defaulted values set. For example, if
	// you did not specify a limit for the number of Objects to return, the
	// server will always bound this to a default value. That default will be
	// set in this field.
	Options option.Options
	// Marker contains the UUID of the last Object on a previous "page" of
	// results returned from a call to ObjectList. This value can be passed in
	// subsequent calls to ObjectList to "continue" the list from that Object.
	Marker string
}

// More returns true if there are more Objects to be retrieved.
func (r ObjectListResult) More() bool {
	return r.Marker != ""
}

// ObjectLister lists zero or more [types.Object] from persistent storage.
type ObjectLister interface {
	// ObjectList lists zero or more [types.Object] from persistent storage.
	ObjectList(
		ctx context.Context,
		expr types.Expression,
		opts ...option.Option,
	) (ObjectListResult, error)
}
