package errors

var (
	ErrDomainInvalid                   = New("invalid domain", WithCode(ErrCodeBadRequest))
	ErrDomainNameEmpty                 = New("domain name cannot be empty", WithWrap(ErrDomainInvalid))
	ErrDomainNameInvalid               = New("invalid domain name", WithWrap(ErrDomainInvalid))
	ErrDomainNameInvalidCharacters     = New("invalid characters", WithWrap(ErrDomainNameInvalid))
	ErrDomainNameMaxLengthExceeded     = New("max length exceeded", WithWrap(ErrDomainNameInvalid))
	ErrDomainNameRepeatedPeriods       = New("repeated periods", WithWrap(ErrDomainNameInvalid))
	ErrDomainNameInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrDomainNameInvalid))
	ErrDomainUUIDRequired              = New("uuid required", WithWrap(ErrDomainInvalid))
)
