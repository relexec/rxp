package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
)

type GenerationPredicate struct {
	BasePredicate
}

func (p GenerationPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
	switch v := v.(type) {
	case []api.Generation:
		return nil
	case api.Generation:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// GenerationEqual returns an Expression that will match things having a
// particular Generation.
func GenerationEqual(generation api.Generation) Expression {
	return UnaryExpression{
		GenerationPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: generation,
			},
		},
	}
}

// GenerationNotEqual returns an Expression that will match things not having a
// particular Generation.
func GenerationNotEqual(generation api.Generation) Expression {
	return UnaryExpression{
		GenerationPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   generation,
			},
		},
	}
}

// GenerationIn returns an Expression that will match things that have any of a
// set of specified Generations.
func GenerationIn(generations ...api.Generation) Expression {
	// flatten IN to = when there's only one value...
	if len(generations) == 1 {
		return GenerationEqual(generations[0])
	}
	return UnaryExpression{
		GenerationPredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: generations,
			},
		},
	}
}

// GenerationNotIn returns an Expression that will match things that do not
// have any of a set of specified Generations.
func GenerationNotIn(generations ...api.Generation) Expression {
	return UnaryExpression{
		GenerationPredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   generations,
			},
		},
	}
}
