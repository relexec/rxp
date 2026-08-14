package api

import (
	"github.com/relexec/rxp/errors"
)

type Kind struct {
	// System contains the System containing the Kind. This is a pointer to a
	// System to allow for backends to default missing System information to
	// their host system.
	System *System
	// UUID stores the Kind's globally-unique identifier.
	UUID string
	// Name is the name of the Kind.
	Name KindName
	// Scope is the uniqueness constraint of the names of Objects having this
	// Kind.
	Scope Scope
}

// Validate returns an error if the Kind is not valid.
func (k Kind) Validate() error {
	if k.UUID == "" {
		return errors.ErrKindUUIDRequired
	}
	err := k.Name.Validate()
	if err != nil {
		return err
	}
	if k.System != nil {
		return k.System.Validate()
	}
	return nil
}
