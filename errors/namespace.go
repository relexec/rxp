package errors

var (
	ErrNamespaceInvalid                   = New("invalid namespace", WithCode(ErrCodeBadRequest))
	ErrNamespaceDomainRequired            = New("domain required", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceUUIDRequired              = New("uuid required", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceNameRequired              = New("name required", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceNameInvalid               = New("invalid namespace name", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceNameInvalidCharacters     = New("invalid characters", WithWrap(ErrNamespaceNameInvalid))
	ErrNamespaceNameRepeatedPeriods       = New("repeated periods", WithWrap(ErrNamespaceNameInvalid))
	ErrNamespaceNameInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrNamespaceNameInvalid))
)
