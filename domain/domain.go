package domain

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Domain describes a top-level division or partition of things managed by rxp.
type Domain struct {
	// system contains the system identifier for the Domain.
	system string
	// name contains the Domain name.
	//
	// A valid Domain Name is a DNS-formatted (RFC 1035-compliant) name less than
	// 254 characters.
	//
	// A Domain's Name must be unique within the scope of the `rxp` system
	// installation.
	name types.DomainName
}

// Validate returns an error if the Domain is invalid.
func (d Domain) Validate() error {
	if d.name == "" {
		return errors.ErrDomainNameEmpty
	}
	return d.name.Validate()
}

// System returns the System of the Domain.
func (d Domain) System() string {
	return d.system
}

// SetSystem sets the System of Domain.
func (d *Domain) SetSystem(system string) {
	d.system = system
}

// Name returns the Name of the Domain.
func (d Domain) Name() types.DomainName {
	return d.name
}

// SetName sets the Name of Domain.
func (d *Domain) SetName(name types.DomainName) {
	d.name = name
}

var _ types.Domain = (*Domain)(nil)
