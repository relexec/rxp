package system

import (
	"github.com/samber/lo"

	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/query/expression"
)

type UUIDPredicate struct {
	expression.BasePredicate
}

func (p UUIDPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// UUIDEqual returns an Expression that will match things having a
// particular UUID.
func UUIDEqual(uuid string) expression.Expression {
	return expression.UnaryExpression{
		Predicate: UUIDPredicate{
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
		Predicate: UUIDPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorEqual,
				Negated: true,
				Value:   uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match things that have any of a
// set of specified UUIDs.
func UUIDIn(uuids ...string) expression.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return UUIDEqual(uuids[0])
	}
	return expression.UnaryExpression{
		Predicate: UUIDPredicate{
			expression.BasePredicate{
				Op:    expression.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified UUIDs.
func UUIDNotIn(uuids ...string) expression.Expression {
	return expression.UnaryExpression{
		Predicate: UUIDPredicate{
			expression.BasePredicate{
				Op:      expression.PredicateOperatorIn,
				Negated: true,
				Value:   uuids,
			},
		},
	}
}

// Equal returns an Expression that will match things having a particular
// System.
func Equal(s *System) expression.Expression {
	return UUIDEqual(s.UUID())
}

// NotEqual returns an Expression that will match things not having a
// particular System.
func NotEqual(s *System) expression.Expression {
	return UUIDNotEqual(s.UUID())
}

// In returns an Expression that will match things that have any of a set
// of specified System.
func In(ss ...*System) expression.Expression {
	uuids := lo.Map(ss, func(s *System, _ int) string {
		return s.UUID()
	})
	return UUIDIn(uuids...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified System.
func NotIn(ss ...*System) expression.Expression {
	uuids := lo.Map(ss, func(s *System, _ int) string {
		return s.UUID()
	})
	return UUIDNotIn(uuids...)
}
