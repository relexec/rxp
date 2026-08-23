package apirun

import "github.com/relexec/rxp/api"

// Identifiers contains all the information needed to identify a particular
// Run.
type Identifiers struct {
	api.Record
	// UUID is the globally-unique string identifiers for the Run.
	UUID string
	// Root points at the root Run. If this *is* a root Run, this is nil.
	Root *Identifiers
	// Parent points at the Run that spawned this Run, if any. If Root is nil,
	// Parent will always be nil. If Root is non-nil, Parent will always be
	// non-nil.
	Parent *Identifiers
}
