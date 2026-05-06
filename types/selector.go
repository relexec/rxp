package types

// Selector allows a Reader to select a specific target.
type Selector interface {
	Validatable
	// ID returns globally-unique identifier to select.
	ID() string
	// Domain returns the Domain to select.
	Domain() string
	// Namespace returns the Namespace to select.
	Namespace() string
	// Name returns the name to select.
	Name() string
}
