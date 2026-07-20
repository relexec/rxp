package domain

import (
	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"

	"github.com/relexec/rxp/api"
)

var (
	FieldPathSystem = fieldpath.FromString("system")
	FieldPathUUID   = fieldpath.FromString("uuid")
	FieldPathName   = fieldpath.FromString("name")
	FieldPathRoot   = fieldpath.FromString("root")
	FieldPathParent = fieldpath.FromString("parent")
)

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func Diff(d api.Domain, subject any) (*delta.Delta, error) {
	var other *api.Domain
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return diffNew(d)
	case api.Domain:
		other = &subject
	case *api.Domain:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(d, subject)
	}

	del := &delta.Delta{}

	thisSystem := d.System()
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

	thisUUID := d.UUID()
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

	thisName := d.Name()
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

	thisRoot := d.Parent()
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

	thisParent := d.Parent()
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
func diffNew(d api.Domain) (*delta.Delta, error) {
	del := &delta.Delta{}

	if d.System() != nil {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathSystem,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        d.System().UUID(),
			},
		)
	}
	del.Push(
		delta.Difference{
			FieldPath: FieldPathUUID,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        d.UUID(),
		},
	)
	del.Push(
		delta.Difference{
			FieldPath: FieldPathName,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        string(d.Name()),
		},
	)
	if d.Root() != nil {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathRoot,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        d.Root().UUID(),
			},
		)
	}
	if d.Parent() != nil {
		del.Push(
			delta.Difference{
				FieldPath: FieldPathParent,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        d.Parent().UUID(),
			},
		)
	}
	return del, nil
}
