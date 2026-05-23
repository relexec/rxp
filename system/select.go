package system

import "github.com/relexec/rxp/errors"

// Selector selects a single System.
type Selector struct {
	// uuid is the globally-unique identifier of the System being selected.
	uuid string
}

// UUID returns the globally-unique identifier of the System being selected.
func (s Selector) UUID() string {
	return s.uuid
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid == "" {
		return errors.ErrSelectorUUIDRequired
	}
	return nil
}

// ByUUID returns a Selector that looks up a System having the supplied UUID.
func ByUUID(uuid string) Selector {
	return Selector{uuid: uuid}
}
