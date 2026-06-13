package errors

var (
	ErrDomainInvalid                   = New("invalid domain", WithCode(ErrCodeBadRequest))
	ErrDomainParentSystemDifferent     = New("parent system must be same", WithWrap(ErrDomainInvalid))
	ErrDomainParentNotFound            = New("parent not found", WithWrap(ErrDomainInvalid))
	ErrDomainUUIDRequired              = New("uuid required", WithWrap(ErrDomainInvalid))
	ErrDomainNameRequired              = New("name required", WithWrap(ErrDomainInvalid))
	ErrDomainNameInvalid               = New("invalid domain name", WithWrap(ErrDomainInvalid))
	ErrDomainNameInvalidCharacters     = New("invalid characters", WithWrap(ErrDomainNameInvalid))
	ErrDomainNameMaxLengthExceeded     = New("max length exceeded", WithWrap(ErrDomainNameInvalid))
	ErrDomainNameRepeatedPeriods       = New("repeated periods", WithWrap(ErrDomainNameInvalid))
	ErrDomainNameInvalidFirstCharacter = New("first character must be letter or number", WithWrap(ErrDomainNameInvalid))
)
