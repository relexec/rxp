package types

// Object is able to be written to and read from persistent storage.
type Object interface {
	// System returns the System associated with the Object.
	System() System
	// KindVersion returns the Object's KindVersion which uniquely identifies
	// the type and version of the Object.
	KindVersion() KindVersion
	// UUID returns the Object's globally-unique identifier.
	UUID() string
	// Domain returns the Object's optional Domain.
	Domain() Domain
	// Namespace returns the Object's optional Namespace.
	Namespace() Namespace
	// Name returns the Object's name.
	//
	// The scope of uniqueness for the name is dictated by the Namescope
	// associated with the Object's KindVersion.
	Name() string
	// Labels returns the Object's Labels.
	Labels() Labels
	// Generation returns the Object's generation number. The Generation is
	// incremented on each mutation of an Object's desired state.
	Generation() Generation
	// Spec returns the Object's desired state as a JSON-encoded string.
	Spec() string
}
