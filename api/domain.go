package api

import (
	apicore "github.com/relexec/rxp/api/core"
	"github.com/relexec/rxp/errors"
)

// Domain describes a division or partition of a System.
type Domain struct {
	apicore.NestedSetRecord
	// System contains the System containing the Domain. This is a pointer to a
	// System to allow for backends to default missing System information to
	// their host system.
	System *System `json:"system,omitempty"`
	// UUID stores the Domain's globally-unique identifier.
	UUID string `json:"uuid"`
	// Name contains the Domain name.
	//
	// A valid Domain Name is a DNS-formatted (RFC 1035-compliant) name less than
	// 254 characters.
	//
	// A Domain's Name must be unique within the scope of the `rxp` system
	// installation.
	Name DomainName `json:"name"`
	// Root contains a pointer to the root Domain, if any. If empty, the Domain
	// is itself the root Domain.
	Root *Domain `json:"root,omitempty"`
	// Parent contains a pointer to the parent Domain, if any.
	Parent *Domain `json:"parent,omitempty"`
}

// Validate returns an error if the Domain is invalid.
func (d Domain) Validate() error {
	if d.UUID == "" {
		return errors.ErrDomainUUIDRequired
	}
	if d.Root != nil {
		rootSystem := d.Root.System
		if d.System != nil && rootSystem != nil {
			if rootSystem.UUID != d.System.UUID {
				return errors.ErrDomainRootSystemDifferent
			}
		}
	}
	if d.Parent != nil {
		if d.Root == nil {
			return errors.ErrDomainParentRootRequired
		}
		parentSystem := d.Parent.System
		if d.System != nil && parentSystem != nil {
			if parentSystem.UUID != d.System.UUID {
				return errors.ErrDomainParentSystemDifferent
			}
		}
	}
	return d.Name.Validate()
}

// IsRoot returns true if the Domain is itself the root domain.
func (d Domain) IsRoot() bool {
	return d.Root == nil || d.Root.UUID == d.UUID
}
