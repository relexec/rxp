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
}

// BasePredicate is the base struct from which all specialized Predicates
// derive.
type BasePredicate struct {
	// Op contains the Predicate's Operator.
	Op PredicateOperator
	// negated indicates whether the Predicate's Operator should be negated.
	Negated bool
	// value contains the value that is compared to the target field.
	Value any
}

// Validate returns an error if the Predicate is not valid.
func (p BasePredicate) Validate() error {
	if p.Value == nil {
		return errors.ErrPredicateValueRequired
	}
	return nil
}

// ContainsPredicate returns true if the supplied [Expression] has a Predicate
// whose type is any of a specified list of types. If the supplied expression
// is an [expression.OrExpression] or [expression.AndExpression], this function
// recursively checks sub-expressions to ensure that a NamePredicate is present
// in all sub-expressions.
func ContainsPredicate(expr Expression, filter func(p Predicate) bool) bool {
	switch expr := expr.(type) {
	case UnaryExpression:
		pred := expr.Predicate
		return filter(pred)
	case OrExpression:
		exprs := expr.Expressions()
		for _, e := range exprs {
			if ContainsPredicate(e, filter) {
				// At least one of the OR'd expressions was a NamePredicate.
				return true
			}
		}
		return false
	case AndExpression:
		exprs := expr.Expressions()
		for _, e := range exprs {
			if ContainsPredicate(e, filter) {
				// At least one of the AND'd expressions was a NamePredicate.
				return true
			}
		}
		return false
	default:
		return false
	}
}
