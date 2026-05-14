package predicate

import (
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

// Predicate is the base struct from which all specialized Predicates
// derive.
type Predicate struct {
	// op contains the Predicate's Operator.
	op types.PredicateOperator
	// negated indicates whether the Predicate's Operator should be negated.
	negated bool
	// values contains the values that are compared to the target field.
	values []any
}

// Operator returns the Predicate's Operator.
func (p Predicate) Operator() types.PredicateOperator {
	return p.op
}

// Negated returns true if the Predicate's Operator should be negated. For
// example, if Operator() returns OperatorEqual and Negated() returns true,
// the Predicate will evaluate to true if the target field is NOT equal to
// the Values.
func (p Predicate) Negated() bool {
	return p.negated
}

// Values returns the values that are compared to the target field.
func (p Predicate) Values() []any {
	return p.values
}

// Validate returns an error if the Predicate is not valid.
func (p Predicate) Validate() error {
	if len(p.values) == 0 {
		return errors.ErrPredicateValueRequired
	}
	return nil
}
