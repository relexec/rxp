package api

// Object is an *instance* of a KindVersion.
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
	// kindVersionName is the kind and version identifier for the type of
	// Object.
	kindVersionName KindVersionName
	// system contains the system identifier for the Object.
	system *System
	// uuid is the globally-unique string identifier.
	uuid string
	// domain is the optional Domain.
	domain *Domain
	// name is the Name.
	name string
	// labels is the collection of Labels.
	labels Labels
	// generation contains the generation of the Object's desired state.
	generation Generation
	// spec contains the Object's desired state encoded as a JSON string.
	spec string
}

// KindName returns the DNS-formatted name of the Kind of Object, e.g.
// `flow.temporal.io`.
func (o Object) KindName() KindName {
	return o.kindVersionName.Kind()
}

// KindVersionName returns the KindVersionName of the Object. This string
// uniquely identifies the type of an Object.
func (o Object) KindVersionName() KindVersionName {
	return o.kindVersionName
}

// SetKindVersionName sets the Object's kind version name which uniquely
// identifies the type of an Object.
func (o *Object) SetKindVersionName(kvn KindVersionName) {
	o.kindVersionName = kvn
}

// System returns the System of the Object.
func (o Object) System() *System {
	return o.system
}

// SetSystem sets the System of Object.
func (o *Object) SetSystem(system *System) {
	o.system = system
}

// UUID returns the globally-unique string identifier.
func (o Object) UUID() string {
	return o.uuid
}

// SetUUID sets the globally-unique string identifier.
func (o *Object) SetUUID(uuid string) {
	o.uuid = uuid
}

// Domain returns the optional Domain.
func (o Object) Domain() *Domain {
	return o.domain
}

// SetDomain sets the Domain.
func (o *Object) SetDomain(domain *Domain) {
	o.domain = domain
}

// Name returns the name. The Scope of the Object's Kind is used to determine
// whether the name is unique globally, within a Kind + System or within a Kind
// + Domain.
func (o Object) Name() string {
	return o.name
}

// SetName sets the name.
func (o *Object) SetName(name string) {
	o.name = name
}

// Labels returns the collection of Labels.
func (o Object) Labels() Labels {
	return o.labels
}

// SetLabels sets the collection of Labels.
func (o *Object) SetLabels(labels Labels) {
	o.labels = labels
}

// Generation returns the Object's Generation, which represents the number of
// mutations to the Object's desired state.
func (o Object) Generation() Generation {
	return o.generation
}

// SetGeneration sets the Object's Generation, which represents the number of
// mutations to the Object's desired state.
func (o *Object) SetGeneration(generation Generation) {
	o.generation = generation
}

// Spec returns the Object's desired state as a JSON-encoded string.
func (o Object) Spec() string {
	return o.spec
}

// SetSpec sets the Object's desired state as a JSON-encoded string.
func (o *Object) SetSpec(spec string) {
	o.spec = spec
}

// Clone returns a copy of the Object.
func (o Object) Clone() Object {
	return o
}
