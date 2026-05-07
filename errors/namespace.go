package errors

var (
	ErrNamespaceInvalid                   = New("invalid namespace", WithCode(ErrCodeBadRequest))
	ErrNamespaceDomainEmpty               = New("domain cannot be empty", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceNameEmpty                 = New("namespace name cannot be empty", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceNameInvalid               = New("invalid namespace name", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceNameInvalidCharacters     = New("invalid characters", WithWrap(ErrNamespaceNameInvalid))
	ErrNamespaceNameRepeatedPeriods       = New("repeated periods", WithWrap(ErrNamespaceNameInvalid))
	ErrNamespaceNameInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrNamespaceNameInvalid))
)
