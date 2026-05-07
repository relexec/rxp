package system

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
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
		return errors.ErrSystemUUIDEmpty
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

var _ types.System = (*System)(nil)
