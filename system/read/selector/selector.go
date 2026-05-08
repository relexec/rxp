package selector

import (
	"github.com/relexec/rxp/errors"
)

// Selector allows a SystemReader to select a specific target.
type Selector struct {
	// uuid is the globally-unique identifier to look up the System by.
	uuid string
	// name is the human-readable name to look up the System by.
	name string
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid == "" && s.name == "" {
		return errors.ErrSelectorUUIDOrNameRequired
	}
	return nil
}

// UUID is the globally-unique identifier to look up the System by.
func (s Selector) UUID() string {
	return s.uuid
}

// Name returns the optional human-readable name to look up the System by.
func (s Selector) Name() string {
	return s.name
}
