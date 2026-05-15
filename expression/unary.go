package expression

import "github.com/relexec/rxp/types"

// UnaryExpression is a filtering expression that contains a single Predicate.
type UnaryExpression struct {
	types.Predicate
}

// Unary returns true if the Expression can be reduced to a single Predicate.
func (u UnaryExpression) Unary() bool {
	return true
}

var _ types.Expression = (*UnaryExpression)(nil)
