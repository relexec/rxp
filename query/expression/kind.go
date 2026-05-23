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
	vals := p.BasePredicate.Values()
	for _, v := range vals {
		k, ok := v.(api.KindName)
		if !ok {
			return errors.PredicateUnsupportedValueType(v)
		}
		err = k.Validate()
		if err != nil {
			return errors.PredicateInvalid(err.Error())
		}
	}
	return nil
}

// KindNameEqual returns an Expression that will match Objects of a particular
// KindName.
func KindNameEqual(k api.KindName) Expression {
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:     PredicateOperatorEqual,
				values: []any{k},
			},
		},
	}
}

// KindNameNotEqual returns an Expression that will match Objects not of a
// particular KindName.
func KindNameNotEqual(k api.KindName) Expression {
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				values:  []any{k},
			},
		},
	}
}

// KindNameIn returns an Expression that will match Objects that are any of a
// set of specified KindNames.
func KindNameIn(kinds ...api.KindName) Expression {
	values := make([]any, 0, len(kinds))
	for _, k := range kinds {
		values = append(values, k)
	}
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:     PredicateOperatorIn,
				values: values,
			},
		},
	}
}

// KindNameNotIn returns an Expression that will match Objects that are not any
// of a set of specified KindNames.
func KindNameNotIn(kinds ...api.KindName) Expression {
	values := make([]any, 0, len(kinds))
	for _, k := range kinds {
		values = append(values, k)
	}
	return UnaryExpression{
		KindNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				values:  values,
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
