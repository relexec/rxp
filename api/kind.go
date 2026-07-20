package api

import (
	"github.com/relexec/rxp/errors"
)

type Kind struct {
	// system contains the System containing the Kind.
	system *System
	// uuid stores the Kind's globally-unique identifier.
	uuid string
	// name is the name of the Kind.
	name KindName
	// scope is the uniqueness constraint of the names of Objects having this
	// Kind.
	scope Scope
}

// Validate returns an error if the Kind is not valid.
func (k Kind) Validate() error {
	if k.uuid == "" {
		return errors.ErrKindUUIDRequired
	}
	err := k.name.Validate()
	if err != nil {
		return err
	}
	if k.system != nil {
		return k.system.Validate()
	}
	return nil
}

// System returns the System of the Kind.
func (k Kind) System() *System {
	return k.system
}

// SetSystem sets the System of Kind.
func (k *Kind) SetSystem(system *System) {
	k.system = system
}

// UUID returns the globally-unique identifier of the Kind.
func (k Kind) UUID() string {
	return k.uuid
}

// SetUUID sets the globally-unique identifier of the Kind.
func (k *Kind) SetUUID(uuid string) {
	k.uuid = uuid
}

// Name returns the name of the Kind.
func (k Kind) Name() KindName {
	return k.name
}

// SetName sets the Name of the Kind.
func (k *Kind) SetName(name KindName) {
	k.name = name
}

// Scope returns the name uniqueness constraint for Objects having this
// Kind.
func (k Kind) Scope() Scope {
	return k.scope
}

// SetScope sets the name uniqueness constraint for Objects having this
// Kind.
func (k *Kind) SetScope(scope Scope) {
	k.scope = scope
}
