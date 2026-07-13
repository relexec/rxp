package event

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/query"
)

type SequencePredicate struct {
	query.BasePredicate
}

func (p SequencePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []int:
		return nil
	case int:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// SequenceEqual returns an Expression that will match things having a
// particular Sequence.
func SequenceEqual(sequence int) query.Expression {
	return query.UnaryExpression{
		Predicate: SequencePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: sequence,
			},
		},
	}
}

// SequenceNotEqual returns an Expression that will match things not having a
// particular Sequence.
func SequenceNotEqual(sequence int) query.Expression {
	return query.UnaryExpression{
		Predicate: SequencePredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   sequence,
			},
		},
	}
}

// SequenceIn returns an Expression that will match things that have any of a
// set of specified Sequences.
func SequenceIn(sequences ...int) query.Expression {
	// flatten IN to = when there's only one value...
	if len(sequences) == 1 {
		return SequenceEqual(sequences[0])
	}
	return query.UnaryExpression{
		Predicate: SequencePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: sequences,
			},
		},
	}
}

// SequenceNotIn returns an Expression that will match things that do not
// have any of a set of specified Sequences.
func SequenceNotIn(sequences ...int) query.Expression {
	return query.UnaryExpression{
		Predicate: SequencePredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorIn,
				Negated: true,
				Value:   sequences,
			},
		},
	}
}

type TypePredicate struct {
	query.BasePredicate
}

func (p TypePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []Type:
		return nil
	case Type:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// TypeEqual returns an Expression that will match things having a
// particular Type.
func TypeEqual(typ Type) query.Expression {
	return query.UnaryExpression{
		Predicate: TypePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: typ,
			},
		},
	}
}

// TypeNotEqual returns an Expression that will match things not having a
// particular Type.
func TypeNotEqual(typ Type) query.Expression {
	return query.UnaryExpression{
		Predicate: TypePredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   typ,
			},
		},
	}
}

// TypeIn returns an Expression that will match things that have any of a
// set of specified Types.
func TypeIn(typs ...Type) query.Expression {
	// flatten IN to = when there's only one value...
	if len(typs) == 1 {
		return TypeEqual(typs[0])
	}
	return query.UnaryExpression{
		Predicate: TypePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: typs,
			},
		},
	}
}

// TypeNotIn returns an Expression that will match things that do not
// have any of a set of specified Types.
func TypeNotIn(typs ...Type) query.Expression {
	return query.UnaryExpression{
		Predicate: TypePredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorIn,
				Negated: true,
				Value:   typs,
			},
		},
	}
}
