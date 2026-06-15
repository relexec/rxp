package object

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/query/expression"
)

type NamePredicate struct {
	expression.BasePredicate
}

func (p NamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value
	switch v.(type) {
	case []string, string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// NameEqual returns an Expression that will match things having a
// particular Name.
func NameEqual(name string) expression.Expression {
	return expression.UnaryExpression{
		NamePredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match things not having a
// particular Name.
func NameNotEqual(name string) expression.Expression {
	return expression.UnaryExpression{
		NamePredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorEqual,
				Negated: true,
				Value:   name,
			},
		},
	}
}

// NameIn returns an Expression that will match things that have any of a
// set of specified Names.
func NameIn(names ...string) expression.Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NameEqual(names[0])
	}
	return expression.UnaryExpression{
		NamePredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match things that do not
// have any of a set of specified Names.
func NameNotIn(names ...string) expression.Expression {
	return expression.UnaryExpression{
		NamePredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorIn,
				Negated: true,
				Value:   names,
			},
		},
	}
}

type UUIDPredicate struct {
	expression.BasePredicate
}

func (p UUIDPredicate) Validate() error {
	return p.BasePredicate.Validate()
}

// UUIDEqual returns an Expression that will match things having a particular
// UUID.
func UUIDEqual(uuid string) expression.Expression {
	return expression.UnaryExpression{
		UUIDPredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// UUIDNotEqual returns an Expression that will match things not having a
// particular UUID.
func UUIDNotEqual(uuid string) expression.Expression {
	return expression.UnaryExpression{
		UUIDPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorEqual,
				Negated: true,
				Value:   uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match things that have any of a set
// of specified UUIDs.
func UUIDIn(uuids ...string) expression.Expression {
	return expression.UnaryExpression{
		UUIDPredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not have any
// of a set of specified UUIDs.
func UUIDNotIn(uuids ...string) expression.Expression {
	return expression.UnaryExpression{
		UUIDPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorIn,
				Negated: true,
				Value:   uuids,
			},
		},
	}
}

type GenerationPredicate struct {
	expression.BasePredicate
}

func (p GenerationPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value
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
func GenerationEqual(generation api.Generation) expression.Expression {
	return expression.UnaryExpression{
		GenerationPredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorEqual,
				Value: generation,
			},
		},
	}
}

// GenerationNotEqual returns an Expression that will match things not having a
// particular Generation.
func GenerationNotEqual(generation api.Generation) expression.Expression {
	return expression.UnaryExpression{
		GenerationPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorEqual,
				Negated: true,
				Value:   generation,
			},
		},
	}
}

// GenerationIn returns an Expression that will match things that have any of a
// set of specified Generations.
func GenerationIn(generations ...api.Generation) expression.Expression {
	// flatten IN to = when there's only one value...
	if len(generations) == 1 {
		return GenerationEqual(generations[0])
	}
	return expression.UnaryExpression{
		GenerationPredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorIn,
				Value: generations,
			},
		},
	}
}

// GenerationNotIn returns an Expression that will match things that do not
// have any of a set of specified Generations.
func GenerationNotIn(generations ...api.Generation) expression.Expression {
	return expression.UnaryExpression{
		GenerationPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorIn,
				Negated: true,
				Value:   generations,
			},
		},
	}
}
