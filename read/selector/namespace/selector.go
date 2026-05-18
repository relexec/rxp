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
	// domain is the Domain within which the Namespace is scoped.
	domain types.Domain
	// namespaceName is the NamespaceName to use when looking up the target via
	// name.
	namespaceName types.NamespaceName
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid == "" && s.namespaceName == "" {
		return errors.ErrSelectorUUIDOrNameRequired
	}
	if s.uuid != "" {
		return nil
	}
	if s.domain == nil {
		return errors.ErrSelectorDomainRequired
	}
	err := s.domain.Validate()
	if err != nil {
		return err
	}
	// We should have been given a valid namespace name.
	if s.namespaceName != "" {
		err := s.namespaceName.Validate()
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
	if s.domain == nil {
		// This is not really valid, but validate ensures that we don't end up
		// here...
		return name.New(string(s.namespaceName))
	}
	return name.New(string(s.namespaceName), name.WithDomain(s.domain))
}

var _ types.Selector = (*Selector)(nil)
