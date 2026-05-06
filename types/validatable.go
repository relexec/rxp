package types

// Validatable can check if it is valid.
type Validatable interface {
	// Validate returns an error if the thing is not valid.
	Validate() error
}
