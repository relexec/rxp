package expression

import "github.com/relexec/rxp/types"

// AndExpression is a filtering expression that evaluates multiple Predicates
// using AND logic.
type AndExpression struct {
	predicates []types.Predicate
}

// Predicates returns the contained Predicates.
func (e AndExpression) Predicates() []types.Predicate {
	return e.predicates
}

// And returns a new AndExpression that evaluates the supplied Predicates using
// AND logic.
func And(preds ...types.Predicate) AndExpression {
	return AndExpression{predicates: preds}
}

var _ types.Expression = (*AndExpression)(nil)
