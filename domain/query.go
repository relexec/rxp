package domain

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
	v := p.BasePredicate.Value
	switch v := v.(type) {
	case []api.DomainName:
		for _, dn := range v {
			if err := dn.Validate(); err != nil {
				return errors.PredicateInvalid(err.Error())
			}
		}
	case api.DomainName:
		return v.Validate()
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
	return nil
}

// NameEqual returns an Expression that will match things having a
// particular DomainName.
func NameEqual(name api.DomainName) query.Expression {
	return query.UnaryExpression{
		Predicate: NamePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

// NameNotEqual returns an Expression that will match things not having a
// particular DomainName.
func NameNotEqual(name api.DomainName) query.Expression {
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

// NameIn returns an Expression that will match things that have any of a
// set of specified DomainNames.
func NameIn(names ...api.DomainName) query.Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NameEqual(names[0])
	}
	return query.UnaryExpression{
		Predicate: NamePredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match things that do not
// have any of a set of specified DomainNames.
func NameNotIn(names ...api.DomainName) query.Expression {
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

type DomainPredicate struct {
	query.BasePredicate
}

// Equal returns an Expression that will match things having a particular
// Domain.
func Equal(dom *Domain) query.Expression {
	if dom.UUID() != "" {
		return UUIDEqual(dom.UUID())
	}
	if dom.System() == nil {
		return NameEqual(dom.Name())
	}
	return query.UnaryExpression{
		Predicate: DomainPredicate{
			query.BasePredicate{
				Op:    query.PredicateOperatorEqual,
				Value: dom,
			},
		},
	}
}

// NotEqual returns an Expression that will match things not having a
// particular Domain.
func NotEqual(dom *Domain) query.Expression {
	if dom.UUID() != "" {
		return UUIDNotEqual(dom.UUID())
	}
	if dom.System() == nil {
		return NameNotEqual(dom.Name())
	}
	return query.UnaryExpression{
		Predicate: DomainPredicate{
			query.BasePredicate{
				Op:      query.PredicateOperatorEqual,
				Negated: true,
				Value:   dom,
			},
		},
	}
}

// In returns an Expression that will match things that have any of a set
// of specified Domain.
func In(doms ...*Domain) query.Expression {
	uuids := lo.Map(doms, func(dom *Domain, _ int) string {
		return dom.UUID()
	})
	if !lo.Contains(uuids, "") {
		return UUIDIn(uuids...)
	}
	exprs := make([]query.Expression, 0, len(doms))
	for _, dom := range doms {
		exprs = append(exprs, Equal(dom))
	}
	return query.Or(exprs...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified Domain.
func NotIn(doms ...*Domain) query.Expression {
	uuids := lo.Map(doms, func(dom *Domain, _ int) string {
		return dom.UUID()
	})
	if !lo.Contains(uuids, "") {
		return UUIDNotIn(uuids...)
	}
	exprs := make([]query.Expression, 0, len(doms))
	for _, dom := range doms {
		exprs = append(exprs, NotEqual(dom))
	}
	return query.And(exprs...)
}
