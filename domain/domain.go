package domain

import (
	"github.com/relexec/rxp/types"
)

// Domain describes a top-level division or partition of things managed by rxp.
type Domain struct {
	// system contains the System containing the Domain.
	system types.System
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
	return d.name.Validate()
}

// System returns the System of the Domain.
func (d Domain) System() types.System {
	return d.system
}

// SetSystem sets the System of Domain.
func (d *Domain) SetSystem(system types.System) {
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
