package errors

var (
	ErrSystemInvalid      = New("invalid system", WithCode(ErrCodeBadRequest))
	ErrSystemUUIDRequired = New("uuid required", WithWrap(ErrSystemInvalid))
)
