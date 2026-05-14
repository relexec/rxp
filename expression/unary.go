package expression

import "github.com/relexec/rxp/types"

// UnaryExpression is a filtering expression that contains a single Predicate.
type UnaryExpression struct {
	types.Predicate
}

// Predicates returns the contained single Predicate.
func (e UnaryExpression) Predicates() []types.Predicate {
	return []types.Predicate{e}
}

// Unary returns a new UnaryExpression that evaluates the single supplied
// Predicate.
func Unary(pred types.Predicate) UnaryExpression {
	return UnaryExpression{pred}
}

var _ types.Expression = (*UnaryExpression)(nil)
