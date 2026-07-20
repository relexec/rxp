package kind

import (
	"github.com/samber/lo"

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
	v := p.Value
	switch v := v.(type) {
	case []api.KindName:
		for _, kn := range v {
			if err := kn.Validate(); err != nil {
				return errors.PredicateInvalid(err.Error())
			}
		}
	case api.KindName:
		return v.Validate()
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
	return nil
}

// NameEqual returns an Expression that will match Objects of a particular
// KindName.
func NameEqual(name api.KindName) query.Expression {
	return query.UnaryExpression{
		Predicate: NamePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match Objects not of a
// particular KindName.
func NameNotEqual(name api.KindName) query.Expression {
	return query.UnaryExpression{
		Predicate: NamePredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   name,
			},
		},
	}
}

// KindNameIn returns an Expression that will match Objects that are any of a
// set of specified KindNames.
func KindNameIn(names ...api.KindName) query.Expression {
	return query.UnaryExpression{
		Predicate: NamePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match Objects that are not any
// of a set of specified KindNames.
func NameNotIn(names ...api.KindName) query.Expression {
	return query.UnaryExpression{
		Predicate: NamePredicate{
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
		return errors.PredicateUnsupportedValueType(v)
	}
}

// UUIDEqual returns an Expression that will match things having a
// particular UUID.
func UUIDEqual(uuid string) query.Expression {
	return query.UnaryExpression{
		Predicate: UUIDPredicate{
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
		Predicate: UUIDPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match things that have any of a
// set of specified UUIDs.
func UUIDIn(uuids ...string) query.Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return UUIDEqual(uuids[0])
	}
	return query.UnaryExpression{
		Predicate: UUIDPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified UUIDs.
func UUIDNotIn(uuids ...string) query.Expression {
	return query.UnaryExpression{
		Predicate: UUIDPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorIn,
				Negated: true,
				Value:   uuids,
			},
		},
	}
}

type KindPredicate struct {
	query.BasePredicate
}

// Equal returns an Expression that will match things having a particular
// Kind.
func Equal(k *api.Kind) query.Expression {
	if k.UUID() != "" {
		return UUIDEqual(k.UUID())
	}
	if k.System() == nil {
		return NameEqual(k.Name())
	}
	return query.UnaryExpression{
		Predicate: KindPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: k,
			},
		},
	}
}

// NotEqual returns an Expression that will match things not having a
// particular Kind.
func NotEqual(k *api.Kind) query.Expression {
	if k.UUID() != "" {
		return UUIDNotEqual(k.UUID())
	}
	if k.System() == nil {
		return NameNotEqual(k.Name())
	}
	return query.UnaryExpression{
		Predicate: KindPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   k,
			},
		},
	}
}

// In returns an Expression that will match things that have any of a set
// of specified Kind.
func In(kinds ...*api.Kind) query.Expression {
	uuids := lo.Map(kinds, func(k *api.Kind, _ int) string {
		return k.UUID()
	})
	if !lo.Contains(uuids, "") {
		return UUIDIn(uuids...)
	}
	exprs := make([]query.Expression, 0, len(kinds))
	for _, k := range kinds {
		exprs = append(exprs, Equal(k))
	}
	return query.Or(exprs...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified Kind.
func NotIn(kinds ...*api.Kind) query.Expression {
	uuids := lo.Map(kinds, func(k *api.Kind, _ int) string {
		return k.UUID()
	})
	if !lo.Contains(uuids, "") {
		return UUIDNotIn(uuids...)
	}
	exprs := make([]query.Expression, 0, len(kinds))
	for _, k := range kinds {
		exprs = append(exprs, NotEqual(k))
	}
	return query.And(exprs...)
}
