package kind

import (
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathName      = fieldpath.FromString("name")
	FieldPathNamescope = fieldpath.FromString("namescope")
)

type Kind struct {
	// system contains the System containing the Kind.
	system types.System
	// name is the name of the Kind.
	name types.KindName
	// namescope is the uniqueness constraint of the names of Objects having
	// this Kind.
	namescope types.Namescope
}

// Validate returns an error if the Kind is not valid.
func (k Kind) Validate() error {
	err := k.name.Validate()
	if err != nil {
		return err
	}
	return k.namescope.Validate()
}

// System returns the System of the Kind.
func (k Kind) System() types.System {
	return k.system
}

// SetSystem sets the System of Kind.
func (k *Kind) SetSystem(system types.System) {
	k.system = system
}

// Name returns the name of the Kind.
func (k Kind) Name() types.KindName {
	return k.name
}

// SetName sets the Name of the Kind.
func (k *Kind) SetName(name types.KindName) {
	k.name = name
}

// Namescope returns the name uniqueness constraint for Objects having this
// KindVersion.
func (k Kind) Namescope() types.Namescope {
	return k.namescope
}

// SetNamescope sets the name uniqueness constraint for Objects having this
// KindVersion.
func (k *Kind) SetNamescope(namescope types.Namescope) {
	k.namescope = namescope
}

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (k Kind) Diff(subject any) (*cmp.Delta, error) {
	var other types.Kind
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
	if k.namescope != other.Namescope() {
		d.Push(
			cmp.NewDifference(
				FieldPathNamescope,
				cmp.DifferenceTypeModify,
				k.namescope,
				other.Namescope(),
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
			FieldPathNamescope,
			cmp.DifferenceTypeAdd,
			k.namescope,
			nil,
		),
	)
	return d, nil
}

var _ types.Kind = (*Kind)(nil)
