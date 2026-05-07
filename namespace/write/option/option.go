package option

import (
	"github.com/relexec/rxp/types"
)

// Option can be used to control how [NamespaceWriter.NamespaceWrite] behaves.
type Option func(*Options)

// Options controls how [NamespaceWriter.NamespaceWrite] behaves.
type Options struct {
	// generation is the Generation of the Spec expected of the Namespace being
	// written.
	generation types.Generation
}

// Generation returns the Generation that the existing Namespace is expected to
// be.
func (o Options) Generation() types.Generation {
	return o.generation
}

// ExpectingGeneration is used to ensure that the Generation of the Namespace's
// Spec being written is equal to the supplied value.
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
