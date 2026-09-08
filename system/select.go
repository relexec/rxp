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

// SelectOption modifies the [Selector] returned from [Select].
type SelectOption func(*Selector)

// ByUUID sets the Selector's UUID.
func ByUUID(uuid string) SelectOption {
	return func(s *Selector) {
		s.uuid = uuid
	}
}

// Select returns a new [Selector]
func Select(opts ...SelectOption) Selector {
	s := Selector{}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}
