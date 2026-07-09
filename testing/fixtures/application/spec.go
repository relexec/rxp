package application

import (
	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"
)

var (
	FieldPathDescription = fieldpath.FromString("description")
	FieldPathType        = fieldpath.FromString("type")
	FieldPathOwner       = fieldpath.FromString("owner")
)

type Spec_V1_0_0 struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func (s Spec_V1_0_0) Diff(subject any) (*delta.Delta, error) {
	var other Spec_V1_0_0
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return s.diffNew()
	case Spec_V1_0_0:
		other = subject
	case *Spec_V1_0_0:
		other = *subject
	default:
		return nil, delta.CannotCompareTypes(s, subject)
	}

	d := &delta.Delta{}

	if s.Description != other.Description {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathDescription,
				Type:      delta.DifferenceTypeModify,
				From:      s.Description,
				To:        other.Description,
			},
		)
	}
	if s.Type != other.Type {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathType,
				Type:      delta.DifferenceTypeModify,
				From:      s.Type,
				To:        other.Type,
			},
		)
	}
	return d, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the
// Spec_V1_0_0 as a new Spec_V1_0_0 (i.e. for the first generation)
func (s Spec_V1_0_0) diffNew() (*delta.Delta, error) {
	d := &delta.Delta{}
	d.Push(
		delta.Difference{
			FieldPath: FieldPathDescription,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.Description,
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathType,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.Type,
		},
	)
	return d, nil
}

type Spec_V1_0_1 struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Owner       string `json:"owner"`
}

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func (s Spec_V1_0_1) Diff(subject any) (*delta.Delta, error) {
	var other Spec_V1_0_1
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return s.diffNew()
	case Spec_V1_0_1:
		other = subject
	case *Spec_V1_0_1:
		other = *subject
	default:
		return nil, delta.CannotCompareTypes(s, subject)
	}

	d := &delta.Delta{}

	if s.Description != other.Description {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathDescription,
				Type:      delta.DifferenceTypeModify,
				From:      s.Description,
				To:        other.Description,
			},
		)
	}
	if s.Type != other.Type {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathType,
				Type:      delta.DifferenceTypeModify,
				From:      s.Type,
				To:        other.Type,
			},
		)
	}
	if s.Owner != other.Owner {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathOwner,
				Type:      delta.DifferenceTypeModify,
				From:      s.Owner,
				To:        other.Owner,
			},
		)
	}
	return d, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the
// Spec_V1_0_1 as a new Spec_V1_0_1 (i.e. for the first generation)
func (s Spec_V1_0_1) diffNew() (*delta.Delta, error) {
	d := &delta.Delta{}
	d.Push(
		delta.Difference{
			FieldPath: FieldPathDescription,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.Description,
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathType,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.Type,
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathOwner,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.Owner,
		},
	)
	return d, nil
}
