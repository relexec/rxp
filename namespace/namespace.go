package namespace

import (
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathDomain = fieldpath.FromString("domain")
	FieldPathUUID   = fieldpath.FromString("uuid")
	FieldPathName   = fieldpath.FromString("name")
)

// Namespace describes a logical division within a Domain.
type Namespace struct {
	// domain contains the Namespace's Domain.
	domain types.Domain
	// uuid stores the Namespace's globally-unique identifier.
	uuid string
	// name contains the Namespace name.
	//
	// A valid Namespace Name is a DNS-formatted (RFC 1035-compliant) name less than
	// 254 characters.
	//
	// A Namespace's Name must be unique within the scope of its Domain.
	name types.NamespaceName
}

// Validate returns an error if the Domain is invalid.
func (n Namespace) Validate() error {
	if n.domain == nil {
		return errors.ErrNamespaceDomainRequired
	}
	if err := n.domain.Validate(); err != nil {
		return err
	}
	if n.uuid == "" {
		return errors.ErrNamespaceUUIDRequired
	}
	return n.name.Validate()
}

// Domain returns the Domain of the Namespace.
func (n Namespace) Domain() types.Domain {
	return n.domain
}

// SetDomain sets the Domain of Namespace.
func (n *Namespace) SetDomain(domain types.Domain) {
	n.domain = domain
}

// UUID returns the globally-unique identifier of the Namespace.
func (n Namespace) UUID() string {
	return n.uuid
}

// SetUUID sets the globally-unique identifier of the Namespace.
func (n *Namespace) SetUUID(uuid string) {
	n.uuid = uuid
}

// Name returns the Name of the Namespace.
func (n Namespace) Name() types.NamespaceName {
	return n.name
}

// SetName sets the Name of Namespace.
func (n *Namespace) SetName(name types.NamespaceName) {
	n.name = name
}

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (n Namespace) Diff(subject any) (*cmp.Delta, error) {
	var other types.Namespace
	switch subject := subject.(type) {
	case cmp.ZeroGen:
		return n.diffNew()
	case Namespace:
		other = &subject
	case *Namespace:
		other = subject
	default:
		return nil, cmp.CannotCompareTypes(n, subject)
	}

	delta := &cmp.Delta{}

	thisDomain := n.domain
	otherDomain := other.Domain()
	if thisDomain != nil {
		thisDomainName := n.domain.Name()
		if otherDomain == nil {
			delta.Push(
				cmp.NewDifference(
					FieldPathDomain,
					cmp.DifferenceTypeRemove,
					thisDomainName,
					nil,
				),
			)
		} else {
			otherDomainName := otherDomain.Name()
			if thisDomainName != otherDomain.Name() {
				delta.Push(
					cmp.NewDifference(
						FieldPathDomain,
						cmp.DifferenceTypeModify,
						thisDomainName,
						otherDomainName,
					),
				)
			}
		}
	} else if otherDomain != nil {
		otherDomainName := otherDomain.Name()
		delta.Push(
			cmp.NewDifference(
				FieldPathDomain,
				cmp.DifferenceTypeAdd,
				nil,
				otherDomainName,
			),
		)
	}

	thisUUID := n.uuid
	otherUUID := other.UUID()
	if thisUUID != otherUUID {
		delta.Push(
			cmp.NewDifference(
				FieldPathUUID,
				cmp.DifferenceTypeModify,
				string(thisUUID),
				string(otherUUID),
			),
		)
	}

	thisName := n.name
	otherName := other.Name()
	if thisName != otherName {
		delta.Push(
			cmp.NewDifference(
				FieldPathName,
				cmp.DifferenceTypeModify,
				string(thisName),
				string(otherName),
			),
		)
	}
	return delta, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the Namespace
// as a new Namespace (i.e. for the first generation)
func (n Namespace) diffNew() (*cmp.Delta, error) {
	delta := &cmp.Delta{}

	if n.domain != nil {
		delta.Push(
			cmp.NewDifference(
				FieldPathDomain,
				cmp.DifferenceTypeAdd,
				n.domain.Name(),
				nil,
			),
		)
	}
	delta.Push(
		cmp.NewDifference(
			FieldPathUUID,
			cmp.DifferenceTypeAdd,
			n.uuid,
			nil,
		),
	)
	delta.Push(
		cmp.NewDifference(
			FieldPathName,
			cmp.DifferenceTypeAdd,
			string(n.name),
			nil,
		),
	)
	return delta, nil
}

var _ types.Namespace = (*Namespace)(nil)
