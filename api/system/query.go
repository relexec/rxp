package apisystem

import (
	"github.com/samber/lo"

	apierrors "github.com/relexec/rxp/api/errors"
	apiquery "github.com/relexec/rxp/api/query"
)

type UUIDPredicate struct {
	apiquery.BasePredicate
}

func (p UUIDPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.Value
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return apierrors.PredicateUnsupportedValueType(v)
	}
}

// UUIDEqual returns an Expression that will match things having a
// particular UUID.
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

// UUIDIn returns an Expression that will match things that have any of a
// set of specified UUIDs.
func UUIDIn(uuids ...string) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return UUIDEqual(uuids[0])
	}
	return apiquery.UnaryExpression{
		Predicate: UUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified UUIDs.
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

// Equal returns an Expression that will match things having a particular
// System.
func Equal(s *System) apiquery.Expression {
	return UUIDEqual(s.UUID)
}

// NotEqual returns an Expression that will match things not having a
// particular System.
func NotEqual(s *System) apiquery.Expression {
	return UUIDNotEqual(s.UUID)
}

// In returns an Expression that will match things that have any of a set
// of specified System.
func In(ss ...*System) apiquery.Expression {
	uuids := lo.Map(ss, func(s *System, _ int) string {
		return s.UUID
	})
	return UUIDIn(uuids...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified System.
func NotIn(ss ...*System) apiquery.Expression {
	uuids := lo.Map(ss, func(s *System, _ int) string {
		return s.UUID
	})
	return UUIDNotIn(uuids...)
}
