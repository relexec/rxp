package result

import (
	"github.com/relexec/rxp/object/list"
	"github.com/relexec/rxp/object/list/option"
	"github.com/relexec/rxp/types"
)

// Result wraps the [types.Object] returned from a succesful call to
// ObjectList.
type Result struct {
	// objects is the set of of Objects returned from the single call to
	// ObjectList.
	objects []types.Object
	// options is the set of Options that were used in the call to ObjectList.
	// These will have any server-side defaulted values set. For example, if
	// you did not specify a limit for the number of Objects to return, the
	// server will always bound this to a default value. That default will be
	// set in this field.
	options option.Options
	// marker contains the UUID of the last Object on a previous "page" of
	// results returned from a call to ObjectList. This value can be passed in
	// subsequent calls to ObjectList to "continue" the list from that Object.
	marker string
}

// Objects returns the set of of Objects returned from the single call to
// ObjectList.
func (r Result) Objects() []types.Object {
	return r.objects
}

// Options returns the Options that were used in the call to ObjectList.
// These will have any server-side defaulted values set. For example, if
// you did not specify a limit for the number of Objects to return, the
// server will always bound this to a default value. That default will be
// set in this field.
func (r Result) Options() option.Options {
	return r.options
}

// Marker returns the UUID of the last Object on a previous "page" of
// results returned from a call to ObjectList. This value can be passed in
// subsequent calls to ObjectList to "continue" the list from that Object.
func (r Result) Marker() string {
	return r.marker
}

// More returns true if there are more Objects to be retrived.
func (r Result) More() bool {
	return r.marker != ""
}

var _ list.Result = (*Result)(nil)
