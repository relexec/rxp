package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
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
// recursively checks sub-expressions to ensure that a KindNamePredicate is
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
