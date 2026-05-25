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
	v := p.BasePredicate.Value()
	switch v := v.(type) {
	case []api.DomainName:
		for _, dn := range v {
			if err := dn.Validate(); err != nil {
				return errors.PredicateInvalid(err.Error())
			}
		}
	case api.DomainName:
		return v.Validate()
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
	return nil
}

// DomainNameEqual returns an Expression that will match things having a
// particular DomainName.
func DomainNameEqual(name api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: name,
			},
		},
	}
}

// DomainNameNotEqual returns an Expression that will match things not having a
// particular DomainName.
func DomainNameNotEqual(name api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   name,
			},
		},
	}
}

// DomainNameIn returns an Expression that will match things that have any of a
// set of specified DomainNames.
func DomainNameIn(names ...api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: names,
			},
		},
	}
}

// DomainNameNotIn returns an Expression that will match things that do not
// have any of a set of specified DomainNames.
func DomainNameNotIn(names ...api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   names,
			},
		},
	}
}
