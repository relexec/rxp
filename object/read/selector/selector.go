package selector

import (
	"fmt"

	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Selector allows a ObjectReader to select a specific target.
type Selector struct {
	// system is the System to find the Object in.
	system types.System
	// uuid is the globally-unique string identifier to look up the Object for.
	uuid string
	// domain is the Domain to use when looking up the Object via name.
	domain types.Domain
	// namespace is the Namespace to use when looking up the Object via name.
	namespace types.Namespace
	// name is the Name to use when looking up the Object via name.
	name string
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid == "" && s.name == "" {
		return errors.ErrUUIDOrNameRequired
	}
	if s.system != nil && s.domain != nil {
		domainSys := s.domain.System()
		if domainSys != nil && s.system.UUID() != domainSys.UUID() {
			return fmt.Errorf("Selector System does not match Selector's Domain System")
		}
	}
	return nil
}

// System is the System to search for the Object in.
func (s Selector) System() types.System {
	return s.system
}

// UUID returns the globally-unique string identifier to look up the Object
// for.
func (s Selector) UUID() string {
	return s.uuid
}

// Domain returns the Domain to use when looking up the Object via name.
func (s Selector) Domain() types.Domain {
	return s.domain
}

// Namespace returns the Namespace to use when looking up the Object via name.
func (s Selector) Namespace() types.Namespace {
	return s.namespace
}

// Name returns the Name to use when looking up the Object via name.
func (s Selector) Name() string {
	return s.name
}
