package expression

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/expression"
	"github.com/relexec/rxp/predicate"
	"github.com/relexec/rxp/types"
)

type DomainPredicate struct {
	predicate.Predicate
}

func (p DomainPredicate) Validate() error {
	err := p.Predicate.Validate()
	if err != nil {
		return err
	}
	vals := p.Predicate.Values()
	for _, v := range vals {
		d, ok := v.(types.DomainName)
		if !ok {
			return errors.PredicateUnsupportedValueType(v)
		}
		err = d.Validate()
		if err != nil {
			return errors.PredicateInvalid(err.Error())
		}
	}
	return nil
}

// DomainEqual returns an Expression that will match Objects of a particular
// Domain.
func DomainEqual(k types.DomainName) types.Expression {
	return expression.Unary(
		DomainPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithValues(k),
			),
		},
	)
}

// DomainNotEqual returns an Expression that will match Objects not of a
// particular Domain.
func DomainNotEqual(k types.DomainName) types.Expression {
	return expression.Unary(
		DomainPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithNegated(true),
				predicate.WithValues(k),
			),
		},
	)
}

// DomainIn returns an Expression that will match Objects that are any of a
// set of specified Domains.
func DomainIn(domains ...types.DomainName) types.Expression {
	values := make([]any, 0, len(domains))
	for _, k := range domains {
		values = append(values, k)
	}
	return expression.Unary(
		DomainPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithValues(values...),
			),
		},
	)
}

// DomainNotIn returns an Expression that will match Objects that are not any
// of a set of specified Domains.
func DomainNotIn(domains ...types.DomainName) types.Expression {
	values := make([]any, 0, len(domains))
	for _, k := range domains {
		values = append(values, k)
	}
	return expression.Unary(
		DomainPredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithNegated(true),
				predicate.WithValues(values...),
			),
		},
	)
}
