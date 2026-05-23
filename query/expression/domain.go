package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
)

type DomainNamePredicate struct {
	BasePredicate
}

func (p DomainNamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	vals := p.BasePredicate.Values()
	for _, v := range vals {
		d, ok := v.(api.DomainName)
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
func DomainNameEqual(k api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:     PredicateOperatorEqual,
				values: []any{k},
			},
		},
	}
}

// DomainNameNotEqual returns an Expression that will match Objects not of a
// particular DomainName.
func DomainNameNotEqual(k api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				values:  []any{k},
			},
		},
	}
}

// DomainNameIn returns an Expression that will match Objects that are any of a
// set of specified DomainNames.
func DomainNameIn(domains ...api.DomainName) Expression {
	values := make([]any, 0, len(domains))
	for _, k := range domains {
		values = append(values, k)
	}
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:     PredicateOperatorIn,
				values: values,
			},
		},
	}
}

// DomainNameNotIn returns an Expression that will match Objects that are not any
// of a set of specified DomainNames.
func DomainNameNotIn(domains ...api.DomainName) Expression {
	values := make([]any, 0, len(domains))
	for _, k := range domains {
		values = append(values, k)
	}
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				values:  values,
			},
		},
	}
}
