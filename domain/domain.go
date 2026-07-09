package domain

import (
	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/system"
)

var (
	FieldPathSystem = fieldpath.FromString("system")
	FieldPathUUID   = fieldpath.FromString("uuid")
	FieldPathName   = fieldpath.FromString("name")
	FieldPathRoot   = fieldpath.FromString("root")
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

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func (d Domain) Diff(subject any) (*delta.Delta, error) {
	var other *Domain
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return d.diffNew()
	case Domain:
		other = &subject
	case *Domain:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(d, subject)
	}

	del := &delta.Delta{}

	thisSystem := d.system
	otherSystem := other.System()
	if thisSystem != nil {
		thisSystemUUID := thisSystem.UUID()
		if otherSystem == nil {
			del.Push(
				delta.Difference{
					FieldPath: FieldPathSystem,
					Type:      delta.DifferenceTypeRemove,
					From:      thisSystemUUID,
					To:        nil,
				},
			)
		} else {
			otherSystemUUID := otherSystem.UUID()
			if thisSystemUUID != otherSystem.UUID() {
				del.Push(
					delta.Difference{
						FieldPath: FieldPathSystem,
						Type:      delta.DifferenceTypeModify,
						From:      thisSystemUUID,
						To:        otherSystemUUID,
					},
				)
			}
		}
	} else if otherSystem != nil {
		otherSystemUUID := otherSystem.UUID()
		del.Push(
			delta.Difference{
				FieldPath: FieldPathSystem,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        otherSystemUUID,
			},
		)
	}

	thisUUID := d.uuid
	otherUUID := other.UUID()
	if thisUUID != otherUUID {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathUUID,
				Type:      delta.DifferenceTypeModify,
				From:      string(thisUUID),
				To:        string(otherUUID),
			},
		)
	}

	thisName := d.name
	otherName := other.Name()
	if thisName != otherName {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathName,
				Type:      delta.DifferenceTypeModify,
				From:      string(thisName),
				To:        string(otherName),
			},
		)
	}

	thisRoot := d.parent
	otherRoot := other.Root()
	if thisRoot != nil {
		thisRootUUID := thisRoot.UUID()
		if otherRoot == nil {
			del.Push(
				delta.Difference{
					FieldPath: FieldPathRoot,
					Type:      delta.DifferenceTypeRemove,
					From:      thisRootUUID,
					To:        nil,
				},
			)
		} else {
			otherRootUUID := otherRoot.UUID()
			if thisRootUUID != otherRoot.UUID() {
				del.Push(
					delta.Difference{
						FieldPath: FieldPathRoot,
						Type:      delta.DifferenceTypeModify,
						From:      thisRootUUID,
						To:        otherRootUUID,
					},
				)
			}
		}
	} else if otherRoot != nil {
		otherRootUUID := otherRoot.UUID()
		del.Push(
			delta.Difference{
				FieldPath: FieldPathRoot,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        otherRootUUID,
			},
		)
	}

	thisParent := d.parent
	otherParent := other.Parent()
	if thisParent != nil {
		thisParentUUID := thisParent.UUID()
		if otherParent == nil {
			del.Push(
				delta.Difference{
					FieldPath: FieldPathParent,
					Type:      delta.DifferenceTypeRemove,
					From:      thisParentUUID,
					To:        nil,
				},
			)
		} else {
			otherParentUUID := otherParent.UUID()
			if thisParentUUID != otherParent.UUID() {
				del.Push(
					delta.Difference{
						FieldPath: FieldPathParent,
						Type:      delta.DifferenceTypeModify,
						From:      thisParentUUID,
						To:        otherParentUUID,
					},
				)
			}
		}
	} else if otherParent != nil {
		otherParentUUID := otherParent.UUID()
		del.Push(
			delta.Difference{
				FieldPath: FieldPathParent,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        otherParentUUID,
			},
		)
	}
	return del, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the Domain as a
// new Domain (i.e. for the first generation)
func (d Domain) diffNew() (*delta.Delta, error) {
	del := &delta.Delta{}

	if d.system != nil {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathSystem,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        d.system.UUID(),
			},
		)
	}
	del.Push(
		delta.Difference{
			FieldPath: FieldPathUUID,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        d.uuid,
		},
	)
	del.Push(
		delta.Difference{
			FieldPath: FieldPathName,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        string(d.name),
		},
	)
	if d.root != nil {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathRoot,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        d.root.UUID(),
			},
		)
	}
	if d.parent != nil {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathParent,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        d.parent.UUID(),
			},
		)
	}
	return del, nil
}
