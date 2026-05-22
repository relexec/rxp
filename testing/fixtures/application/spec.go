package application

import (
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/types"
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

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (s Spec_V1_0_0) Diff(subject any) (*cmp.Delta, error) {
	var other Spec_V1_0_0
	switch subject := subject.(type) {
	case cmp.ZeroGen:
		return s.diffNew()
	case Spec_V1_0_0:
		other = subject
	case *Spec_V1_0_0:
		other = *subject
	default:
		return nil, cmp.CannotCompareTypes(s, subject)
	}

	d := &cmp.Delta{}

	if s.Description != other.Description {
		d.Push(
			cmp.NewDifference(
				FieldPathDescription,
				cmp.DifferenceTypeModify,
				s.Description,
				other.Description,
			),
		)
	}
	if s.Type != other.Type {
		d.Push(
			cmp.NewDifference(
				FieldPathType,
				cmp.DifferenceTypeModify,
				s.Type,
				other.Type,
			),
		)
	}
	return d, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the
// Spec_V1_0_0 as a new Spec_V1_0_0 (i.e. for the first generation)
func (s Spec_V1_0_0) diffNew() (*cmp.Delta, error) {
	d := &cmp.Delta{}
	d.Push(
		cmp.NewDifference(
			FieldPathDescription,
			cmp.DifferenceTypeAdd,
			nil,
			s.Description,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathType,
			cmp.DifferenceTypeAdd,
			nil,
			s.Type,
		),
	)
	return d, nil
}

var _ types.Spec = (*Spec_V1_0_0)(nil)

type Spec_V1_0_1 struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Owner       string `json:"owner"`
}

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (s Spec_V1_0_1) Diff(subject any) (*cmp.Delta, error) {
	var other Spec_V1_0_1
	switch subject := subject.(type) {
	case cmp.ZeroGen:
		return s.diffNew()
	case Spec_V1_0_1:
		other = subject
	case *Spec_V1_0_1:
		other = *subject
	default:
		return nil, cmp.CannotCompareTypes(s, subject)
	}

	d := &cmp.Delta{}

	if s.Description != other.Description {
		d.Push(
			cmp.NewDifference(
				FieldPathDescription,
				cmp.DifferenceTypeModify,
				s.Description,
				other.Description,
			),
		)
	}
	if s.Type != other.Type {
		d.Push(
			cmp.NewDifference(
				FieldPathType,
				cmp.DifferenceTypeModify,
				s.Type,
				other.Type,
			),
		)
	}
	if s.Owner != other.Owner {
		d.Push(
			cmp.NewDifference(
				FieldPathOwner,
				cmp.DifferenceTypeModify,
				s.Owner,
				other.Owner,
			),
		)
	}
	return d, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the
// Spec_V1_0_1 as a new Spec_V1_0_1 (i.e. for the first generation)
func (s Spec_V1_0_1) diffNew() (*cmp.Delta, error) {
	d := &cmp.Delta{}
	d.Push(
		cmp.NewDifference(
			FieldPathDescription,
			cmp.DifferenceTypeAdd,
			nil,
			s.Description,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathType,
			cmp.DifferenceTypeAdd,
			nil,
			s.Type,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathOwner,
			cmp.DifferenceTypeAdd,
			nil,
			s.Owner,
		),
	)
	return d, nil
}

var _ types.Spec = (*Spec_V1_0_1)(nil)
