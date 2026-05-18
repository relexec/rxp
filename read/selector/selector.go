package selector

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Selector enables `rxp` backend implementations to read a *single thing*.
//
// Either UUID() or Name() must return a non-empty value.
type Selector struct {
	// uuid is the globally-unique string identifier to look up the target for.
	uuid string
	// name is the Name to use when looking up the target via name.
	name types.Name
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.uuid == "" && s.name == nil {
		return errors.ErrSelectorUUIDOrNameRequired
	}
	if s.name != nil {
		err := s.name.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// UUID returns the globally-unique string identifier to look up the target
// for.
func (s Selector) UUID() string {
	return s.uuid
}

// Namespace returns the Name to use when looking up the target via name.
func (s Selector) Name() types.Name {
	return s.name
}

var _ types.Selector = (*Selector)(nil)
