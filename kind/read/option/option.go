package option

import (
	"github.com/relexec/rxp/types"
)

// Options controls how a call to [KindReader.KindRead] behaves.
type Options struct {
	// generation is the Generation of the Kind that should be read. If this
	// is 0, the Kind's latest generation is read.
	generation types.Generation
}

// Generation returns the Generation to read.
func (o Options) Generation() types.Generation {
	return o.generation
}
