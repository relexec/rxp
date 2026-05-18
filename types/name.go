package types

// Name describes a string name that is optionally qualified by a
// System, Domain or Namescope.
//
// Things in `rxp`'s  data model can be selected by UUID or by name. When
// selected by name and the Kind of thing being selected has a Namescope of
// NamescopeDomain or NamescopeNamespace, the name must be qualified with the
// Domain or Namespace, respectively.
type Name interface {
	Validatable
	// Name returns the string name of the thing being selected. If the Kind of
	// thing being selected has a Namescope of NamescopeDomain or
	// NamescopeNamespace the returned values from Domain() or Namespace() must
	// be non-nil.
	Name() string
	// System returns the System in which the thing being selected is scoped.
	// If nil, the host System of the `rxp` endpoint being called is assumed.
	System() System
	// Domain returns the Domain in which the thing being selected is scoped.
	// Must be non-nil if the Kind of thing being selected has a Namescope of
	// NamescopeDomain.
	Domain() Domain
	// Namespace returns the Namespace in which the thing being selected is
	// scoped.  Must be non-nil if the Kind of thing being selected has a
	// Namescope of NamescopeNamespace.
	Namespace() Namespace
}
