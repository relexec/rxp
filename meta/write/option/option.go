package option

import (
	"github.com/relexec/rxp/types"
)

// Option can be used to control how [MetaWriter.MetaWrite] method behaves.
type Option func(*Options)

// Options controls how [MetaWriter.MetaWrite] method behaves.
type Options struct {
	// generation is the Generation expected of the Meta being written.
	generation types.Generation
}

// Generation returns the Generation that the existing Meta is expected to be.
func (o Options) Generation() types.Generation {
	return o.generation
}

// ExpectingGeneration is used to ensure that the Generation of the Meta being
// written is equal to the supplied value.
func ExpectingGeneration(generation types.Generation) Option {
	return func(o *Options) {
		o.generation = generation
	}
}

// New returns a new Options given zero or more Option modifiers.
func New(opts ...Option) Options {
	o := Options{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
