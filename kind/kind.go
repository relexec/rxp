package kind

import (
	"github.com/relexec/rxp/core"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/system"
)

type Kind struct {
	core.Record
	// System contains the System containing the Kind. This is a pointer to a
	// System to allow for backends to default missing System information to
	// their host system.
	System *system.System
	// UUID stores the Kind's globally-unique identifier.
	UUID string
	// Name is the name of the Kind.
	Name Name
	// Scope is the uniqueness constraint of the names of Objects having this
	// Kind.
	Scope core.Scope
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
