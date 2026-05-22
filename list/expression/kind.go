package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/predicate"
	"github.com/relexec/rxp/types"
)

type KindNamePredicate struct {
	predicate.Predicate
}

func (p KindNamePredicate) Validate() error {
	err := p.Predicate.Validate()
	if err != nil {
		return err
	}
	vals := p.Predicate.Values()
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
func KindNameEqual(k api.KindName) types.Expression {
	return UnaryExpression{
		KindNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithValues(k),
			),
		},
	}
}

// KindNameNotEqual returns an Expression that will match Objects not of a
// particular KindName.
func KindNameNotEqual(k api.KindName) types.Expression {
	return UnaryExpression{
		KindNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithNegated(true),
				predicate.WithValues(k),
			),
		},
	}
}

// KindNameIn returns an Expression that will match Objects that are any of a
// set of specified KindNames.
func KindNameIn(kinds ...api.KindName) types.Expression {
	values := make([]any, 0, len(kinds))
	for _, k := range kinds {
		values = append(values, k)
	}
	return UnaryExpression{
		KindNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithValues(values...),
			),
		},
	}
}

// KindNameNotIn returns an Expression that will match Objects that are not any
// of a set of specified KindNames.
func KindNameNotIn(kinds ...api.KindName) types.Expression {
	values := make([]any, 0, len(kinds))
	for _, k := range kinds {
		values = append(values, k)
	}
	return UnaryExpression{
		KindNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithNegated(true),
				predicate.WithValues(values...),
			),
		},
	}
}

// ContainsKindPredicate returns true if the supplied [types.Expression] has a
// KindNamePredicate. If the supplied expression is an
// [expression.OrExpression] or [expression.AndExpression], this function
// recursively checks sub-expressions to ensure that a KindNamePredicate is
// present in all sub-expressions.
func ContainsKindPredicate(expr types.Expression) bool {
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
