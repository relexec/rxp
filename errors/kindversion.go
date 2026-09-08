package errors

var (
	ErrKindVersionInvalid         = New("invalid kindversion", WithCode(ErrCodeBadRequest))
	ErrKindVersionKindRequired    = New("kind required", WithWrap(ErrKindVersionInvalid))
	ErrKindVersionVersionRequired = New("version required", WithWrap(ErrKindVersionInvalid))
	ErrKindVersionSchemaRequired  = New("schema required", WithWrap(ErrKindVersionInvalid))
	ErrKindVersionUnknown         = New("unknown kind version", WithCode(ErrCodeBadRequest))
)
