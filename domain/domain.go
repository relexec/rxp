package domain

import (
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathSystem = fieldpath.FromString("system")
	FieldPathName   = fieldpath.FromString("name")
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

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (d Domain) Diff(subject any) (*cmp.Delta, error) {
	var other types.Domain
	switch subject := subject.(type) {
	case cmp.ZeroGen:
		return d.diffNew()
	case Domain:
		other = &subject
	case *Domain:
		other = subject
	default:
		return nil, cmp.CannotCompareTypes(d, subject)
	}

	delta := &cmp.Delta{}

	thisSystem := d.system
	otherSystem := other.System()
	if thisSystem != nil {
		thisSystemUUID := d.system.UUID()
		if otherSystem == nil {
			delta.Push(
				cmp.NewDifference(
					FieldPathSystem,
					cmp.DifferenceTypeRemove,
					thisSystemUUID,
					nil,
				),
			)
		} else {
			otherSystemUUID := otherSystem.UUID()
			delta.Push(
				cmp.NewDifference(
					FieldPathSystem,
					cmp.DifferenceTypeModify,
					thisSystemUUID,
					otherSystemUUID,
				),
			)
		}
	} else if otherSystem != nil {
		otherSystemUUID := otherSystem.UUID()
		delta.Push(
			cmp.NewDifference(
				FieldPathSystem,
				cmp.DifferenceTypeAdd,
				nil,
				otherSystemUUID,
			),
		)
	}
	return delta, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the Domain as a
// new Domain (i.e. for the first generation)
func (d Domain) diffNew() (*cmp.Delta, error) {
	delta := &cmp.Delta{}

	if d.system != nil {
		delta.Push(
			cmp.NewDifference(
				FieldPathSystem,
				cmp.DifferenceTypeAdd,
				d.system.UUID(),
				nil,
			),
		)
	}
	delta.Push(
		cmp.NewDifference(
			FieldPathName,
			cmp.DifferenceTypeAdd,
			string(d.name),
			nil,
		),
	)
	return delta, nil
}

var _ types.Domain = (*Domain)(nil)
