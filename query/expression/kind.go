package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/kind"
	"github.com/samber/lo"
)

type KindNamePredicate struct {
	BasePredicate
}

func (p KindNamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
	switch v := v.(type) {
	case []api.KindName:
		for _, dn := range v {
			if err := dn.Validate(); err != nil {
				return errors.PredicateInvalid(err.Error())
			}
		}
	case api.KindName:
		return v.Validate()
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
	return nil
}

// KindNameEqual returns an Expression that will match Objects of a particular
// KindName.
func KindNameEqual(name api.KindName) Expression {
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: name,
			},
		},
	}
}

// KindNameNotEqual returns an Expression that will match Objects not of a
// particular KindName.
func KindNameNotEqual(name api.KindName) Expression {
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   name,
			},
		},
	}
}

// KindNameIn returns an Expression that will match Objects that are any of a
// set of specified KindNames.
func KindNameIn(names ...api.KindName) Expression {
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: names,
			},
		},
	}
}

// KindNameNotIn returns an Expression that will match Objects that are not any
// of a set of specified KindNames.
func KindNameNotIn(names ...api.KindName) Expression {
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   names,
			},
		},
	}
}

// ContainsKindPredicate returns true if the supplied [Expression] has a
// KindNamePredicate. If the supplied expression is an
// [expression.OrExpression] or [expression.AndExpression], this function
// recursively checkinds sub-expressions to ensure that a KindNamePredicate is
// present in all sub-expressions.
func ContainsKindPredicate(expr Expression) bool {
	switch expr := expr.(type) {
	case UnaryExpression:
		pred := expr.Predicate
		_, ok := pred.(KindNamePredicate)
		return ok
	case OrExpression:
		exprs := expr.Expressions()
		for _, e := range exprs {
			if ContainsKindPredicate(e) {
				// At least one of the OR'd expressions was a KindNamePredicate.
				return true
			}
		}
		return false
	case AndExpression:
		exprs := expr.Expressions()
		for _, e := range exprs {
			if ContainsKindPredicate(e) {
				// At least one of the AND'd expressions was a KindNamePredicate.
				return true
			}
		}
		return false
	default:
		return false
	}
}

type KindUUIDPredicate struct {
	BasePredicate
}

func (p KindUUIDPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// KindUUIDEqual returns an Expression that will match things having a
// particular KindUUID.
func KindUUIDEqual(uuid string) Expression {
	return UnaryExpression{
		KindUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: uuid,
			},
		},
	}
}

// KindUUIDNotEqual returns an Expression that will match things not having a
// particular KindUUID.
func KindUUIDNotEqual(uuid string) Expression {
	return UnaryExpression{
		KindUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   uuid,
			},
		},
	}
}

// KindUUIDIn returns an Expression that will match things that have any of a
// set of specified KindUUIDs.
func KindUUIDIn(uuids ...string) Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return KindUUIDEqual(uuids[0])
	}
	return UnaryExpression{
		KindUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: uuids,
			},
		},
	}
}

// KindUUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified KindUUIDs.
func KindUUIDNotIn(uuids ...string) Expression {
	return UnaryExpression{
		KindUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   uuids,
			},
		},
	}
}

type KindPredicate struct {
	BasePredicate
}

// KindEqual returns an Expression that will match things having a particular
// Kind.
func KindEqual(k *kind.Kind) Expression {
	if k.UUID() != "" {
		return KindUUIDEqual(k.UUID())
	}
	if k.System() == nil {
		return KindNameEqual(k.Name())
	}
	return UnaryExpression{
		KindPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: k,
			},
		},
	}
}

// KindNotEqual returns an Expression that will match things not having a
// particular Kind.
func KindNotEqual(k *kind.Kind) Expression {
	if k.UUID() != "" {
		return KindUUIDNotEqual(k.UUID())
	}
	if k.System() == nil {
		return KindNameNotEqual(k.Name())
	}
	return UnaryExpression{
		KindPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   k,
			},
		},
	}
}

// KindIn returns an Expression that will match things that have any of a set
// of specified Kind.
func KindIn(kinds ...*kind.Kind) Expression {
	uuids := lo.Map(kinds, func(k *kind.Kind, _ int) string {
		return k.UUID()
	})
	if !lo.Contains(uuids, "") {
		return KindUUIDIn(uuids...)
	}
	exprs := make([]Expression, 0, len(kinds))
	for _, k := range kinds {
		exprs = append(exprs, KindEqual(k))
	}
	return Or(exprs...)
}

// KindNotIn returns an Expression that will match things that do not
// have any of a set of specified Kind.
func KindNotIn(kinds ...*kind.Kind) Expression {
	uuids := lo.Map(kinds, func(k *kind.Kind, _ int) string {
		return k.UUID()
	})
	if !lo.Contains(uuids, "") {
		return KindUUIDNotIn(uuids...)
	}
	exprs := make([]Expression, 0, len(kinds))
	for _, k := range kinds {
		exprs = append(exprs, KindNotEqual(k))
	}
	return And(exprs...)
}
