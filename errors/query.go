package errors

import "fmt"

var (
	ErrQueryExpressionRequired            = New("query expression required", WithWrap(ErrInvalidQueryRequest))
	ErrInvalidQueryExpression             = New("invalid query expression", WithWrap(ErrInvalidQueryRequest))
	ErrInvalidQueryExpressionKindRequired = New("at least one kind required", WithWrap(ErrInvalidQueryExpression))
	ErrInvalidQueryKindPredicate          = New("kind predicate not allowed in object query expression", WithWrap(ErrInvalidQueryExpression))
)

// UnsupportedExpression returns a wrapped ErrInvalidQueryExpression indicating
// the supplied expression type is not supported.
func UnsupportedExpression(expr any) error {
	return New(
		fmt.Sprintf("unsupported expression %T", expr),
		WithWrap(ErrInvalidQueryExpression),
	)
}

// UnsupportedPredicate returns a wrapped ErrInvalidQueryExpression indicating
// the supplied predicate type is not supported.
func UnsupportedPredicate(p any) error {
	return New(
		fmt.Sprintf("unsupported predicate %T", p),
		WithWrap(ErrInvalidQueryExpression),
	)
}

// UnsupportedPredicateOperator returns a wrapped ErrInvalidQueryExpression
// indicating the supplied predicate operator is not supported.
func UnsupportedPredicateOperator(op any) error {
	return New(
		fmt.Sprintf("unsupported predicate operator %T", op),
		WithWrap(ErrInvalidQueryExpression),
	)
}
