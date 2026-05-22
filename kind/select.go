package kind

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/system"
)

// Selector selects a single Kind.
type Selector struct {
	// uuid is the globally-unique identifier of the Kind being selected.
	uuid string
	// system is the System to find the Kind in.
	system *system.System
	// name is the name to look up the Kind for.
	name api.KindName
}

// UUID returns the globally-unique identifier of the Kind being selected.
func (s Selector) UUID() string {
	return s.uuid
}

// System is the rxp system identifier to search for the Kind in.
func (s Selector) System() *system.System {
	return s.system
}

// Name returns the name to use when looking up the Kind.
func (s Selector) Name() api.KindName {
	return s.name
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid != "" {
		return nil
	}
	if s.name == "" {
		return errors.ErrSelectorNameRequired
	}
	// Note that if a nil system is provided, the host System is assumed.
	if s.system != nil {
		err := s.system.Validate()
		if err != nil {
			return err
		}
	}
	return s.name.Validate()
}

// ByUUID returns a Selector that looks up a Kind having the supplied UUID.
func ByUUID(uuid string) Selector {
	return Selector{uuid: uuid}
}

// ByName returns a Selector that looks up a Kind having the supplied KindName.
// The containing System for the selected Kind is assumed to be the host
// System.
func ByName(name api.KindName) Selector {
	return Selector{name: name}
}

// BySystemAndName returns a Selector that looks up a Kind having the supplied
// System and KindName.
func BySystemAndName(sys *system.System, name api.KindName) Selector {
	return Selector{system: sys, name: name}
}
