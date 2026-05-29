package kindversion

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/system"
)

// Selector selects a single KindVersion.
type Selector struct {
	// uuid is the globally-unique identifier of the KindVersion being
	// selected.
	uuid string
	// name is the KindVersionName of the KindVersion being selected.
	name api.KindVersionName
	// system is the System to find the KindVersion in.
	system *system.System
	// domain is the Domain to find the KindVersion in.
	domain *domain.Domain
	// namespace is the Namespace to find the KindVersion in.
	namespace *namespace.Namespace
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
func (s Selector) System() *system.System {
	return s.system
}

// Domain is the Domain to search for the KindVersion in.
func (s Selector) Domain() *domain.Domain {
	return s.domain
}

// Namespace is the Namespace to search for the KindVersion in.
func (s Selector) Namespace() *namespace.Namespace {
	return s.namespace
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
