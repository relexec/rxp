package errors

var (
	ErrNamespaceInvalid               = New("invalid domain", WithCode(ErrCodeBadRequest))
	ErrNamespaceInvalidCharacters     = New("invalid characters", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceRepeatedPeriods       = New("repeated periods", WithWrap(ErrNamespaceInvalid))
	ErrNamespaceInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrNamespaceInvalid))
)
