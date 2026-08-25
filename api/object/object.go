package apiobject

import (
	apicore "github.com/relexec/rxp/api/core"
	apidomain "github.com/relexec/rxp/api/domain"
	apikind "github.com/relexec/rxp/api/kind"
	apikindversion "github.com/relexec/rxp/api/kindversion"
	apisystem "github.com/relexec/rxp/api/system"
)

// Object is an *instance* of a KindVersion.
//
// Object is designed for read-heavy, write-seldom entities that need strict
// desired state mutation tracking.
//
// Each Object has a UUID globally-unique identifier.
//
// Objects have a `Name`. An Object's `Name` is unique within the
// Scope associated with the Object's Kind.
//
// If that Scope is `ScopeDomain`, the Object will have a Domain.
//
// Objects may have zero or more `Labels` associated with them. `Labels` are
// structures with a `Key` and optional `Value` that can be used to categorize
// Objects and filter them in query operations.
type Object struct {
	apicore.Record
	// KindVersionName is the kind and version identifier for the type of
	// Object.
	KindVersionName apikindversion.Name
	// System contains the System containing the Object. This is a pointer to a
	// System to allow for backends to default missing System information to
	// their host system.
	System *apisystem.System
	// UUID stores the Object's globally-unique identifier.
	UUID string
	// Domain is the optional Domain.
	Domain *apidomain.Domain
	// Name is the human-readable name for the Object. The uniqueness of the
	// Name is guaranteed within the Kind's Scope.
	Name string
	// Labels is the Object's collection of Labels.
	Labels apicore.Labels
	// Generation contains the generation of the Object's desired state.
	Generation apicore.Generation
	// Spec contains the Object's desired state encoded as a JSON string.
	Spec string
}

// KindName returns the DNS-formatted name of the Kind of Object, e.g.
// `flow.temporal.io`.
func (o Object) KindName() apikind.Name {
	return o.KindVersionName.Kind()
}

// Clone returns a copy of the Object.
func (o Object) Clone() Object {
	return o
}
