package errors

var (
	ErrQueryExpressionRequired            = New("query expression required", WithWrap(ErrInvalidQueryRequest))
	ErrInvalidQueryExpression             = New("invalid query expression", WithWrap(ErrInvalidQueryRequest))
	ErrInvalidQueryExpressionKindRequired = New("at least one kind required", WithWrap(ErrInvalidQueryExpression))
)
