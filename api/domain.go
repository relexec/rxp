package api

import (
	"github.com/relexec/rxp/errors"
)

// Domain describes a division or partition of a System.
type Domain struct {
	// system contains the System containing the Domain.
	system *System
	// uuid stores the Domain's globally-unique identifier.
	uuid string
	// name contains the Domain name.
	//
	// A valid Domain Name is a DNS-formatted (RFC 1035-compliant) name less than
	// 254 characters.
	//
	// A Domain's Name must be unique within the scope of the `rxp` system
	// installation.
	name DomainName
	// root contains a pointer to the root Domain, if any. If empty, the Domain
	// is itself the root Domain.
	root *Domain
	// parent contains a pointer to the parent Domain, if any.
	parent *Domain
}

// Validate returns an error if the Domain is invalid.
func (d Domain) Validate() error {
	if d.uuid == "" {
		return errors.ErrDomainUUIDRequired
	}
	if d.root != nil {
		rootSystem := d.root.System()
		if d.system != nil && rootSystem != nil {
			if rootSystem.UUID() != d.system.UUID() {
				return errors.ErrDomainRootSystemDifferent
			}
		}
	}
	if d.parent != nil {
		if d.root == nil {
			return errors.ErrDomainParentRootRequired
		}
		parentSystem := d.parent.System()
		if d.system != nil && parentSystem != nil {
			if parentSystem.UUID() != d.system.UUID() {
				return errors.ErrDomainParentSystemDifferent
			}
		}
	}
	return d.name.Validate()
}

// System returns the System of the Domain.
func (d Domain) System() *System {
	return d.system
}

// SetSystem sets the System of Domain.
func (d *Domain) SetSystem(system *System) {
	d.system = system
}

// UUID returns the globally-unique identifier of the Domain.
func (d Domain) UUID() string {
	return d.uuid
}

// SetUUID sets the globally-unique identifier of the Domain.
func (d *Domain) SetUUID(uuid string) {
	d.uuid = uuid
}

// Name returns the Name of the Domain.
func (d Domain) Name() DomainName {
	return d.name
}

// SetName sets the Name of Domain.
func (d *Domain) SetName(name DomainName) {
	d.name = name
}

// Parent returns the Parent of the Domain.
func (d Domain) Parent() *Domain {
	return d.parent
}

// SetParent sets the Parent of Domain.
func (d *Domain) SetParent(parent *Domain) {
	d.parent = parent
}

// Root returns the Root of the Domain. If nil, the Domain itself is the root
// Domain.
func (d Domain) Root() *Domain {
	return d.root
}

// SetRoot sets the Root of Domain.
func (d *Domain) SetRoot(root *Domain) {
	d.root = root
}

// IsRoot returns true if the Domain is itself the root domain.
func (d *Domain) IsRoot() bool {
	return d.root == nil
}
