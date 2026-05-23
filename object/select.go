package object

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/system"
)

// Selector selects a single Object.
type Selector struct {
	// uuid is the globally-unique string identifier to look up the Object for.
	uuid string
	// system is the System to find the Object in.
	system *system.System
	// domain is the Domain to use when looking up the Object via name.
	domain *domain.Domain
	// namespace is the Namespace to use when looking up the Object via name.
	namespace *namespace.Namespace
	// name is the Name to use when looking up the Object via name.
	name string
	// generation is the specific generation of the Object to select.
	generation api.Generation
}

// System is the System to search for the Object in.
func (s Selector) System() *system.System {
	return s.system
}

// UUID returns the globally-unique string identifier to look up the Object
// for.
func (s Selector) UUID() string {
	return s.uuid
}

// Domain returns the Domain to use when looking up the Object via name.
func (s Selector) Domain() *domain.Domain {
	return s.domain
}

// Namespace returns the Namespace to use when looking up the Object via name.
func (s Selector) Namespace() *namespace.Namespace {
	return s.namespace
}

// Name returns the Name to use when looking up the Object via name.
func (s Selector) Name() string {
	return s.name
}

// Generation returns the specific Generation of the Object to select. If this
// returns 0, the latest Object is selected.
func (s Selector) Generation() api.Generation {
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
		if err := s.domain.Validate(); err != nil {
			return err
		}
	}
	if s.namespace != nil {
		if err := s.namespace.Validate(); err != nil {
			return err
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
func BySystem(system *system.System) SelectOption {
	return func(s *Selector) {
		s.system = system
	}
}

// ByDomain sets the Selector's Domain.
func ByDomain(domain *domain.Domain) SelectOption {
	return func(s *Selector) {
		s.domain = domain
	}
}

// ByNamespace sets the Selector's Namespace.
func ByNamespace(namespace *namespace.Namespace) SelectOption {
	return func(s *Selector) {
		s.namespace = namespace
	}
}

// ByName sets the Selector's Name.
func ByName(name string) SelectOption {
	return func(s *Selector) {
		s.name = name
	}
}

// ByGeneration sets the Selector's Generation.
func ByGeneration(generation api.Generation) SelectOption {
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
