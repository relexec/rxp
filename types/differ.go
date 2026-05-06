package types

import (
	"github.com/relexec/rxp/cmp"
)

// Differ is something that can return a [cmp.Delta] representing the
// difference between itself and something else of the same type.
type Differ interface {
	// Diff returns a [cmp.Delta] representing the difference between itself
	// and something else of the same type.
	//
	// If the argument is the [cmp.DoesNotExist] sentinel, the returned
	// [cmp.Delta] represents instructions to create the thing.
	Diff(any) (*cmp.Delta, error)
}
