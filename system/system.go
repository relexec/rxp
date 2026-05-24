package system

import (
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/errors"
)

var (
	FieldPathUUID = fieldpath.FromString("uuid")
	FieldPathTag  = fieldpath.FromString("tag")
)

// System represents the boundaries of an rxp system installation.
type System struct {
	// uuid contains the System's globally-unique identifier.
	uuid string
	// tag contains an optional string tag for the System. Note this is not
	// called "name" because a Name in rxp has a specific semantic meaning that
	// reflects the uniqueness constraint its value. Tags have no such
	// uniqueness constraint.
	tag string
}

// Validate returns an error if the System is invalid.
func (s System) Validate() error {
	if s.uuid == "" {
		return errors.ErrSystemUUIDRequired
	}
	return nil
}

// UUID returns the globally-unique identifier of the System.
func (s System) UUID() string {
	return s.uuid
}

// SetUUID sets the globally-unique identifier of the System.
func (s *System) SetUUID(uuid string) {
	s.uuid = uuid
}

// Tag returns an optional string tag for the System. Note this is not called
// "name" because a Name in rxp has a specific semantic meaning that reflects
// the uniqueness constraint its value. Tags have no such uniqueness
// constraint.
func (s System) Tag() string {
	return s.tag
}

// SetTag sets the optional string tag for the System.
func (s *System) SetTag(tag string) {
	s.tag = tag
}

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (s System) Diff(subject any) (*cmp.Delta, error) {
	var other *System
	switch subject := subject.(type) {
	case cmp.ZeroGen:
		return s.diffNew()
	case System:
		other = &subject
	case *System:
		other = subject
	default:
		return nil, cmp.CannotCompareTypes(s, subject)
	}

	delta := &cmp.Delta{}

	thisSystemUUID := s.uuid
	otherSystemUUID := other.UUID()
	if thisSystemUUID != otherSystemUUID {
		delta.Push(
			cmp.NewDifference(
				FieldPathUUID,
				cmp.DifferenceTypeModify,
				thisSystemUUID,
				otherSystemUUID,
			),
		)
	}

	thisSystemTag := s.tag
	otherSystemTag := other.Tag()
	if thisSystemTag != otherSystemTag {
		delta.Push(
			cmp.NewDifference(
				FieldPathTag,
				cmp.DifferenceTypeModify,
				thisSystemTag,
				otherSystemTag,
			),
		)
	}
	return delta, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the System as a
// new System (i.e. for the first generation)
func (s System) diffNew() (*cmp.Delta, error) {
	delta := &cmp.Delta{}

	delta.Push(
		cmp.NewDifference(
			FieldPathUUID,
			cmp.DifferenceTypeAdd,
			s.uuid,
			nil,
		),
	)
	delta.Push(
		cmp.NewDifference(
			FieldPathTag,
			cmp.DifferenceTypeAdd,
			s.tag,
			nil,
		),
	)
	return delta, nil
}
