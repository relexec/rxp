package apiobject

import (
	apicore "github.com/relexec/rxp/api/core"
	apidomain "github.com/relexec/rxp/api/domain"
	apisystem "github.com/relexec/rxp/api/system"
	"github.com/relexec/rxp/errors"
)

// Selector selects a single Object.
type Selector struct {
	// uuid is the globally-unique string identifier to look up the Object for.
	uuid string
	// system is the System to find the Object in.
	system *apisystem.System
	// domain is the Domain to use when looking up the Object via name.
	domain *apidomain.Domain
	// name is the Name to use when looking up the Object via name.
	name string
	// generation is the specific generation of the Object to select.
	generation apicore.Generation
}

// System is the System to search for the Object in.
func (s Selector) System() *apisystem.System {
	return s.system
}

// UUID returns the globally-unique string identifier to look up the Object
// for.
func (s Selector) UUID() string {
	return s.uuid
}

// Domain returns the Domain to use when looking up the Object via name.
func (s Selector) Domain() *apidomain.Domain {
	return s.domain
}

// Name returns the Name to use when looking up the Object via name.
func (s Selector) Name() string {
	return s.name
}

// Generation returns the specific Generation of the Object to select. If this
// returns 0, the latest Object is selected.
func (s Selector) Generation() apicore.Generation {
	return s.generation
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid != "" {
		return nil
	}
	if s.name == "" {
		return errors.ErrSelectorUUIDOrNameRequired
	}
	if s.system != nil {
		if err := s.system.Validate(); err != nil {
			return err
		}
	}
	if s.domain != nil {
		// When looking up by Domain, we need either the domain's UUID or its
		// Name, not both.
		if s.domain.UUID == "" && s.domain.Name == "" {
			return errors.ErrSelectorDomainUUIDOrNameRequired
		}
	}
	return nil
}

// SelectOption modifies the [Selector] returned from [Select].
type SelectOption func(*Selector)

// ByUUID sets the Selector's UUID.
func ByUUID(uuid string) SelectOption {
	return func(s *Selector) {
		s.uuid = uuid
	}
}

// BySystem sets the Selector's System.
func BySystem(system *apisystem.System) SelectOption {
	return func(s *Selector) {
		s.system = system
	}
}

// ByDomain sets the Selector's Domain.
func ByDomain(domain *apidomain.Domain) SelectOption {
	return func(s *Selector) {
		s.domain = domain
	}
}

// ByDomainUUID sets the Selector's Domain UUID.
func ByDomainUUID(uuid string) SelectOption {
	return func(s *Selector) {
		if s.domain == nil {
			s.domain = &apidomain.Domain{}
		}
		s.domain.UUID = uuid
	}
}

// ByName sets the Selector's Name.
func ByName(name string) SelectOption {
	return func(s *Selector) {
		s.name = name
	}
}

// ByGeneration sets the Selector's Generation.
func ByGeneration(generation apicore.Generation) SelectOption {
	return func(s *Selector) {
		s.generation = generation
	}
}

// Select returns a new [Selector]
func Select(opts ...SelectOption) Selector {
	s := Selector{}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}
