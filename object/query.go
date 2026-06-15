package object

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/query"
)

type NamePredicate struct {
	query.BasePredicate
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
func NameEqual(name string) query.Expression {
	return query.UnaryExpression{
		NamePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match things not having a
// particular Name.
func NameNotEqual(name string) query.Expression {
	return query.UnaryExpression{
		NamePredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   name,
			},
		},
	}
}

// NameIn returns an Expression that will match things that have any of a
// set of specified Names.
func NameIn(names ...string) query.Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NameEqual(names[0])
	}
	return query.UnaryExpression{
		NamePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match things that do not
// have any of a set of specified Names.
func NameNotIn(names ...string) query.Expression {
	return query.UnaryExpression{
		NamePredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorIn,
				Negated: true,
				Value:   names,
			},
		},
	}
}

type UUIDPredicate struct {
	query.BasePredicate
}

func (p UUIDPredicate) Validate() error {
	return p.BasePredicate.Validate()
}

// UUIDEqual returns an Expression that will match things having a particular
// UUID.
func UUIDEqual(uuid string) query.Expression {
	return query.UnaryExpression{
		UUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// UUIDNotEqual returns an Expression that will match things not having a
// particular UUID.
func UUIDNotEqual(uuid string) query.Expression {
	return query.UnaryExpression{
		UUIDPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match things that have any of a set
// of specified UUIDs.
func UUIDIn(uuids ...string) query.Expression {
	return query.UnaryExpression{
		UUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not have any
// of a set of specified UUIDs.
func UUIDNotIn(uuids ...string) query.Expression {
	return query.UnaryExpression{
		UUIDPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorIn,
				Negated: true,
				Value:   uuids,
			},
		},
	}
}

type GenerationPredicate struct {
	query.BasePredicate
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
func GenerationEqual(generation api.Generation) query.Expression {
	return query.UnaryExpression{
		GenerationPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: generation,
			},
		},
	}
}

// GenerationNotEqual returns an Expression that will match things not having a
// particular Generation.
func GenerationNotEqual(generation api.Generation) query.Expression {
	return query.UnaryExpression{
		GenerationPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   generation,
			},
		},
	}
}

// GenerationIn returns an Expression that will match things that have any of a
// set of specified Generations.
func GenerationIn(generations ...api.Generation) query.Expression {
	// flatten IN to = when there's only one value...
	if len(generations) == 1 {
		return GenerationEqual(generations[0])
	}
	return query.UnaryExpression{
		GenerationPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: generations,
			},
		},
	}
}

// GenerationNotIn returns an Expression that will match things that do not
// have any of a set of specified Generations.
func GenerationNotIn(generations ...api.Generation) query.Expression {
	return query.UnaryExpression{
		GenerationPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorIn,
				Negated: true,
				Value:   generations,
			},
		},
	}
}
