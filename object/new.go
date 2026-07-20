package object

import (
	"github.com/relexec/rxp/api"
)

// Option modifies an Object returned from New.
type Option func(*api.Object)

// New returns a new [Object]
func New(opts ...Option) *api.Object {
	o := &api.Object{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// WithSystem sets the Object's System.
func WithSystem(system *api.System) Option {
	return func(o *api.Object) {
		o.SetSystem(system)
	}
}

// WithKindVersionName sets the Object's KindVersionName.
func WithKindVersionName(kv api.KindVersionName) Option {
	return func(o *api.Object) {
		o.SetKindVersionName(kv)
	}
}

// WithUUID sets the Object's UUID.
func WithUUID(uuid string) Option {
	return func(o *api.Object) {
		o.SetUUID(uuid)
	}
}

// WithDomain sets the Object's Domain.
func WithDomain(domain *api.Domain) Option {
	return func(o *api.Object) {
		o.SetDomain(domain)
	}
}

// WithName sets the Object's Name.
func WithName(name string) Option {
	return func(o *api.Object) {
		o.SetName(name)
	}
}

// WithLabels sets the Object's Labels.
func WithLabels(labels api.Labels) Option {
	return func(o *api.Object) {
		o.SetLabels(labels)
	}
}

// WithGeneration sets the Object's Generation.
func WithGeneration(generation api.Generation) Option {
	return func(o *api.Object) {
		o.SetGeneration(generation)
	}
}

// WithSpec sets the Object's Spec to the supplied JSON-encoded string of
// desired state.
func WithSpec(spec string) Option {
	return func(o *api.Object) {
		o.SetSpec(spec)
	}
}
