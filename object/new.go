package object

import (
	"github.com/relexec/rxp/types"
)

// Option modifies an Object returned from New.
type Option func(*Object)

// New returns a new [Object]
func New(opts ...Option) *Object {
	o := &Object{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithSystem sets the Object's System.
func WithSystem(system types.System) Option {
	return func(o *Object) {
		o.system = system
	}
}

// WithKindVersion sets the Object's KindVersion.
func WithKindVersion(kindVersion types.KindVersion) Option {
	return func(o *Object) {
		o.kindVersion = kindVersion
	}
}

// WithUUID sets the Object's UUID.
func WithUUID(uuid string) Option {
	return func(o *Object) {
		o.uuid = uuid
	}
}

// WithDomain sets the Object's Domain.
func WithDomain(domain types.Domain) Option {
	return func(o *Object) {
		o.domain = domain
	}
}

// WithNamespace sets the Object's Namespace.
func WithNamespace(namespace types.Namespace) Option {
	return func(o *Object) {
		o.namespace = namespace
	}
}

// WithName sets the Object's Name.
func WithName(name string) Option {
	return func(o *Object) {
		o.name = name
	}
}

// WithLabels sets the Object's Labels.
func WithLabels(labels types.Labels) Option {
	return func(o *Object) {
		o.labels = labels
	}
}

// WithGeneration sets the Object's Generation.
func WithGeneration(generation types.Generation) Option {
	return func(o *Object) {
		o.generation = generation
	}
}

// WithSpec sets the Object's Spec to the supplied JSON-encoded string of
// desired state.
func WithSpec(spec string) Option {
	return func(o *Object) {
		o.spec = spec
	}
}
