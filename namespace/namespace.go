package namespace

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Namespace describes a logical division within a Domain.
type Namespace struct {
	// domain contains the Namespace's Domain.
	domain types.Domain
	// name contains the Namespace name.
	//
	// A valid Namespace Name is a DNS-formatted (RFC 1035-compliant) name less than
	// 254 characters.
	//
	// A Namespace's Name must be unique within the scope of its Domain.
	name types.NamespaceName
}

// Validate returns an error if the Domain is invalid.
func (n Namespace) Validate() error {
	if n.domain == nil {
		return errors.ErrNamespaceDomainEmpty
	}
	if err := n.domain.Validate(); err != nil {
		return err
	}
	return n.name.Validate()
}

// Domain returns the Domain of the Namespace.
func (n Namespace) Domain() types.Domain {
	return n.domain
}

// SetDomain sets the Domain of Namespace.
func (n *Namespace) SetDomain(domain types.Domain) {
	n.domain = domain
}

// Name returns the Name of the Namespace.
func (n Namespace) Name() types.NamespaceName {
	return n.name
}

// SetName sets the Name of Namespace.
func (n *Namespace) SetName(name types.NamespaceName) {
	n.name = name
}

var _ types.Namespace = (*Namespace)(nil)
