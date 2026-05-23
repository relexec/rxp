package meta

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/system"
)

// Selector selects a single Meta.
type Selector struct {
	// uuid is the globally-unique identifier of the Meta being selected.
	uuid string
	// kindVersion is the KindVersion of the Meta being selected.
	kindVersion api.KindVersion
	// system is the System to find the Meta in.
	system *system.System
	// domain is the Domain to find the Meta in.
	domain *domain.Domain
	// namespace is the Namespace to find the Meta in.
	namespace *namespace.Namespace
}

// UUID returns the globally-unique identifier of the Meta being selected.
func (s Selector) UUID() string {
	return s.uuid
}

// KindVersion returns the KindVersion to use when looking up the Meta.
func (s Selector) KindVersion() api.KindVersion {
	return s.kindVersion
}

// System is the System to search for the Meta in.
func (s Selector) System() *system.System {
	return s.system
}

// Domain is the Domain to search for the Meta in.
func (s Selector) Domain() *domain.Domain {
	return s.domain
}

// Namespace is the Namespace to search for the Meta in.
func (s Selector) Namespace() *namespace.Namespace {
	return s.namespace
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid != "" {
		return nil
	}
	err := s.kindVersion.Validate()
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

// ByKindVersion sets the Selector's KindVersion.
func ByKindVersion(kindVersion api.KindVersion) SelectOption {
	return func(s *Selector) {
		s.kindVersion = kindVersion
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
