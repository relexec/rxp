package system

import (
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/errors"
)

var (
	FieldPathUUID = fieldpath.FromString("uuid")
	FieldPathName = fieldpath.FromString("name")
)

// System represents the boundaries of an rxp system installation.
type System struct {
	// uuid contains the System's globally-unique identifier.
	uuid string
	// name contains the optional human-readable System name.
	name string
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

// Name returns the optional human-readable name of the System.
func (s System) Name() string {
	return s.name
}

// SetName sets the optional human-readable name of the System.
func (s *System) SetName(name string) {
	s.name = name
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

	thisSystemName := s.name
	otherSystemName := other.Name()
	if thisSystemName != otherSystemName {
		delta.Push(
			cmp.NewDifference(
				FieldPathName,
				cmp.DifferenceTypeModify,
				thisSystemName,
				otherSystemName,
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
			FieldPathName,
			cmp.DifferenceTypeAdd,
			s.name,
			nil,
		),
	)
	return delta, nil
}
