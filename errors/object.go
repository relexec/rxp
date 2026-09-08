package errors

var (
	ErrObjectInvalid             = New("invalid object", WithCode(ErrCodeBadRequest))
	ErrObjectNil                 = New("nil object parameter", WithWrap(ErrObjectInvalid))
	ErrObjectKindVersionRequired = New("kindversion required", WithWrap(ErrObjectInvalid))
	ErrObjectNameRequired        = New("name required", WithWrap(ErrObjectInvalid))
	ErrObjectUUIDRequired        = New("uuid required", WithWrap(ErrObjectInvalid))
	ErrObjectDomainRequired      = New("domain required", WithWrap(ErrObjectInvalid))
)
