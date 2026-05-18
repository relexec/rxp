package name

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Name describes a string name that is optionally qualified by a
// System, Domain or Namescope.
//
// Things in `rxp`'s  data model can be selected by UUID or by name. When
// selected by name and the Kind of thing being selected has a Namescope of
// NamescopeDomain or NamescopeNamespace, the name must be qualified with the
// Domain or Namespace, respectively.
type Name struct {
	// name is the string name of the thing.
	name string
	// system is the System in which the named thing is scoped.
	system types.System
	// domain is the Domain in which the named thing is scoped.
	domain types.Domain
	// namespace is the Namescope in which the named thing is scoped.
	namespace types.Namespace
}

// Name returns the string name of the thing.
func (n Name) Name() string {
	return n.name
}

// System returns the System in which the named thing is scoped.
func (n Name) System() types.System {
	return n.system
}

// Domain returns the Domain in which the named thing is scoped.
func (n Name) Domain() types.Domain {
	return n.domain
}

// Namespace returns the Namespace in which the named thing is scoped.
func (n Name) Namespace() types.Namespace {
	return n.namespace
}

// Validate returns an error if the Name is not valid.
func (n Name) Validate() error {
	if n.name == "" {
		return errors.ErrNameNameRequired
	}
	if n.system != nil {
		err := n.system.Validate()
		if err != nil {
			return err
		}
	}
	if n.domain != nil {
		err := n.domain.Validate()
		if err != nil {
			return err
		}
	}
	if n.namespace != nil {
		err := n.namespace.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
