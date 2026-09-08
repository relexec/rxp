package errors

var (
	ErrCallerInvalid          = New("invalid caller", WithCode(ErrCodeBadRequest))
	ErrCallerIdentityRequired = New("caller identity required", WithWrap(ErrCallerInvalid))
)
