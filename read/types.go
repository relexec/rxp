package read

import (
	"context"

	"github.com/relexec/rxp/read/option"
	"github.com/relexec/rxp/types"
)

type Options interface {
	// KindVersion returns the KindVersion of the target to look up. If empty,
	// the latest version of the Kind specified in the Selector is used.
	KindVersion() types.KindVersion
	// Generation returns the Generation of the target that should be read.
	// Only applicable for things that can have multiple generations
	// representing mutations to desired state (e.g. Object).
	//
	// If the Kind of target being read supports multiple generations and this
	// method returns 0, the target's latest generation is read.
	Generation() types.Generation
}

type Result[T any] interface {
	// Items returns the set of things returned from the single call to List.
	Items() []T
	// Options returns the Options that were used in the call to List.  These
	// will have any server-side defaulted values set. For example, if you did
	// not specify a limit for the number of items to return, the server will
	// always bound this to a default value. That default will be set in this
	// field.
	Options() Options
	// Marker returns the UUID of the last item on a previous "page" of
	// results returned from a call to List. This value can be passed in
	// subsequent calls to List to "continue" the list from that item.
	Marker() string
	// More returns true if there are more items to be retrieved.
	More() bool
}

// Reader reads a single thing from persistent storage.
type Reader[T any] interface {
	// Read reads a single thing from persistent storage.
	Read(
		context.Context,
		types.Selector,
		...option.Option,
	) (T, error)
}
