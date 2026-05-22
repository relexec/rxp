package object

import (
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

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid == "" && s.name == "" {
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

// ByUUID returns a Selector that looks up an Object having the supplied
// UUID.
func ByUUID(uuid string) Selector {
	return Selector{uuid: uuid}
}

// BySystemAndName returns a Selector that looks up a Object having the
// supplied Name in the supplied System.
func BySystemAndName(sys *system.System, name string) Selector {
	return Selector{system: sys, name: name}
}

// ByDomainAndName returns a Selector that looks up a Object having the
// supplied Name in the supplied Domain.
func ByDomainAndName(dom *domain.Domain, name string) Selector {
	return Selector{domain: dom, name: name}
}

// ByNamespaceAndName returns a Selector that looks up a Object having
// the supplied Name in the supplied Namespace.
func ByNamespaceAndName(ns *namespace.Namespace, name string) Selector {
	return Selector{namespace: ns, name: name}
}
