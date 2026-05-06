package v1

import (
	"time"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathAuthor      = fieldpath.FromString("author")
	FieldPathPublishedOn = fieldpath.FromString("published_on")
	FieldPathNumPages    = fieldpath.FromString("num_pages")
)

type Spec_V1_0_0 struct {
	Author      string    `json:"author"`
	PublishedOn time.Time `json:"published_on"`
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

	if s.Author != other.Author {
		d.Push(
			cmp.NewDifference(
				FieldPathAuthor,
				cmp.DifferenceTypeModify,
				s.Author,
				other.Author,
			),
		)
	}
	if !s.PublishedOn.Equal(other.PublishedOn) {
		d.Push(
			cmp.NewDifference(
				FieldPathPublishedOn,
				cmp.DifferenceTypeModify,
				s.PublishedOn,
				other.PublishedOn,
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
			FieldPathAuthor,
			cmp.DifferenceTypeAdd,
			nil,
			s.Author,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathPublishedOn,
			cmp.DifferenceTypeAdd,
			nil,
			s.PublishedOn,
		),
	)
	return d, nil
}

var _ types.Spec = (*Spec_V1_0_0)(nil)

type Spec_V1_0_1 struct {
	Author      string    `json:"author"`
	PublishedOn time.Time `json:"published_on"`
	NumPages    int       `json:"num_pages"`
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

	if s.Author != other.Author {
		d.Push(
			cmp.NewDifference(
				FieldPathAuthor,
				cmp.DifferenceTypeModify,
				s.Author,
				other.Author,
			),
		)
	}
	if !s.PublishedOn.Equal(other.PublishedOn) {
		d.Push(
			cmp.NewDifference(
				FieldPathPublishedOn,
				cmp.DifferenceTypeModify,
				s.PublishedOn,
				other.PublishedOn,
			),
		)
	}
	if s.NumPages != other.NumPages {
		d.Push(
			cmp.NewDifference(
				FieldPathNumPages,
				cmp.DifferenceTypeModify,
				s.NumPages,
				other.NumPages,
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
			FieldPathAuthor,
			cmp.DifferenceTypeAdd,
			nil,
			s.Author,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathPublishedOn,
			cmp.DifferenceTypeAdd,
			nil,
			s.PublishedOn,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathNumPages,
			cmp.DifferenceTypeAdd,
			nil,
			s.NumPages,
		),
	)
	return d, nil
}

var _ types.Spec = (*Spec_V1_0_1)(nil)
