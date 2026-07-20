package kindversion

import (
	"github.com/relexec/rxp/api"
)

// Selector selects a single KindVersion.
type Selector struct {
	// uuid is the globally-unique identifier of the KindVersion being
	// selected.
	uuid string
	// name is the KindVersionName of the KindVersion being selected.
	name api.KindVersionName
	// system is the System to find the KindVersion in.
	system *api.System
	// domain is the Domain to find the KindVersion in.
	domain *api.Domain
}

// UUID returns the globally-unique identifier of the KindVersion being
// selected.
func (s Selector) UUID() string {
	return s.uuid
}

// Name returns the KindVersionName to use when looking up the KindVersion.
func (s Selector) Name() api.KindVersionName {
	return s.name
}

// System is the System to search for the KindVersion in.
func (s Selector) System() *api.System {
	return s.system
}

// Domain is the Domain to search for the KindVersion in.
func (s Selector) Domain() *api.Domain {
	return s.domain
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid != "" {
		return nil
	}
	err := s.name.Validate()
	if err != nil {
		return err
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
func BySystem(system *api.System) SelectOption {
	return func(s *Selector) {
		s.system = system
	}
}

// ByDomain sets the Selector's Domain.
func ByDomain(domain *api.Domain) SelectOption {
	return func(s *Selector) {
		s.domain = domain
	}
}

// ByName sets the Selector's Name.
func ByName(name api.KindVersionName) SelectOption {
	return func(s *Selector) {
		s.name = name
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
