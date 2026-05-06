package errors

var (
	ErrDomainInvalid               = New("invalid domain", WithCode(ErrCodeBadRequest))
	ErrDomainInvalidCharacters     = New("invalid characters", WithWrap(ErrDomainInvalid))
	ErrDomainMaxLengthExceeded     = New("max length exceeded", WithWrap(ErrDomainInvalid))
	ErrDomainRepeatedPeriods       = New("repeated periods", WithWrap(ErrDomainInvalid))
	ErrDomainInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrDomainInvalid))
)
