package api

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const (
	DomainNameMaxLength = 253
	DomainNamePattern   = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`
)

// DomainName is a DNS-formatted (RFC 1035-compliant) name less than 254
// characters.
//
// A DomainName must be unique within the scope of the `rxp` system
// installation.
type DomainName string

// Validate returns an error if the Domain is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (n DomainName) Validate() error {
	if len(n) == 0 {
		return errors.ErrDomainNameRequired
	}
	if len(n) > DomainNameMaxLength {
		return errors.ErrDomainNameMaxLengthExceeded
	}
	first := rune(n[0])
	if !unicode.IsLetter(first) && !unicode.IsNumber(first) {
		return errors.ErrDomainNameInvalidFirstCharacter
	}
	hasNonValidChars := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return false
		}
		return true
	}
	if strings.ContainsFunc(string(n), hasNonValidChars) {
		return errors.ErrDomainNameInvalidCharacters
	}
	if strings.Contains(string(n), "..") {
		return errors.ErrDomainNameRepeatedPeriods
	}
	return nil
}
