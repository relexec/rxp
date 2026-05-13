package errors

var (
	ErrKindInvalid                   = New("invalid kind", WithCode(ErrCodeBadRequest))
	ErrKindNameInvalid               = New("invalid kind name", WithCode(ErrCodeBadRequest))
	ErrKindNameEmpty                 = New("kind name cannot be empty", WithWrap(ErrKindNameInvalid))
	ErrKindNameInvalidCharacters     = New("invalid characters", WithWrap(ErrKindNameInvalid))
	ErrKindNameRepeatedPeriods       = New("repeated periods", WithWrap(ErrKindNameInvalid))
	ErrKindNameInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrKindNameInvalid))

	ErrKindVersionUnknown = New("unknown kind version", WithCode(ErrCodeBadRequest))
)
