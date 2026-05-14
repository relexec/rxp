package expression

import "github.com/relexec/rxp/types"

// OrExpression is a filtering expression that evaluates multiple Predicates
// using OR logic.
type OrExpression struct {
	predicates []types.Predicate
}

// Predicates returns the contained Predicates.
func (e OrExpression) Predicates() []types.Predicate {
	return e.predicates
}

// Or returns a new OrExpression that evaluates the supplied Predicates using
// OR logic.
func Or(preds ...types.Predicate) OrExpression {
	return OrExpression{predicates: preds}
}

var _ types.Expression = (*OrExpression)(nil)
