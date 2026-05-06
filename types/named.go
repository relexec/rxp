package types

// Named have a string Name that is unique within a Namescope.
type Named interface {
	// Domain returns the Named's Domain, if any.
	Domain() string
	// SetDomain sets the Named's Domain.
	SetDomain(string)
	// Namespace returns Named's Namespace, if any.
	Namespace() string
	// SetNamespace sets the Named's Namespace.
	SetNamespace(string)
	// Name returns the Named's name.
	//
	// The scope of uniqueness for the name is dictated by the Namescope.
	Name() string
	// SetName sets the Named's name
	SetName(string)
}
