package api

import (
	"github.com/relexec/rxp/errors"
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
