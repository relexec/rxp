package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/namespace"
	"github.com/samber/lo"
)

type NamespaceNamePredicate struct {
	BasePredicate
}

func (p NamespaceNamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
	switch v := v.(type) {
	case []api.NamespaceName:
		for _, dn := range v {
			if err := dn.Validate(); err != nil {
				return errors.PredicateInvalid(err.Error())
			}
		}
	case api.NamespaceName:
		return v.Validate()
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
	return nil
}

// NamespaceNameEqual returns an Expression that will match things having a
// particular NamespaceName.
func NamespaceNameEqual(name api.NamespaceName) Expression {
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: name,
			},
		},
	}
}

// NamespaceNameNotEqual returns an Expression that will match things not having a
// particular NamespaceName.
func NamespaceNameNotEqual(name api.NamespaceName) Expression {
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   name,
			},
		},
	}
}

// NamespaceNameIn returns an Expression that will match things that have any of a
// set of specified NamespaceNames.
func NamespaceNameIn(names ...api.NamespaceName) Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NamespaceNameEqual(names[0])
	}
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: names,
			},
		},
	}
}

// NamespaceNameNotIn returns an Expression that will match things that do not
// have any of a set of specified NamespaceNames.
func NamespaceNameNotIn(names ...api.NamespaceName) Expression {
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   names,
			},
		},
	}
}

type NamespaceUUIDPredicate struct {
	BasePredicate
}

// NamespaceUUIDEqual returns an Expression that will match things having a
// particular NamespaceUUID.
func NamespaceUUIDEqual(uuid string) Expression {
	return UnaryExpression{
		NamespaceUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: uuid,
			},
		},
	}
}

// NamespaceUUIDNotEqual returns an Expression that will match things not having a
// particular NamespaceUUID.
func NamespaceUUIDNotEqual(uuid string) Expression {
	return UnaryExpression{
		NamespaceUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   uuid,
			},
		},
	}
}

// NamespaceUUIDIn returns an Expression that will match things that have any of a
// set of specified NamespaceUUIDs.
func NamespaceUUIDIn(uuids ...string) Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return NamespaceUUIDEqual(uuids[0])
	}
	return UnaryExpression{
		NamespaceUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: uuids,
			},
		},
	}
}

// NamespaceUUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified NamespaceUUIDs.
func NamespaceUUIDNotIn(uuids ...string) Expression {
	return UnaryExpression{
		NamespaceUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   uuids,
			},
		},
	}
}

type NamespacePredicate struct {
	BasePredicate
}

// NamespaceEqual returns an Expression that will match things having a particular
// Namespace.
func NamespaceEqual(ns *namespace.Namespace) Expression {
	if ns.UUID() != "" {
		return NamespaceUUIDEqual(ns.UUID())
	}
	return UnaryExpression{
		NamespacePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: ns,
			},
		},
	}
}

// NamespaceNotEqual returns an Expression that will match things not having a
// particular Namespace.
func NamespaceNotEqual(ns *namespace.Namespace) Expression {
	if ns.UUID() != "" {
		return NamespaceUUIDNotEqual(ns.UUID())
	}
	return UnaryExpression{
		NamespacePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   ns,
			},
		},
	}
}

// NamespaceIn returns an Expression that will match things that have any of a set
// of specified Namespace.
func NamespaceIn(nss ...*namespace.Namespace) Expression {
	uuids := lo.Map(nss, func(ns *namespace.Namespace, _ int) string {
		return ns.UUID()
	})
	if !lo.Contains(uuids, "") {
		return NamespaceUUIDIn(uuids...)
	}
	exprs := make([]Expression, 0, len(nss))
	for _, ns := range nss {
		exprs = append(exprs, NamespaceEqual(ns))
	}
	return Or(exprs...)
}

// NamespaceNotIn returns an Expression that will match things that do not have
// any of a set of specified Namespace.
func NamespaceNotIn(nss ...*namespace.Namespace) Expression {
	uuids := lo.Map(nss, func(ns *namespace.Namespace, _ int) string {
		return ns.UUID()
	})
	if !lo.Contains(uuids, "") {
		return NamespaceUUIDNotIn(uuids...)
	}
	exprs := make([]Expression, 0, len(nss))
	for _, ns := range nss {
		exprs = append(exprs, NamespaceNotEqual(ns))
	}
	return And(exprs...)
}
