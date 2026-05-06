package selector

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Selector allows a Reader to select a specific target.
type Selector struct {
	// id is the globally-unique string identifier to look up the Object for.
	id string
	// domain is the Domain to use when looking up the Object via name.
	domain string
	// namespace is the Namespace to use when looking up the Object via name.
	namespace string
	// name is the Name to use when looking up the Object via name.
	name string
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	if s.id == "" && s.name == "" {
		return errors.ErrIDOrNameRequired
	}
	return nil
}

// ID returns the globally-unique string identifier to look up the Object for.
func (s Selector) ID() string {
	return s.id
}

// Domain returns the Domain to use when looking up the Object via name.
func (s Selector) Domain() string {
	return s.domain
}

// Namespace returns the Namespace to use when looking up the Object via name.
func (s Selector) Namespace() string {
	return s.namespace
}

// Name returns the Name to use when looking up the Object via name.
func (s Selector) Name() string {
	return s.name
}

var _ types.Selector = (*Selector)(nil)
