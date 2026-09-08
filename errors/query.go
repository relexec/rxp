package errors

import "fmt"

var (
	ErrQueryInvalid                    = New("invalid query", WithCode(ErrCodeBadRequest))
	ErrQueryExpressionRequired         = New("query expression required", WithWrap(ErrQueryInvalid))
	ErrQueryInvalidExpression          = New("invalid query expression", WithWrap(ErrQueryInvalid))
	ErrQueryKindPredicateInObjectQuery = New("kind predicate not allowed in object query expression", WithWrap(ErrQueryInvalidExpression))
)

// UnsupportedExpression returns a wrapped ErrQueryInvalidExpression indicating
// the supplied expression type is not supported.
func UnsupportedExpression(expr any) error {
	return New(
		fmt.Sprintf("unsupported expression %T", expr),
		WithWrap(ErrQueryInvalidExpression),
	)
}

// UnsupportedPredicate returns a wrapped ErrQueryInvalidExpression indicating
// the supplied predicate type is not supported.
func UnsupportedPredicate(p any) error {
	return New(
		fmt.Sprintf("unsupported predicate %T", p),
		WithWrap(ErrQueryInvalidExpression),
	)
}

// UnsupportedPredicateOperator returns a wrapped ErrQueryInvalidExpression
// indicating the supplied predicate operator is not supported.
func UnsupportedPredicateOperator(op any) error {
	return New(
		fmt.Sprintf("unsupported predicate operator %T", op),
		WithWrap(ErrQueryInvalidExpression),
	)
}
