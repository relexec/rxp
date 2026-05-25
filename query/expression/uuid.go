package expression

type UUIDPredicate struct {
	BasePredicate
}

func (p UUIDPredicate) Validate() error {
	return p.BasePredicate.Validate()
}

// UUIDEqual returns an Expression that will match things having a particular
// UUID.
func UUIDEqual(uuid string) Expression {
	return UnaryExpression{
		UUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorEqual,
				value: uuid,
			},
		},
	}
}

// UUIDNotEqual returns an Expression that will match things not having a
// particular UUID.
func UUIDNotEqual(uuid string) Expression {
	return UnaryExpression{
		UUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorEqual,
				negated: true,
				value:   uuid,
			},
		},
	}
}

// UUIDIn returns an Expression that will match things that have any of a set
// of specified UUIDs.
func UUIDIn(uuids ...string) Expression {
	return UnaryExpression{
		UUIDPredicate{
			BasePredicate{
				op:    PredicateOperatorIn,
				value: uuids,
			},
		},
	}
}

// UUIDNotIn returns an Expression that will match things that do not have any
// of a set of specified UUIDs.
func UUIDNotIn(uuids ...string) Expression {
	return UnaryExpression{
		UUIDPredicate{
			BasePredicate{
				op:      PredicateOperatorIn,
				negated: true,
				value:   uuids,
			},
		},
	}
}
