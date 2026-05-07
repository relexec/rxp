package errors

var (
	ErrSystemInvalid   = New("invalid system", WithCode(ErrCodeBadRequest))
	ErrSystemUUIDEmpty = New("UUID cannot be empty", WithWrap(ErrSystemInvalid))
)
