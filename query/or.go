package query

// OrExpression is a filtering expression that evaluates multiple expressions
// using OR logic.
type OrExpression struct {
	expressions []Expression
}

// Unary returns true if the Expression can be reduced to a single Predicate.
func (u OrExpression) Unary() bool {
	return false
}

// Expressions returns the contained Expressions.
func (e OrExpression) Expressions() []Expression {
	return e.expressions
}

// Or returns a new OrExpression that evaluates the supplied Predicates using
// AND logic.
func Or(exprs ...Expression) OrExpression {
	return OrExpression{expressions: exprs}
}

var _ Expression = (*OrExpression)(nil)
