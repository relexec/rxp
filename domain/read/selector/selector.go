package selector

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Selector allows a DomainReader to select a specific target.
type Selector struct {
	// system is the System to find the Domain in.
	system types.System
	// name is the name to look up the Domain for.
	name types.DomainName
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.name == "" {
		return errors.ErrSelectorNameRequired
	}
	if s.system != nil {
		err := s.system.Validate()
		if err != nil {
			return err
		}
	}
	return s.name.Validate()
}

// System is the rxp system identifier to search for the Meta in.
func (s Selector) System() types.System {
	return s.system
}

// Name returns the name to use when looking up the Domain.
func (s Selector) Name() types.DomainName {
	return s.name
}
