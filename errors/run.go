package errors

var (
	ErrRunRequestInvalid        = New("invalid run request", WithCode(ErrCodeBadRequest))
	ErrRunRequestUUIDRequired   = New("uuid required", WithWrap(ErrRunRequestInvalid))
	ErrRunRequestTargetRequired = New("target required", WithWrap(ErrRunRequestInvalid))
	ErrRunRequestOnRequired     = New("on timestamp required", WithWrap(ErrRunRequestInvalid))
)
