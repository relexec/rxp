package expression

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/samber/lo"
)

type KindVersionNamePredicate struct {
	BasePredicate
}

func (p KindVersionNamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
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

// KindVersionNameEqual returns an Expression that will match Objects of a particular
// KindVersionName.
func KindVersionNameEqual(name api.KindVersionName) Expression {
	return UnaryExpression{
		KindVersionNamePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: name,
			},
		},
	}
}

// KindVersionNameNotEqual returns an Expression that will match Objects not of a
// particular KindVersionName.
func KindVersionNameNotEqual(name api.KindVersionName) Expression {
	return UnaryExpression{
		KindVersionNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   name,
			},
		},
	}
}

// KindVersionNameIn returns an Expression that will match Objects that are any of a
// set of specified KindVersionNames.
func KindVersionNameIn(names ...api.KindVersionName) Expression {
	return UnaryExpression{
		KindVersionNamePredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: names,
			},
		},
	}
}

// KindVersionNameNotIn returns an Expression that will match Objects that are not any
// of a set of specified KindVersionNames.
func KindVersionNameNotIn(names ...api.KindVersionName) Expression {
	return UnaryExpression{
		KindVersionNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   names,
			},
		},
	}
}

type KindVersionPredicate struct {
	BasePredicate
}

// KindVersionEqual returns an Expression that will match things having a
// particular KindVersion.
func KindVersionEqual(kv *kindversion.KindVersion) Expression {
	return KindVersionNameEqual(kv.Name())
}

// KindVersionNotEqual returns an Expression that will match things not having
// a particular KindVersion.
func KindVersionNotEqual(kv *kindversion.KindVersion) Expression {
	return KindVersionNameNotEqual(kv.Name())
}

// KindVersionIn returns an Expression that will match things that have any of
// a set of specified KindVersion.
func KindVersionIn(kvs ...*kindversion.KindVersion) Expression {
	names := lo.Map(
		kvs,
		func(kv *kindversion.KindVersion, _ int) api.KindVersionName {
			return kv.Name()
		},
	)
	return KindVersionNameIn(names...)
}

// KindVersionNotIn returns an Expression that will match things that do not
// have any of a set of specified KindVersion.
func KindVersionNotIn(kvs ...*kindversion.KindVersion) Expression {
	names := lo.Map(
		kvs,
		func(kv *kindversion.KindVersion, _ int) api.KindVersionName {
			return kv.Name()
		},
	)
	return KindVersionNameNotIn(names...)
}
