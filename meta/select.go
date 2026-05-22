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
	// kindVersion is the KindVersion to look up the Meta in.
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

// ByUUID returns a Selector that looks up a Namespace having the supplied UUID.
func ByUUID(uuid string) Selector {
	return Selector{uuid: uuid}
}

// ByKindVersion returns a Selector that looks up a Meta having the supplied
// KindVersion.
func ByKindVersion(kv api.KindVersion) Selector {
	return Selector{kindVersion: kv}
}

// BySystemAndKindVersion returns a Selector that looks up a Meta having the
// supplied KindVersion in the supplied System.
func BySystemAndKindVersion(sys *system.System, kv api.KindVersion) Selector {
	return Selector{system: sys, kindVersion: kv}
}

// ByDomainAndKindVersion returns a Selector that looks up a Meta having the
// supplied KindVersion in the supplied Domain.
func ByDomainAndKindVersion(dom *domain.Domain, kv api.KindVersion) Selector {
	return Selector{domain: dom, kindVersion: kv}
}

// ByNamespaceAndKindVersion returns a Selector that looks up a Meta having the
// supplied KindVersion in the supplied Namespace.
func ByNamespaceAndKindVersion(ns *namespace.Namespace, kv api.KindVersion) Selector {
	return Selector{namespace: ns, kindVersion: kv}
}
