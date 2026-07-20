package kind

import (
	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"

	"github.com/relexec/rxp/api"
)

var (
	FieldPathSystem = fieldpath.FromString("system")
	FieldPathUUID   = fieldpath.FromString("uuid")
	FieldPathName   = fieldpath.FromString("name")
	FieldPathScope  = fieldpath.FromString("scope")
)

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func Diff(k api.Kind, subject any) (*delta.Delta, error) {
	var other *api.Kind
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return diffNew(k)
	case api.Kind:
		other = &subject
	case *api.Kind:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(k, subject)
	}

	d := &delta.Delta{}

	thisSystem := k.System()
	otherSystem := other.System()
	if thisSystem != nil {
		thisSystemUUID := k.System().UUID()
		if otherSystem == nil {
			d.Push(
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
				d.Push(
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
		d.Push(
			delta.Difference{
				FieldPath: FieldPathSystem,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        otherSystemUUID,
			},
		)
	}

	thisUUID := k.UUID()
	otherUUID := other.UUID()
	if thisUUID != otherUUID {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathUUID,
				Type:      delta.DifferenceTypeModify,
				From:      string(thisUUID),
				To:        string(otherUUID),
			},
		)
	}

	thisName := string(k.Name())
	otherName := string(other.Name())
	if thisName != otherName {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathName,
				Type:      delta.DifferenceTypeModify,
				From:      thisName,
				To:        otherName,
			},
		)
	}
	if k.Scope() != other.Scope() {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathScope,
				Type:      delta.DifferenceTypeModify,
				From:      k.Scope(),
				To:        other.Scope(),
			},
		)
	}
	return d, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the Kind as a
// new Kind (i.e. for the first generation)
func diffNew(k api.Kind) (*delta.Delta, error) {
	d := &delta.Delta{}
	if k.System() != nil {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathSystem,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        k.System().UUID(),
			},
		)
	}
	d.Push(
		delta.Difference{
			FieldPath: FieldPathUUID,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        k.UUID(),
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathName,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        string(k.Name()),
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathScope,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        k.Scope(),
		},
	)
	return d, nil
}
