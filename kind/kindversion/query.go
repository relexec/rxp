package kindversion

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
	case []api.KindVersionName:
		for _, dn := range v {
			if err := dn.Validate(); err != nil {
				return errors.PredicateInvalid(err.Error())
			}
		}
	case api.KindVersionName:
		return v.Validate()
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
	return nil
}

// NameEqual returns an Expression that will match Objects of a particular
// KindVersionName.
func NameEqual(name api.KindVersionName) query.Expression {
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
// particular KindVersionName.
func NameNotEqual(name api.KindVersionName) query.Expression {
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

// NameIn returns an Expression that will match Objects that are any of a
// set of specified KindVersionNames.
func NameIn(names ...api.KindVersionName) query.Expression {
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
// of a set of specified KindVersionNames.
func NameNotIn(names ...api.KindVersionName) query.Expression {
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

type KindVersionPredicate struct {
	query.BasePredicate
}

// Equal returns an Expression that will match things having a
// particular KindVersion.
func Equal(kv *KindVersion) query.Expression {
	return NameEqual(kv.Name())
}

// NotEqual returns an Expression that will match things not having
// a particular KindVersion.
func NotEqual(kv *KindVersion) query.Expression {
	return NameNotEqual(kv.Name())
}

// In returns an Expression that will match things that have any of
// a set of specified KindVersion.
func In(kvs ...*KindVersion) query.Expression {
	names := lo.Map(
		kvs,
		func(kv *KindVersion, _ int) api.KindVersionName {
			return kv.Name()
		},
	)
	return NameIn(names...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified KindVersion.
func NotIn(kvs ...*KindVersion) query.Expression {
	names := lo.Map(
		kvs,
		func(kv *KindVersion, _ int) api.KindVersionName {
			return kv.Name()
		},
	)
	return NameNotIn(names...)
}
