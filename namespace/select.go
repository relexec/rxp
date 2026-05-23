package namespace

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/errors"
)

// Selector selects a single Namespace.
type Selector struct {
	// uuid is the globally-unique identifier of the Namespace being selected.
	uuid string
	// domain is the Domain to find the Namespace in.
	domain *domain.Domain
	// name is the name to look up the Namespace for.
	name api.NamespaceName
}

// UUID returns the globally-unique identifier of the Domain being selected.
func (s Selector) UUID() string {
	return s.uuid
}

// Domain is the Domain to search for the Namespace in.
func (s Selector) Domain() *domain.Domain {
	return s.domain
}

// Name returns the name to use when looking up the Namespace.
func (s Selector) Name() api.NamespaceName {
	return s.name
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid != "" {
		return nil
	}
	if s.name == "" {
		return errors.ErrSelectorUUIDOrNameRequired
	}
	if s.domain == nil {
		return errors.ErrSelectorDomainRequired
	} else {
		if err := s.domain.Validate(); err != nil {
			return err
		}
	}
	return s.name.Validate()
}

// SelectOption modifies the [Selector] returned from [Select].
type SelectOption func(*Selector)

// ByUUID sets the Selector's UUID.
func ByUUID(uuid string) SelectOption {
	return func(s *Selector) {
		s.uuid = uuid
	}
}

// ByDomain sets the Selector's Domain.
func ByDomain(dom *domain.Domain) SelectOption {
	return func(s *Selector) {
		s.domain = dom
	}
}

// ByName sets the Selector's Name.
func ByName(name api.NamespaceName) SelectOption {
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
