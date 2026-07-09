package kind

import (
	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/system"
)

var (
	FieldPathSystem = fieldpath.FromString("system")
	FieldPathUUID   = fieldpath.FromString("uuid")
	FieldPathName   = fieldpath.FromString("name")
	FieldPathScope  = fieldpath.FromString("scope")
)

type Kind struct {
	// system contains the System containing the Kind.
	system *system.System
	// uuid stores the Kind's globally-unique identifier.
	uuid string
	// name is the name of the Kind.
	name api.KindName
	// scope is the uniqueness constraint of the names of Objects having this
	// Kind.
	scope api.Scope
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
func (k Kind) System() *system.System {
	return k.system
}

// SetSystem sets the System of Kind.
func (k *Kind) SetSystem(system *system.System) {
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
func (k Kind) Name() api.KindName {
	return k.name
}

// SetName sets the Name of the Kind.
func (k *Kind) SetName(name api.KindName) {
	k.name = name
}

// Scope returns the name uniqueness constraint for Objects having this
// Kind.
func (k Kind) Scope() api.Scope {
	return k.scope
}

// SetScope sets the name uniqueness constraint for Objects having this
// Kind.
func (k *Kind) SetScope(scope api.Scope) {
	k.scope = scope
}

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func (k Kind) Diff(subject any) (*delta.Delta, error) {
	var other *Kind
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return k.diffNew()
	case Kind:
		other = &subject
	case *Kind:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(k, subject)
	}

	d := &delta.Delta{}

	thisSystem := k.system
	otherSystem := other.System()
	if thisSystem != nil {
		thisSystemUUID := k.system.UUID()
		if otherSystem == nil {
			d.Push(
				delta.Difference{
					FieldPath: FieldPathSystem,
					Type:      delta.DifferenceTypeRemove,
					From:      thisSystemUUID,
					To:        nil,
				},
			)
		} else {
			otherSystemUUID := otherSystem.UUID()
			if thisSystemUUID != otherSystem.UUID() {
				d.Push(
					delta.Difference{
						FieldPath: FieldPathSystem,
						Type:      delta.DifferenceTypeModify,
						From:      thisSystemUUID,
						To:        otherSystemUUID,
					},
				)
			}
		}
	} else if otherSystem != nil {
		otherSystemUUID := otherSystem.UUID()
		d.Push(
			delta.Difference{
				FieldPath: FieldPathSystem,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        otherSystemUUID,
			},
		)
	}

	thisUUID := k.uuid
	otherUUID := other.UUID()
	if thisUUID != otherUUID {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathUUID,
				Type:      delta.DifferenceTypeModify,
				From:      string(thisUUID),
				To:        string(otherUUID),
			},
		)
	}

	thisName := string(k.name)
	otherName := string(other.Name())
	if thisName != otherName {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathName,
				Type:      delta.DifferenceTypeModify,
				From:      thisName,
				To:        otherName,
			},
		)
	}
	if k.scope != other.Scope() {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathScope,
				Type:      delta.DifferenceTypeModify,
				From:      k.scope,
				To:        other.Scope(),
			},
		)
	}
	return d, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the Kind as a
// new Kind (i.e. for the first generation)
func (k Kind) diffNew() (*delta.Delta, error) {
	d := &delta.Delta{}
	if k.system != nil {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathSystem,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        k.system.UUID(),
			},
		)
	}
	d.Push(
		delta.Difference{
			FieldPath: FieldPathUUID,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        k.uuid,
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathName,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        string(k.name),
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathScope,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        k.scope,
		},
	)
	return d, nil
}
