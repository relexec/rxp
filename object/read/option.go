package read

import (
	"github.com/relexec/rxp/types"
)

// Option can be used to control how [ObjectReader.ReadObject] method behaves.
type Option func(*Options)

// Options controls how a call to one of the [ObjectReader.ReadObject] methods
// behaves.
type Options struct {
	// generation is the Generation of the Object's Spec that should be read.
	// If this is 0, the Object Spec's latest generation is returned.
	generation types.Generation
}

// Generation returns the Generation of the Object's Spec to read.
func (o Options) Generation() types.Generation {
	return o.generation
}

// HavingGeneration is used to look up an Object with a Spec having the supplied
// Generation. If HavingGeneration is not used or a value of 0 is supplied to
// HavingGeneration, the Object Spec's latest generation is read.
func HavingGeneration(generation types.Generation) Option {
	return func(o *Options) {
		o.generation = generation
	}
}

// NewOptions returns a new Options given zero or more Option modifiers.
func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}
