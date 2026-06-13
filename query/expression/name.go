package expression

import (
	"github.com/relexec/rxp/errors"
)

type NamePredicate struct {
	BasePredicate
}

func (p NamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
	switch v.(type) {
	case []string, string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// NameEqual returns an Expression that will match things having a
// particular Name.
func NameEqual(name string) Expression {
	return UnaryExpression{
		NamePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match things not having a
// particular Name.
func NameNotEqual(name string) Expression {
	return UnaryExpression{
		NamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   name,
			},
		},
	}
}

// NameIn returns an Expression that will match things that have any of a
// set of specified Names.
func NameIn(names ...string) Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NameEqual(names[0])
	}
	return UnaryExpression{
		NamePredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match things that do not
// have any of a set of specified Names.
func NameNotIn(names ...string) Expression {
	return UnaryExpression{
		NamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   names,
			},
		},
	}
}
