package types

import "fmt"

// Namescope describes the uniqueness constraint of some thing's name.
type Namescope int

// NOTE(jaypipes): These Namescopes are listed in order of specificity, from
// the narrowest specificity (names are unique within the Kind, Namespace and
// Domain) to the broadest specificity (names are globally-unique).

const (
	// NamescopeNamespace means the thing's name is unique within its Kind +
	// Domain + Namespace.
	NamescopeNamespace Namescope = 1
	// NamescopeDomain means the thing's name is unique within its Kind +
	// Domain.
	NamescopeDomain Namescope = 2
	// NamescopeKind means the thing's name is unique within its Kind.
	NamescopeKind Namescope = 3
	// NamescopeSystem means the thing's name is unique within the `rxp` system
	// installation.
	NamescopeSystem Namescope = 4
	// NamescopeGlobal means the thing's name is globally unique.
	NamescopeGlobal Namescope = 5
)

func (n Namescope) String() string {
	switch n {
	case NamescopeNamespace:
		return "namespace"
	case NamescopeDomain:
		return "domain"
	case NamescopeKind:
		return "kind"
	case NamescopeSystem:
		return "system"
	case NamescopeGlobal:
		return "global"
	default:
		return "unknown"
	}
}

// Validate returns an error if the Namescope is not valid.
func (n Namescope) Validate() error {
	switch n {
	case NamescopeNamespace,
		NamescopeDomain,
		NamescopeKind,
		NamescopeSystem,
		NamescopeGlobal:
		return nil
	default:
		return fmt.Errorf("unknown namescope")
	}
}
