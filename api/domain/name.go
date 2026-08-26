package apidomain

import (
	"strings"
	"unicode"

	apierrors "github.com/relexec/rxp/api/errors"
)

const (
	NameMaxLength = 253
	NamePattern   = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`
)

// Name is a DNS-formatted (RFC 1035-compliant) name less than 254
// characters.
//
// A Name must be unique within the scope of the `rxp` system
// installation.
type Name string

// Validate returns an error if the Domain is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (n Name) Validate() error {
	if len(n) == 0 {
		return apierrors.ErrDomainNameRequired
	}
	if len(n) > NameMaxLength {
		return apierrors.ErrDomainNameMaxLengthExceeded
	}
	first := rune(n[0])
	if !unicode.IsLetter(first) && !unicode.IsNumber(first) {
		return apierrors.ErrDomainNameInvalidFirstCharacter
	}
	hasNonValidChars := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return false
		}
		return true
	}
	if strings.ContainsFunc(string(n), hasNonValidChars) {
		return apierrors.ErrDomainNameInvalidCharacters
	}
	if strings.Contains(string(n), "..") {
		return apierrors.ErrDomainNameRepeatedPeriods
	}
	return nil
}
