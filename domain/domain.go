package domain

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/system"
)

var (
	FieldPathSystem = fieldpath.FromString("system")
	FieldPathUUID   = fieldpath.FromString("uuid")
	FieldPathName   = fieldpath.FromString("name")
	FieldPathParent = fieldpath.FromString("parent")
)

// Domain describes a division or partition of a System.
type Domain struct {
	// system contains the System containing the Domain.
	system *system.System
	// uuid stores the Domain's globally-unique identifier.
	uuid string
	// name contains the Domain name.
	//
	// A valid Domain Name is a DNS-formatted (RFC 1035-compliant) name less than
	// 254 characters.
	//
	// A Domain's Name must be unique within the scope of the `rxp` system
	// installation.
	name api.DomainName
	// parent contains a pointer to the parent Domain, if any.
	parent *Domain
}

// Validate returns an error if the Domain is invalid.
func (d Domain) Validate() error {
	if d.uuid == "" {
		return errors.ErrDomainUUIDRequired
	}
	if d.parent != nil {
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
func (d Domain) System() *system.System {
	return d.system
}

// SetSystem sets the System of Domain.
func (d *Domain) SetSystem(system *system.System) {
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
func (d Domain) Name() api.DomainName {
	return d.name
}

// SetName sets the Name of Domain.
func (d *Domain) SetName(name api.DomainName) {
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

// IsRoot returns true if the Domain has no parent.
func (d *Domain) IsRoot() bool {
	return d.parent == nil
}

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (d Domain) Diff(subject any) (*cmp.Delta, error) {
	var other *Domain
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
		thisSystemUUID := thisSystem.UUID()
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
			if thisSystemUUID != otherSystem.UUID() {
				delta.Push(
					cmp.NewDifference(
						FieldPathSystem,
						cmp.DifferenceTypeModify,
						thisSystemUUID,
						otherSystemUUID,
					),
				)
			}
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

	thisUUID := d.uuid
	otherUUID := other.UUID()
	if thisUUID != otherUUID {
		delta.Push(
			cmp.NewDifference(
				FieldPathUUID,
				cmp.DifferenceTypeModify,
				string(thisUUID),
				string(otherUUID),
			),
		)
	}

	thisName := d.name
	otherName := other.Name()
	if thisName != otherName {
		delta.Push(
			cmp.NewDifference(
				FieldPathName,
				cmp.DifferenceTypeModify,
				string(thisName),
				string(otherName),
			),
		)
	}

	thisParent := d.parent
	otherParent := other.Parent()
	if thisParent != nil {
		thisParentUUID := thisParent.UUID()
		if otherParent == nil {
			delta.Push(
				cmp.NewDifference(
					FieldPathParent,
					cmp.DifferenceTypeRemove,
					thisParentUUID,
					nil,
				),
			)
		} else {
			otherParentUUID := otherParent.UUID()
			if thisParentUUID != otherParent.UUID() {
				delta.Push(
					cmp.NewDifference(
						FieldPathParent,
						cmp.DifferenceTypeModify,
						thisParentUUID,
						otherParentUUID,
					),
				)
			}
		}
	} else if otherParent != nil {
		otherParentUUID := otherParent.UUID()
		delta.Push(
			cmp.NewDifference(
				FieldPathParent,
				cmp.DifferenceTypeAdd,
				nil,
				otherParentUUID,
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
			FieldPathUUID,
			cmp.DifferenceTypeAdd,
			d.uuid,
			nil,
		),
	)
	delta.Push(
		cmp.NewDifference(
			FieldPathName,
			cmp.DifferenceTypeAdd,
			string(d.name),
			nil,
		),
	)
	if d.parent != nil {
		delta.Push(
			cmp.NewDifference(
				FieldPathParent,
				cmp.DifferenceTypeAdd,
				d.parent.UUID(),
				nil,
			),
		)
	}
	return delta, nil
}
