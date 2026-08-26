package apiobject

import (
	apicore "github.com/relexec/rxp/api/core"
	apierrors "github.com/relexec/rxp/api/errors"
	apiquery "github.com/relexec/rxp/api/query"
)

type NamePredicate struct {
	apiquery.BasePredicate
}

func (p NamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v.(type) {
	case []string, string:
		return nil
	default:
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// NameEqual returns an Expression that will match things having a
// particular Name.
func NameEqual(name string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: NamePredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match things not having a
// particular Name.
func NameNotEqual(name string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: NamePredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorEqual,
				Negated: true,
				Value:   name,
			},
		},
	}
}

// NameIn returns an Expression that will match things that have any of a
// set of specified Names.
func NameIn(names ...string) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NameEqual(names[0])
	}
	return apiquery.UnaryExpression{
		Predicate: NamePredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match things that do not
// have any of a set of specified Names.
func NameNotIn(names ...string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: NamePredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorIn,
				Negated: true,
				Value:   names,
			},
		},
	}
}

type UUIDPredicate struct {
	apiquery.BasePredicate
}

func (p UUIDPredicate) Validate() error {
	return p.BasePredicate.Validate()
}

// UUIDEqual returns an Expression that will match things having a particular
// UUID.
func UUIDEqual(uuid string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: UUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

// UUIDNotEqual returns an Expression that will match things not having a
// particular UUID.
func UUIDNotEqual(uuid string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: UUIDPredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorEqual,
				Negated: true,
				Value:   uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match things that have any of a set
// of specified UUIDs.
func UUIDIn(uuids ...string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: UUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not have any
// of a set of specified UUIDs.
func UUIDNotIn(uuids ...string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: UUIDPredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorIn,
				Negated: true,
				Value:   uuids,
			},
		},
	}
}

type GenerationPredicate struct {
	apiquery.BasePredicate
}

func (p GenerationPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []apicore.Generation:
		return nil
	case apicore.Generation:
		return nil
	default:
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// GenerationEqual returns an Expression that will match things having a
// particular Generation.
func GenerationEqual(generation apicore.Generation) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: GenerationPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: generation,
			},
		},
	}
}

// GenerationNotEqual returns an Expression that will match things not having a
// particular Generation.
func GenerationNotEqual(generation apicore.Generation) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: GenerationPredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorEqual,
				Negated: true,
				Value:   generation,
			},
		},
	}
}

// GenerationIn returns an Expression that will match things that have any of a
// set of specified Generations.
func GenerationIn(generations ...apicore.Generation) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(generations) == 1 {
		return GenerationEqual(generations[0])
	}
	return apiquery.UnaryExpression{
		Predicate: GenerationPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: generations,
			},
		},
	}
}

// GenerationNotIn returns an Expression that will match things that do not
// have any of a set of specified Generations.
func GenerationNotIn(generations ...apicore.Generation) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: GenerationPredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorIn,
				Negated: true,
				Value:   generations,
			},
		},
	}
}
