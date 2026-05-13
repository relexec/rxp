package option

import (
	"github.com/relexec/rxp/types"
)

// Option can be used to control how [KindReader.KindRead] behaves.
type Option func(*Options)

// HavingGeneration is used to look up a Kind having the supplied Generation.
// If HavingGeneration is not used or a value of 0 is supplied to
// HavingGeneration, the Kind latest generation is read.
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
