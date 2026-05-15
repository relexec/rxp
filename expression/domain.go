package expression

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/predicate"
	"github.com/relexec/rxp/types"
)

type DomainNamePredicate struct {
	predicate.Predicate
}

func (p DomainNamePredicate) Validate() error {
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

// DomainNameEqual returns an Expression that will match Objects of a particular
// DomainName.
func DomainNameEqual(k types.DomainName) types.Expression {
	return UnaryExpression{
		DomainNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithValues(k),
			),
		},
	}
}

// DomainNameNotEqual returns an Expression that will match Objects not of a
// particular DomainName.
func DomainNameNotEqual(k types.DomainName) types.Expression {
	return UnaryExpression{
		DomainNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorEqual),
				predicate.WithNegated(true),
				predicate.WithValues(k),
			),
		},
	}
}

// DomainNameIn returns an Expression that will match Objects that are any of a
// set of specified DomainNames.
func DomainNameIn(domains ...types.DomainName) types.Expression {
	values := make([]any, 0, len(domains))
	for _, k := range domains {
		values = append(values, k)
	}
	return UnaryExpression{
		DomainNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithValues(values...),
			),
		},
	}
}

// DomainNameNotIn returns an Expression that will match Objects that are not any
// of a set of specified DomainNames.
func DomainNameNotIn(domains ...types.DomainName) types.Expression {
	values := make([]any, 0, len(domains))
	for _, k := range domains {
		values = append(values, k)
	}
	return UnaryExpression{
		DomainNamePredicate{
			predicate.New(
				predicate.WithOperator(types.PredicateOperatorIn),
				predicate.WithNegated(true),
				predicate.WithValues(values...),
			),
		},
	}
}
