package namespace

import (
	"github.com/samber/lo"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/query/expression"
)

type NamePredicate struct {
	expression.BasePredicate
}

func (p NamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value
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

// NameEqual returns an Expression that will match things having a
// particular NamespaceName.
func NameEqual(name api.NamespaceName) expression.Expression {
	return expression.UnaryExpression{
		Predicate: NamePredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match things not having a
// particular NamespaceName.
func NameNotEqual(name api.NamespaceName) expression.Expression {
	return expression.UnaryExpression{
		Predicate: NamePredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorEqual,
				Negated: true,
				Value:   name,
			},
		},
	}
}

// NameIn returns an Expression that will match things that have any of a
// set of specified NamespaceNames.
func NameIn(names ...api.NamespaceName) expression.Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NameEqual(names[0])
	}
	return expression.UnaryExpression{
		Predicate: NamePredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match things that do not
// have any of a set of specified NamespaceNames.
func NameNotIn(names ...api.NamespaceName) expression.Expression {
	return expression.UnaryExpression{
		Predicate: NamePredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorIn,
				Negated: true,
				Value:   names,
			},
		},
	}
}

type UUIDPredicate struct {
	expression.BasePredicate
}

// UUIDEqual returns an Expression that will match things having a
// particular UUID.
func UUIDEqual(uuid string) expression.Expression {
	return expression.UnaryExpression{
		Predicate: UUIDPredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// UUIDNotEqual returns an Expression that will match things not having a
// particular UUID.
func UUIDNotEqual(uuid string) expression.Expression {
	return expression.UnaryExpression{
		Predicate: UUIDPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorEqual,
				Negated: true,
				Value:   uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match things that have any of a
// set of specified UUIDs.
func UUIDIn(uuids ...string) expression.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return UUIDEqual(uuids[0])
	}
	return expression.UnaryExpression{
		Predicate: UUIDPredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified UUIDs.
func UUIDNotIn(uuids ...string) expression.Expression {
	return expression.UnaryExpression{
		Predicate: UUIDPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorIn,
				Negated: true,
				Value:   uuids,
			},
		},
	}
}

type NamespacePredicate struct {
	expression.BasePredicate
}

// Equal returns an Expression that will match things having a particular
// Namespace.
func Equal(ns *Namespace) expression.Expression {
	if ns.UUID() != "" {
		return UUIDEqual(ns.UUID())
	}
	return expression.UnaryExpression{
		Predicate: NamespacePredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorEqual,
				Value: ns,
			},
		},
	}
}

// NotEqual returns an Expression that will match things not having a
// particular Namespace.
func NotEqual(ns *Namespace) expression.Expression {
	if ns.UUID() != "" {
		return UUIDNotEqual(ns.UUID())
	}
	return expression.UnaryExpression{
		Predicate: NamespacePredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorEqual,
				Negated: true,
				Value:   ns,
			},
		},
	}
}

// In returns an Expression that will match things that have any of a set
// of specified Namespace.
func In(nss ...*Namespace) expression.Expression {
	uuids := lo.Map(nss, func(ns *Namespace, _ int) string {
		return ns.UUID()
	})
	if !lo.Contains(uuids, "") {
		return UUIDIn(uuids...)
	}
	exprs := make([]expression.Expression, 0, len(nss))
	for _, ns := range nss {
		exprs = append(exprs, Equal(ns))
	}
	return expression.Or(exprs...)
}

// NotIn returns an Expression that will match things that do not have
// any of a set of specified Namespace.
func NotIn(nss ...*Namespace) expression.Expression {
	uuids := lo.Map(nss, func(ns *Namespace, _ int) string {
		return ns.UUID()
	})
	if !lo.Contains(uuids, "") {
		return UUIDNotIn(uuids...)
	}
	exprs := make([]expression.Expression, 0, len(nss))
	for _, ns := range nss {
		exprs = append(exprs, NotEqual(ns))
	}
	return expression.And(exprs...)
}
