package expression

import (
	"github.com/samber/lo"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/errors"
)

type DomainNamePredicate struct {
	BasePredicate
}

func (p DomainNamePredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
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

// DomainNameEqual returns an Expression that will match things having a
// particular DomainName.
func DomainNameEqual(name api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: name,
			},
		},
	}
}

// DomainNameNotEqual returns an Expression that will match things not having a
// particular DomainName.
func DomainNameNotEqual(name api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   name,
			},
		},
	}
}

// DomainNameIn returns an Expression that will match things that have any of a
// set of specified DomainNames.
func DomainNameIn(names ...api.DomainName) Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return DomainNameEqual(names[0])
	}
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: names,
			},
		},
	}
}

// DomainNameNotIn returns an Expression that will match things that do not
// have any of a set of specified DomainNames.
func DomainNameNotIn(names ...api.DomainName) Expression {
	return UnaryExpression{
		DomainNamePredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   names,
			},
		},
	}
}

type DomainUUIDPredicate struct {
	BasePredicate
}

func (p DomainUUIDPredicate) Validate() error {
	err := p.BasePredicate.Validate()
	if err != nil {
		return err
	}
	v := p.BasePredicate.Value()
	switch v := v.(type) {
	case []string:
		return nil
	case string:
		return nil
	default:
		return errors.PredicateUnsupportedValueType(v)
	}
}

// DomainUUIDEqual returns an Expression that will match things having a
// particular DomainUUID.
func DomainUUIDEqual(uuid string) Expression {
	return UnaryExpression{
		DomainUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: uuid,
			},
		},
	}
}

// DomainUUIDNotEqual returns an Expression that will match things not having a
// particular DomainUUID.
func DomainUUIDNotEqual(uuid string) Expression {
	return UnaryExpression{
		DomainUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   uuid,
			},
		},
	}
}

// DomainUUIDIn returns an Expression that will match things that have any of a
// set of specified DomainUUIDs.
func DomainUUIDIn(uuids ...string) Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return DomainUUIDEqual(uuids[0])
	}
	return UnaryExpression{
		DomainUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: uuids,
			},
		},
	}
}

// DomainUUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified DomainUUIDs.
func DomainUUIDNotIn(uuids ...string) Expression {
	return UnaryExpression{
		DomainUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   uuids,
			},
		},
	}
}

type DomainPredicate struct {
	BasePredicate
}

// DomainEqual returns an Expression that will match things having a particular
// Domain.
func DomainEqual(dom *domain.Domain) Expression {
	if dom.UUID() != "" {
		return DomainUUIDEqual(dom.UUID())
	}
	if dom.System() == nil {
		return DomainNameEqual(dom.Name())
	}
	return UnaryExpression{
		DomainPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: dom,
			},
		},
	}
}

// DomainNotEqual returns an Expression that will match things not having a
// particular Domain.
func DomainNotEqual(dom *domain.Domain) Expression {
	if dom.UUID() != "" {
		return DomainUUIDNotEqual(dom.UUID())
	}
	if dom.System() == nil {
		return DomainNameNotEqual(dom.Name())
	}
	return UnaryExpression{
		DomainPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   dom,
			},
		},
	}
}

// DomainIn returns an Expression that will match things that have any of a set
// of specified Domain.
func DomainIn(doms ...*domain.Domain) Expression {
	uuids := lo.Map(doms, func(dom *domain.Domain, _ int) string {
		return dom.UUID()
	})
	if !lo.Contains(uuids, "") {
		return DomainUUIDIn(uuids...)
	}
	exprs := make([]Expression, 0, len(doms))
	for _, dom := range doms {
		exprs = append(exprs, DomainEqual(dom))
	}
	return Or(exprs...)
}

// DomainNotIn returns an Expression that will match things that do not
// have any of a set of specified Domain.
func DomainNotIn(doms ...*domain.Domain) Expression {
	uuids := lo.Map(doms, func(dom *domain.Domain, _ int) string {
		return dom.UUID()
	})
	if !lo.Contains(uuids, "") {
		return DomainUUIDNotIn(uuids...)
	}
	exprs := make([]Expression, 0, len(doms))
	for _, dom := range doms {
		exprs = append(exprs, DomainNotEqual(dom))
	}
	return And(exprs...)
}
