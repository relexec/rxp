package api

import (
	"github.com/relexec/rxp/errors"
)

// System represents the boundaries of an rxp system installation.
type System struct {
	// UUID contains the System's globally-unique identifier.
	UUID string `json:"uuid"`
	// Tag contains an optional string tag for the System. Note this is not
	// called "name" because a Name in rxp has a specific semantic meaning that
	// reflects the uniqueness constraint its value. Tags have no such
	// uniqueness constraint.
	Tag string `json:"tag,omitempty"`
}

// Validate returns an error if the System is invalid.
func (s System) Validate() error {
	if s.UUID == "" {
		return errors.ErrSystemUUIDRequired
	}
	return nil
}
