package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
)

type NamespaceNamePredicate struct {
	BasePredicate
}

func (p NamespaceNamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
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

// NamespaceNameEqual returns an Expression that will match things having a
// particular NamespaceName.
func NamespaceNameEqual(name api.NamespaceName) Expression {
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: name,
			},
		},
	}
}

// NamespaceNameNotEqual returns an Expression that will match things not having a
// particular NamespaceName.
func NamespaceNameNotEqual(name api.NamespaceName) Expression {
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   name,
			},
		},
	}
}

// NamespaceNameIn returns an Expression that will match things that have any of a
// set of specified NamespaceNames.
func NamespaceNameIn(names ...api.NamespaceName) Expression {
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: names,
			},
		},
	}
}

// NamespaceNameNotIn returns an Expression that will match things that do not
// have any of a set of specified NamespaceNames.
func NamespaceNameNotIn(names ...api.NamespaceName) Expression {
	return UnaryExpression{
		NamespaceNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   names,
			},
		},
	}
}
