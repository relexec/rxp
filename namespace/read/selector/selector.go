package selector

import (
	"github.com/relexec/rxp/types"
)

// Selector allows a NamespaceReader to select a specific target.
type Selector struct {
	// domain is the Domain to look up the Namespace for.
	domain types.Domain
	// name is the NamespaceName to look up the Namespace for.
	name string
}

// Validate returns an error if the Selector is not valid.
func (s Selector) Validate() error {
	return nil
}

// Domain returns the Domain to use when looking up the Namespace.
func (s Selector) Domain() types.Domain {
	return s.domain
}

// Name returns the name to use when looking up the Namespace.
func (s Selector) Name() string {
	return s.name
}
