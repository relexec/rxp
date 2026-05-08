package selector

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Selector allows a NamespaceReader to select a specific target.
type Selector struct {
	// domain is the Domain to look up the Namespace for.
	domain types.Domain
	// name is the name to look up the Namespace for.
	name types.NamespaceName
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.domain == nil {
		return errors.ErrSelectorDomainRequired
	} else {
		if err := s.domain.Validate(); err != nil {
			return err
		}
	}
	if s.name == "" {
		return errors.ErrSelectorNameRequired
	}
	return s.name.Validate()
}

// Domain returns the Domain to use when looking up the Namespace.
func (s Selector) Domain() types.Domain {
	return s.domain
}

// Name returns the name to use when looking up the Namespace.
func (s Selector) Name() types.NamespaceName {
	return s.name
}
