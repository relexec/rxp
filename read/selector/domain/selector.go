package selector

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/name"
	"github.com/relexec/rxp/types"
)

// Selector enables `rxp` backend implementations to read a single Domain.
//
// Either UUID() or Name() must return a non-empty value.
type Selector struct {
	// uuid is the globally-unique string identifier to look up the target for.
	uuid string
	// system is the System within which the Domain is scoped. If empty, the
	// host system of the rxp backend is assumed.
	system types.System
	// name is the DomainName to use when looking up the target via name.
	domainName types.DomainName
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid == "" && s.domainName == "" {
		return errors.ErrSelectorUUIDOrNameRequired
	}
	if s.system != nil {
		err := s.system.Validate()
		if err != nil {
			return err
		}
	}
	// We should have been given a valid domain name.
	if s.domainName != "" {
		err := s.domainName.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// UUID returns the globally-unique string identifier to look up the target
// for.
func (s Selector) UUID() string {
	return s.uuid
}

// Namespace returns the Name to use when looking up the target via name.
func (s Selector) Name() types.Name {
	if s.uuid != "" {
		return nil
	}
	if s.system == nil {
		return name.New(string(s.domainName))
	}
	return name.New(string(s.domainName), name.WithSystem(s.system))
}

var _ types.Selector = (*Selector)(nil)
