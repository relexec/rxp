package types

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const (
	DomainMaxLength = 253
	DomainPattern   = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`
)

// Domain describes a top-level division or partition of things managed by rxp.
//
// A valid Domain is a DNS-formatted (RFC 1035-compliant) name less than 254
// characters.
//
// A Domain must be unique within the scope of the `rxp` system installation.
type Domain string

// Validate returns an error if the Domain is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (d Domain) Validate() error {
	if len(d) == 0 {
		return nil
	}
	if len(d) > DomainMaxLength {
		return errors.ErrDomainMaxLengthExceeded
	}
	first := rune(d[0])
	if !unicode.IsLetter(first) && !unicode.IsNumber(first) {
		return errors.ErrDomainInvalidFirstCharacter
	}
	hasNonValidChars := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return false
		}
		return true
	}
	if strings.ContainsFunc(string(d), hasNonValidChars) {
		return errors.ErrDomainInvalidCharacters
	}
	if strings.Contains(string(d), "..") {
		return errors.ErrDomainRepeatedPeriods
	}
	return nil
}
