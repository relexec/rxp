package errors

var (
	ErrKindInvalid               = New("invalid kind", WithCode(ErrCodeBadRequest))
	ErrKindEmpty                 = New("kind cannot be empty", WithWrap(ErrKindInvalid))
	ErrKindInvalidCharacters     = New("invalid characters", WithWrap(ErrKindInvalid))
	ErrKindRepeatedPeriods       = New("repeated periods", WithWrap(ErrKindInvalid))
	ErrKindInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrKindInvalid))

	ErrKindVersionUnknown = New("unknown kind version", WithCode(ErrCodeBadRequest))
)
