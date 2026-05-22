package kind

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/system"
)

var (
	FieldPathName  = fieldpath.FromString("name")
	FieldPathScope = fieldpath.FromString("scope")
)

type Kind struct {
	// system contains the System containing the Kind.
	system *system.System
	// name is the name of the Kind.
	name api.KindName
	// scope is the uniqueness constraint of the names of Objects having this
	// Kind.
	scope api.Scope
}

// Validate returns an error if the Kind is not valid.
func (k Kind) Validate() error {
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

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (k Kind) Diff(subject any) (*cmp.Delta, error) {
	var other *Kind
	switch subject := subject.(type) {
	case cmp.ZeroGen:
		return k.diffNew()
	case Kind:
		other = &subject
	case *Kind:
		other = subject
	default:
		return nil, cmp.CannotCompareTypes(k, subject)
	}

	d := &cmp.Delta{}

	thisName := string(k.name)
	otherName := string(other.Name())
	if thisName != otherName {
		d.Push(
			cmp.NewDifference(
				FieldPathName,
				cmp.DifferenceTypeModify,
				thisName,
				otherName,
			),
		)
	}
	if k.scope != other.Scope() {
		d.Push(
			cmp.NewDifference(
				FieldPathScope,
				cmp.DifferenceTypeModify,
				k.scope,
				other.Scope(),
			),
		)
	}
	return d, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the Kind as a
// new Kind (i.e. for the first generation)
func (k Kind) diffNew() (*cmp.Delta, error) {
	d := &cmp.Delta{}
	d.Push(
		cmp.NewDifference(
			FieldPathName,
			cmp.DifferenceTypeAdd,
			string(k.name),
			nil,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathScope,
			cmp.DifferenceTypeAdd,
			k.scope,
			nil,
		),
	)
	return d, nil
}
