package types

// Selector enables `rxp` backend implementations to read a *single thing*.
//
// Either UUID() or Name() must return a non-empty value.
type Selector interface {
	Validatable
	// UUID returns the globally-unique identifier of the thing being selected.
	UUID() string
	// Name returns the qualified name of the thing being selected.
	Name() Name
}
