package option

import "github.com/relexec/rxp/types"

// Options controls how a call to Read behaves.
type Options struct {
	// kindVersion returns the KindVersion of the target to look up. If empty,
	// the latest version of the Kind specified in the Selector is used.
	kindVersion types.KindVersion
	// igeneration returns the Generation of the target that should be read.  Only
	// applicable for things that can have multiple generations representing
	// mutations to desired state (e.g. Object).
	generation types.Generation
}

// Generation returns the Generation of the target that should be read.  Only
// applicable for things that can have multiple generations representing
// mutations to desired state (e.g. Object).
//
// If the Kind of target being read supports multiple generations and this
// method returns 0, the target's latest generation is read.
func (o Options) Generation() types.Generation {
	return o.generation
}

// KindVersion returns the KindVersion of the target to look up. If empty,
// the latest version of the Kind specified in the Selector is used.
func (o Options) KindVersion() types.KindVersion {
	return o.kindVersion
}
