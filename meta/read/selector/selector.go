package selector

import (
	"github.com/relexec/rxp/types"
)

// Selector allows a MetaReader to select a specific target.
type Selector struct {
	// system is the System to find the Meta in.
	system types.System
	// kindVersion is the KindVersion to look up the Meta in.
	kindVersion types.KindVersion
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	return s.kindVersion.Validate()
}

// System is the System to search for the Meta in.
func (s Selector) System() types.System {
	return s.system
}

// KindVersion returns the KindVersion to use when looking up the Meta.
func (s Selector) KindVersion() types.KindVersion {
	return s.kindVersion
}
