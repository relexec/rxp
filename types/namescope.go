package types

import "fmt"

// Namescope describes the uniqueness constraint of some thing's name.
type Namescope int

// NOTE(jaypipes): These Namescopes are listed in order of specificity, from
// the narrowest specificity (names are unique within the Kind, Namespace and
// Domain) to the broadest specificity (names are globally-unique).

const (
	// NamescopeNamespace means the thing's name is unique within its System,
	// Kind, Domain and Namespace.
	NamescopeNamespace Namescope = 1
	// NamescopeDomain means the thing's name is unique within its System, Kind
	// and Domain.
	NamescopeDomain Namescope = 2
	// NamescopeSystem means the thing's name is unique within its System and
	// Kind.
	NamescopeSystem Namescope = 3
)

func (n Namescope) String() string {
	switch n {
	case NamescopeNamespace:
		return "namespace"
	case NamescopeDomain:
		return "domain"
	case NamescopeSystem:
		return "system"
	default:
		return "unknown"
	}
}

// Validate returns an error if the Namescope is not valid.
func (n Namescope) Validate() error {
	switch n {
	case NamescopeNamespace,
		NamescopeDomain,
		NamescopeSystem:
		return nil
	default:
		return fmt.Errorf("unknown namescope")
	}
}
