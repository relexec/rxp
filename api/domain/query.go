package apidomain

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

// NameEqual returns an Expression that will match things having a
// particular Name.
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

// NameNotEqual returns an Expression that will match things not having a
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

// NameIn returns an Expression that will match things that have any of a
// set of specified Names.
func NameIn(names ...Name) apiquery.Expression {
	// flatten IN to = when there's only one value...
	if len(names) == 1 {
		return NameEqual(names[0])
	}
	return apiquery.UnaryExpression{
		Predicate: NamePredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorIn,
				Value: names,
			},
		},
	}
}

// NameNotIn returns an Expression that will match things that do not
// have any of a set of specified Names.
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

type DomainPredicate struct {
	apiquery.BasePredicate
}

// Equal returns an Expression that will match things having a particular
// Domain.
func Equal(dom *Domain) apiquery.Expression {
	if dom.UUID != "" {
		return UUIDEqual(dom.UUID)
	}
	if dom.System == nil {
		return NameEqual(dom.Name)
	}
	return apiquery.UnaryExpression{
		Predicate: DomainPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: dom,
			},
		},
	}
}

// NotEqual returns an Expression that will match things not having a
// particular Domain.
func NotEqual(dom *Domain) apiquery.Expression {
	if dom.UUID != "" {
		return UUIDNotEqual(dom.UUID)
	}
	if dom.System == nil {
		return NameNotEqual(dom.Name)
	}
	return apiquery.UnaryExpression{
		Predicate: DomainPredicate{
			apiquery.BasePredicate{
				Op:      apiquery.PredicateOperatorEqual,
				Negated: true,
				Value:   dom,
			},
		},
	}
}

// In returns an Expression that will match things that have any of a set
// of specified Domain.
func In(doms ...*Domain) apiquery.Expression {
	uuids := lo.Map(doms, func(dom *Domain, _ int) string {
		return dom.UUID
	})
	if !lo.Contains(uuids, "") {
		return UUIDIn(uuids...)
	}
	exprs := make([]apiquery.Expression, 0, len(doms))
	for _, dom := range doms {
		exprs = append(exprs, Equal(dom))
	}
	return apiquery.Or(exprs...)
}

// NotIn returns an Expression that will match things that do not
// have any of a set of specified Domain.
func NotIn(doms ...*Domain) apiquery.Expression {
	uuids := lo.Map(doms, func(dom *Domain, _ int) string {
		return dom.UUID
	})
	if !lo.Contains(uuids, "") {
		return UUIDNotIn(uuids...)
	}
	exprs := make([]apiquery.Expression, 0, len(doms))
	for _, dom := range doms {
		exprs = append(exprs, NotEqual(dom))
	}
	return apiquery.And(exprs...)
}

type RootNamePredicate struct {
	apiquery.BasePredicate
}

func (p RootNamePredicate) Validate() error {
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

// RootNameEqual returns an Expression that will match domains having a
// particular Name as their root Domain.
func RootNameEqual(name Name) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: RootNamePredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

type RootUUIDPredicate struct {
	apiquery.BasePredicate
}

func (p RootUUIDPredicate) Validate() error {
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

// RootUUIDEqual returns an Expression that will match domains having a root
// domain with a particular UUID.
func RootUUIDEqual(uuid string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: RootUUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

type RootDomainPredicate struct {
	apiquery.BasePredicate
}

// RootEqual returns an Expression that will match domains having a particular
// root Domain.
func RootEqual(dom *Domain) apiquery.Expression {
	if dom.UUID != "" {
		return RootUUIDEqual(dom.UUID)
	}
	if dom.System == nil {
		return RootNameEqual(dom.Name)
	}
	return apiquery.UnaryExpression{
		Predicate: RootDomainPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: dom,
			},
		},
	}
}

type ParentNamePredicate struct {
	apiquery.BasePredicate
}

func (p ParentNamePredicate) Validate() error {
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

// ParentNameEqual returns an Expression that will match domains having a
// particular Name as their parent Domain and any of that Domain's child
// domains.
func ParentNameEqual(name Name) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: ParentNamePredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: name,
			},
		},
	}
}

type ParentUUIDPredicate struct {
	apiquery.BasePredicate
}

func (p ParentUUIDPredicate) Validate() error {
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

// ParentUUIDEqual returns an Expression that will match domains having a
// parent domain with a particular UUID or any of that Domain's child domains.
func ParentUUIDEqual(uuid string) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: ParentUUIDPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: uuid,
			},
		},
	}
}

type ParentDomainPredicate struct {
	apiquery.BasePredicate
}

// ParentEqual returns an Expression that will match domains that are in the
// supplied Domain and any of that Domain's child domains.
func ParentEqual(dom *Domain) apiquery.Expression {
	return apiquery.UnaryExpression{
		Predicate: ParentDomainPredicate{
			apiquery.BasePredicate{
				Op:    apiquery.PredicateOperatorEqual,
				Value: dom,
			},
		},
	}
}
