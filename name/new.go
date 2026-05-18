package name

import "github.com/relexec/rxp/types"

// Option controls the returned [Name] from [New]
type Option func(*Name)

// WithSystem sets the Name's optional System qualifier.
func WithSystem(system types.System) Option {
	return func(n *Name) {
		n.system = system
	}
}

// WithDomain sets the Name's optional Domain qualifier.
func WithDomain(domain types.Domain) Option {
	return func(n *Name) {
		n.domain = domain
	}
}

// WithNamespace sets the Name's optional Namespace qualifier.
func WithNamespace(namespace types.Namespace) Option {
	return func(n *Name) {
		n.namespace = namespace
	}
}

// New returns a new Name given a human-readable string name and zero or more
// Option modifiers.
func New(name string, opts ...Option) Name {
	n := Name{name: name}
	for _, opt := range opts {
		opt(&n)
	}
	return n
}
