package types

import "context"

// Archiver archives something.
//
// The ability to cleanly, safely and efficiently remove something from an
// "active" data set while preserving its change history is a critical function
// for any durable execution platform.
//
// Once archived, something cannot be "unarchived".
type Archiver interface {
	// Archive safely removes the supplied thing from the "active" data set
	// while preserving the thing's change history and state.
	//
	// Upon successful return, any Names associated with the archived thing can
	// be re-used for a future thing of the same type. In other words, once
	// archived, a Named object's name uniqueness guarantee are no longer
	// valid. This means that archival is a one-way process and cannot be
	// undone.
	Archive(context.Context, Object) error
}
