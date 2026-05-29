package object

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/system"
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
func WithSystem(system *system.System) Option {
	return func(o *Object) {
		o.system = system
	}
}

// WithKindVersionName sets the Object's KindVersionName.
func WithKindVersionName(kv api.KindVersionName) Option {
	return func(o *Object) {
		o.kindVersionName = kv
	}
}

// WithUUID sets the Object's UUID.
func WithUUID(uuid string) Option {
	return func(o *Object) {
		o.uuid = uuid
	}
}

// WithDomain sets the Object's Domain.
func WithDomain(domain *domain.Domain) Option {
	return func(o *Object) {
		o.domain = domain
	}
}

// WithNamespace sets the Object's Namespace.
func WithNamespace(namespace *namespace.Namespace) Option {
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
func WithLabels(labels api.Labels) Option {
	return func(o *Object) {
		o.labels = labels
	}
}

// WithGeneration sets the Object's Generation.
func WithGeneration(generation api.Generation) Option {
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
