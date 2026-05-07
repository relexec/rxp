package types

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const NamespaceNamePattern = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`

// NamespaceName is a DNS-formatted (RFC 1035-compliant) name.
//
// Note that unlike RFC 1035, there is no 253 character size limit on Namespace
// string length.
type NamespaceName string

// Validate returns an error if the Namespace is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (n NamespaceName) Validate() error {
	if len(n) == 0 {
		return errors.ErrNamespaceNameEmpty
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

// Namespace describes a logical division within a Domain.
//
// Namespaces are typically used to segregate data by tenancy boundaries.
type Namespace interface {
	Validatable
	// DOmain returns the Namespace's Domain.
	Domain() Domain
	// Name returns the name of the Namespace.
	Name() NamespaceName
}
