package system

import (
	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"

	"github.com/relexec/rxp/api"
)

var (
	FieldPathUUID = fieldpath.FromString("uuid")
	FieldPathTag  = fieldpath.FromString("tag")
)

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func Diff(s api.System, subject any) (*delta.Delta, error) {
	var other *api.System
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return diffNew(s)
	case api.System:
		other = &subject
	case *api.System:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(s, subject)
	}

	d := &delta.Delta{}

	thisSystemUUID := s.UUID()
	otherSystemUUID := other.UUID()
	if thisSystemUUID != otherSystemUUID {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathUUID,
				Type:      delta.DifferenceTypeModify,
				From:      thisSystemUUID,
				To:        otherSystemUUID,
			},
		)
	}

	thisSystemTag := s.Tag()
	otherSystemTag := other.Tag()
	if thisSystemTag != otherSystemTag {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathTag,
				Type:      delta.DifferenceTypeModify,
				From:      thisSystemTag,
				To:        otherSystemTag,
			},
		)
	}
	return d, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the System as a
// new System (i.e. for the first generation)
func diffNew(s api.System) (*delta.Delta, error) {
	d := &delta.Delta{}

	d.Push(
		delta.Difference{
			FieldPath: FieldPathUUID,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.UUID(),
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathTag,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.Tag(),
		},
	)
	return d, nil
}
