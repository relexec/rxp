package option

import (
	"github.com/relexec/rxp/types"
)

// Option can be used to control how [MetaReader.MetaRead] behaves.
type Option func(*Options)

// Options controls how a call to [MetaReader.MetaRead] behaves.
type Options struct {
	// generation is the Generation of the Meta that should be read. If this
	// is 0, the Meta's latest generation is read.
	generation types.Generation
}

// Generation returns the Generation to read.
func (o Options) Generation() types.Generation {
	return o.generation
}

// HavingGeneration is used to look up a Meta having the supplied Generation.
// If HavingGeneration is not used or a value of 0 is supplied to
// HavingGeneration, the Meta latest generation is read.
func HavingGeneration(generation types.Generation) Option {
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
