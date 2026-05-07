package types

// System represents the boundaries of an rxp system installation.
type System interface {
	Validatable
	// UUID returns the System's globally-unique identifier.
	UUID() string
	// Name returns the optional human-readable name of the System.
	Name() string
}
