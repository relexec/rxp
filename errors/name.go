package errors

var (
	ErrInvalidName      = New("invalid selector", WithWrap(ErrInvalidReadRequest))
	ErrNameNameRequired = New("name required", WithWrap(ErrInvalidName))
)
