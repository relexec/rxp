package selector

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Selector allows a ObjectReader to select a specific target.
type Selector struct {
	// system is the System to find the Object in.
	system types.System
	// kindVersion is the KindVersion of the Object to look up.
	kindVersion types.KindVersion
	// uuid is the globally-unique string identifier to look up the Object for.
	uuid string
	// domain is the Domain to use when looking up the Object via name.
	domain types.Domain
	// namespace is the Namespace to use when looking up the Object via name.
	namespace types.Namespace
	// name is the Name to use when looking up the Object via name.
	name string
	// generation is the Generation of the Object's Spec that should be read.
	// If this is 0, the Object Spec's latest generation is returned.
	generation types.Generation
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.kindVersion == "" {
		return errors.ErrSelectorKindVersionRequired
	}
	err := s.kindVersion.Validate()
	if err != nil {
		return err
	}
	if s.uuid == "" && s.name == "" {
		return errors.ErrSelectorUUIDOrNameRequired
	}
	if s.system != nil {
		err = s.system.Validate()
		if err != nil {
			return err
		}
		if s.domain != nil {
			domainSys := s.domain.System()
			if domainSys != nil && s.system.UUID() != domainSys.UUID() {
				return errors.ErrSelectorSystemMismatched
			}
		}
	}
	if s.domain != nil {
		err = s.domain.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// System is the System to search for the Object in.
func (s Selector) System() types.System {
	return s.system
}

// KindVersion is the KindVersion of the Object to look up.
func (s Selector) KindVersion() types.KindVersion {
	return s.kindVersion
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

// Generation returns the Generation of the Object's Spec to read. If this is
// 0, the latest generation of the Object is read.
func (s Selector) Generation() types.Generation {
	return s.generation
}
