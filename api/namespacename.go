package api

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const NamespaceNamePattern = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`

// NamespaceName is a DNS-formatted (RFC 1035-compliant) name.
//
// Note that unlike RFC 1035, there is no 253 character size limit on
// NamespaceName string length.
type NamespaceName string

// Validate returns an error if the NamespaceName is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (n NamespaceName) Validate() error {
	if len(n) == 0 {
		return errors.ErrNamespaceNameRequired
	}
	first := rune(n[0])
	if !unicode.IsLetter(first) && !unicode.IsNumber(first) {
		return errors.ErrNamespaceNameInvalidFirstCharacter
	}
	hasNonValidChars := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return false
		}
		return true
	}
	if strings.ContainsFunc(string(n), hasNonValidChars) {
		return errors.ErrNamespaceNameInvalidCharacters
	}
	if strings.Contains(string(n), "..") {
		return errors.ErrNamespaceNameRepeatedPeriods
	}
	return nil
}
