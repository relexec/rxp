package expression

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/expression"
	"github.com/relexec/rxp/predicate"
	"github.com/relexec/rxp/types"
)

type KindPredicate struct {
	predicate.Predicate
}

func (p KindPredicate) Validate() error {
	err := p.Predicate.Validate()
	if err != nil {
		return err
	}
	vals := p.Predicate.Values()
	for _, v := range vals {
		k, ok := v.(types.KindName)
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

// KindEqual returns an Expression that will match Objects of a particular
// Kind.
func KindEqual(k types.KindName) types.Expression {
	return expression.Unary(
		KindPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithValues(k),
			),
		},
	)
}

// KindNotEqual returns an Expression that will match Objects not of a
// particular Kind.
func KindNotEqual(k types.KindName) types.Expression {
	return expression.Unary(
		KindPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithNegated(true),
				predicate.WithValues(k),
			),
		},
	)
}

// KindIn returns an Expression that will match Objects that are any of a
// set of specified Kinds.
func KindIn(kinds ...types.KindName) types.Expression {
	values := make([]any, 0, len(kinds))
	for _, k := range kinds {
		values = append(values, k)
	}
	return expression.Unary(
		KindPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithValues(values...),
			),
		},
	)
}

// KindNotIn returns an Expression that will match Objects that are not any
// of a set of specified Kinds.
func KindNotIn(kinds ...types.KindName) types.Expression {
	values := make([]any, 0, len(kinds))
	for _, k := range kinds {
		values = append(values, k)
	}
	return expression.Unary(
		KindPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithNegated(true),
				predicate.WithValues(values...),
			),
		},
	)
}
