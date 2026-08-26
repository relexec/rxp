package query

// AndExpression is a filtering expression that evaluates multiple expressions
// using AND logic.
type AndExpression struct {
	expressions []Expression
}

// Unary returns true if the Expression can be reduced to a single Predicate.
func (u AndExpression) Unary() bool {
	return false
}

// Expressions returns the contained Expressions.
func (e AndExpression) Expressions() []Expression {
	return e.expressions
}

// And returns a new AndExpression that evaluates the supplied Predicates using
// AND logic.
func And(exprs ...Expression) AndExpression {
	return AndExpression{expressions: exprs}
}

var _ Expression = (*AndExpression)(nil)
