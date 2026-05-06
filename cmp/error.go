package cmp

import "fmt"

var (
	// ErrTypeComparison is returned when two incompatible types are compared.
	ErrTypeComparison = fmt.Errorf("incompatible type comparison")
	// ErrNotBlueprint indicates the [cmp.Delta] used to try and construct a
	// thing was not a blueprint (i.e. it did not contain only differences of
	// type DifferenceTypeAdd)
	ErrNotBlueprint = fmt.Errorf("supplied cmp.Delta was not a Blueprint")
)

// CannotCompareTypes returns ErrTypeComparison specifying the types of the
// two supplied subjects.
func CannotCompareTypes(a, b any) error {
	return fmt.Errorf("%w: %T not comparable to %T", ErrTypeComparison, a, b)
}
