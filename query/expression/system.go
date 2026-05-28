package expression

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/system"
	"github.com/samber/lo"
)

type SystemUUIDPredicate struct {
	BasePredicate
}

func (p SystemUUIDPredicate) Validate() error {
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

// SystemUUIDEqual returns an Expression that will match things having a
// particular SystemUUID.
func SystemUUIDEqual(uuid string) Expression {
	return UnaryExpression{
		SystemUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: uuid,
			},
		},
	}
}

// SystemUUIDNotEqual returns an Expression that will match things not having a
// particular SystemUUID.
func SystemUUIDNotEqual(uuid string) Expression {
	return UnaryExpression{
		SystemUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   uuid,
			},
		},
	}
}

// SystemUUIDIn returns an Expression that will match things that have any of a
// set of specified SystemUUIDs.
func SystemUUIDIn(uuids ...string) Expression {
	// flatten IN to = when there's only one value...
	if len(uuids) == 1 {
		return SystemUUIDEqual(uuids[0])
	}
	return UnaryExpression{
		SystemUUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: uuids,
			},
		},
	}
}

// SystemUUIDNotIn returns an Expression that will match things that do not
// have any of a set of specified SystemUUIDs.
func SystemUUIDNotIn(uuids ...string) Expression {
	return UnaryExpression{
		SystemUUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   uuids,
			},
		},
	}
}

// SystemEqual returns an Expression that will match things having a particular
// System.
func SystemEqual(dom *system.System) Expression {
	return SystemUUIDEqual(dom.UUID())
}

// SystemNotEqual returns an Expression that will match things not having a
// particular System.
func SystemNotEqual(dom *system.System) Expression {
	return SystemUUIDNotEqual(dom.UUID())
}

// SystemIn returns an Expression that will match things that have any of a set
// of specified System.
func SystemIn(doms ...*system.System) Expression {
	uuids := lo.Map(doms, func(dom *system.System, _ int) string {
		return dom.UUID()
	})
	return SystemUUIDIn(uuids...)
}

// SystemNotIn returns an Expression that will match things that do not
// have any of a set of specified System.
func SystemNotIn(doms ...*system.System) Expression {
	uuids := lo.Map(doms, func(dom *system.System, _ int) string {
		return dom.UUID()
	})
	return SystemUUIDNotIn(uuids...)
}
