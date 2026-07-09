package system

import (
	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"
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

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func (s System) Diff(subject any) (*delta.Delta, error) {
	var other *System
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return s.diffNew()
	case System:
		other = &subject
	case *System:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(s, subject)
	}

	d := &delta.Delta{}

	thisSystemUUID := s.uuid
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

	thisSystemTag := s.tag
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
func (s System) diffNew() (*delta.Delta, error) {
	d := &delta.Delta{}

	d.Push(
		delta.Difference{
			FieldPath: FieldPathUUID,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.uuid,
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathTag,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.tag,
		},
	)
	return d, nil
}
