package errors

var (
	ErrInvalidObject             = New("invalid object", WithCode(ErrCodeBadRequest))
	ErrNilObject                 = New("nil object parameter", WithWrap(ErrInvalidObject))
	ErrObjectKindVersionRequired = New("kindversion required", WithWrap(ErrInvalidObject))
	ErrObjectNameRequired        = New("name required", WithWrap(ErrInvalidObject))
	ErrObjectUUIDRequired        = New("uuid required", WithWrap(ErrInvalidObject))
	ErrObjectDomainRequired      = New("domain required", WithWrap(ErrInvalidObject))
)
