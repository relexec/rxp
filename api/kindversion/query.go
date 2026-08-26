package apikindversion

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
		for _, dn := range v {
			if err := dn.Validate(); err != nil {
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
// Name.
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
// particular Name.
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

// NameIn returns an Expression that will match Objects that are any of a
// set of specified Names.
func NameIn(names ...Name) apiquery.Expression {
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
// of a set of specified Names.
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

type KindVersionPredicate struct {
	apiquery.BasePredicate
}

// Equal returns an Expression that will match things having a
// particular KindVersion.
func Equal(kv *KindVersion) apiquery.Expression {
	return NameEqual(kv.Name())
}

// NotEqual returns an Expression that will match things not having
// a particular KindVersion.
func NotEqual(kv *KindVersion) apiquery.Expression {
	return NameNotEqual(kv.Name())
}

// In returns an Expression that will match things that have any of
// a set of specified KindVersion.
func In(kvs ...*KindVersion) apiquery.Expression {
	names := lo.Map(
		kvs,
		func(kv *KindVersion, _ int) Name {
			return kv.Name()
		},
	)
	return NameIn(names...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified KindVersion.
func NotIn(kvs ...*KindVersion) apiquery.Expression {
	names := lo.Map(
		kvs,
		func(kv *KindVersion, _ int) Name {
			return kv.Name()
		},
	)
	return NameNotIn(names...)
}
