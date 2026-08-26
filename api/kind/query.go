package apikind

import (
	"github.com/samber/lo"

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
	switch v := v.(type) {
	case []Name:
		for _, kn := range v {
			if err := kn.Validate(); err != nil {
				return apierrors.PredicateInvalid(err.Error())
			}
		}
	case Name:
		return v.Validate()
	default:
		return apierrors.PredicateUnsupportedValueType(v)
	}
	return nil
}

// NameEqual returns an Expression that will match Objects of a particular
// KindName.
func NameEqual(name Name) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: NamePredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match Objects not of a
// particular KindName.
func NameNotEqual(name Name) apiquery.Expression {
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

// KindNameIn returns an Expression that will match Objects that are any of a
// set of specified KindNames.
func KindNameIn(names ...Name) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: NamePredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match Objects that are not any
// of a set of specified KindNames.
func NameNotIn(names ...Name) apiquery.Expression {
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

type KindPredicate struct {
	apiquery.BasePredicate
}

// Equal returns an Expression that will match things having a particular
// Kind.
func Equal(k *Kind) apiquery.Expression {
	if k.UUID != "" {
		return UUIDEqual(k.UUID)
	}
	if k.System == nil {
		return NameEqual(k.Name)
	}
	return apiquery.UnaryExpression{
		Predicate: KindPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: k,
			},
		},
	}
}

// NotEqual returns an Expression that will match things not having a
// particular Kind.
func NotEqual(k *Kind) apiquery.Expression {
	if k.UUID != "" {
		return UUIDNotEqual(k.UUID)
	}
	if k.System == nil {
		return NameNotEqual(k.Name)
	}
	return apiquery.UnaryExpression{
		Predicate: KindPredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorEqual,
				Negated: true,
				Value:   k,
			},
		},
	}
}

// In returns an Expression that will match things that have any of a set
// of specified Kind.
func In(kinds ...*Kind) apiquery.Expression {
	uuids := lo.Map(kinds, func(k *Kind, _ int) string {
		return k.UUID
	})
	if !lo.Contains(uuids, "") {
		return UUIDIn(uuids...)
	}
	exprs := make([]apiquery.Expression, 0, len(kinds))
	for _, k := range kinds {
		exprs = append(exprs, Equal(k))
	}
	return apiquery.Or(exprs...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified Kind.
func NotIn(kinds ...*Kind) apiquery.Expression {
	uuids := lo.Map(kinds, func(k *Kind, _ int) string {
		return k.UUID
	})
	if !lo.Contains(uuids, "") {
		return UUIDNotIn(uuids...)
	}
	exprs := make([]apiquery.Expression, 0, len(kinds))
	for _, k := range kinds {
		exprs = append(exprs, NotEqual(k))
	}
	return apiquery.And(exprs...)
}
