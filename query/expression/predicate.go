package expression

import (
	"github.com/relexec/rxp/errors"
)

// PredicateOperator is the type of comparison operator used by a Predicate.
type PredicateOperator int

const (
	PredicateOperatorEqual PredicateOperator = iota
	PredicateOperatorIn
)

// Predicate describes a single boolean expression. Predicates are used by
// `rxp` backend implementations to translate the user's matching/filtering
// intent into backend-specific filtering expressions (e.g. SQL statements).
type Predicate interface {
	// Validate returns an error if the Predicate is not valid.
	Validate() error
	// Operator returns the Predicate's Operator.
	Operator() PredicateOperator
	// Negated returns true if the Predicate's Operator should be negated. For
	// example, if Operator() returns PredicateOperatorEqual and Negated()
	// returns true, the Predicate will evaluate to true if the target field is
	// NOT equal to the Values.
	Negated() bool
	// Values returns the values that are compared to the target field.
	Values() []any
}

// BasePredicate is the base struct from which all specialized Predicates
// derive.
type BasePredicate struct {
	// op contains the Predicate's Operator.
	op PredicateOperator
	// negated indicates whether the Predicate's Operator should be negated.
	negated bool
	// values contains the values that are compared to the target field.
	values []any
}

// Operator returns the Predicate's Operator.
func (p BasePredicate) Operator() PredicateOperator {
	return p.op
}

// Negated returns true if the Predicate's Operator should be negated. For
// example, if Operator() returns OperatorEqual and Negated() returns true,
// the Predicate will evaluate to true if the target field is NOT equal to
// the Values.
func (p BasePredicate) Negated() bool {
	return p.negated
}

// Values returns the values that are compared to the target field.
func (p BasePredicate) Values() []any {
	return p.values
}

// Validate returns an error if the Predicate is not valid.
func (p BasePredicate) Validate() error {
	if len(p.values) == 0 {
		return errors.ErrPredicateValueRequired
	}
	return nil
}
