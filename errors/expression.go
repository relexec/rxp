package errors

var (
	ErrListExpressionRequired            = New("expression required", WithWrap(ErrInvalidListRequest))
	ErrInvalidListExpression             = New("invalid list expression", WithWrap(ErrInvalidListRequest))
	ErrInvalidListExpressionKindRequired = New("at least one kind required", WithWrap(ErrInvalidListExpression))
)
