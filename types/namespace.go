package types

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const NamespacePattern = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`

// Namespace describes a logical division within a Domain.
//
// Namespaces are typically used to segregate data by tenancy boundaries.
//
// A valid Namespace is a DNS-formatted (RFC 1035-compliant) name.
//
// Note that unlike RFC 1035, there is no 253 character size limit on Namespace
// string length.
type Namespace string

// Validate returns an error if the Namespace is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (n Namespace) Validate() error {
	if len(n) == 0 {
		return nil
	}
	first := rune(n[0])
	if !unicode.IsLetter(first) && !unicode.IsNumber(first) {
		return errors.ErrNamespaceInvalidFirstCharacter
	}
	hasNonValidChars := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return false
		}
		return true
	}
	if strings.ContainsFunc(string(n), hasNonValidChars) {
		return errors.ErrNamespaceInvalidCharacters
	}
	if strings.Contains(string(n), "..") {
		return errors.ErrNamespaceRepeatedPeriods
	}
	return nil
}
