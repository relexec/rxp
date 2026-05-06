package v1

import (
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathName      = fieldpath.FromString("name")
	FieldPathPublisher = fieldpath.FromString("publisher")
)

type Spec_V1_0_0 struct {
	Name string `json:"name"`
}

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

	if s.Name != other.Name {
		d.Push(
			cmp.NewDifference(
				FieldPathName,
				cmp.DifferenceTypeModify,
				s.Name,
				other.Name,
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
			FieldPathName,
			cmp.DifferenceTypeAdd,
			nil,
			s.Name,
		),
	)
	return d, nil
}

var _ types.Spec = (*Spec_V1_0_0)(nil)

type Spec_V1_0_1 struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}

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

	if s.Name != other.Name {
		d.Push(
			cmp.NewDifference(
				FieldPathName,
				cmp.DifferenceTypeModify,
				s.Name,
				other.Name,
			),
		)
	}
	if s.Publisher != other.Publisher {
		d.Push(
			cmp.NewDifference(
				FieldPathPublisher,
				cmp.DifferenceTypeModify,
				s.Publisher,
				other.Publisher,
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
			FieldPathName,
			cmp.DifferenceTypeAdd,
			nil,
			s.Name,
		),
	)
	return d, nil
}

var _ types.Spec = (*Spec_V1_0_1)(nil)
