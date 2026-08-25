package apikind

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const NamePattern = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`

// Name is the name of a type of thing managed by `rxp`.
//
// A valid Name is a DNS-formatted (RFC 1035-compliant) name type of Meta,
// e.g.  `flow.temporal.io`.
//
// Conventionally, Names are specified as singular, not plural nouns. So,
// "flow", not "flows".
//
// Furthermore, Names are conventionally all lower-cased, with dots
// separating coarser-grained categories/groups. So, "flow.temporal.io", not
// "TemporalFlow".
//
// You can use only alphanumeric characters and hyphens in the Name,
// separated by periods. Furthermore, the first and last character of the
// Name must be a letter or number, not a hyphen or period.
//
// Note that unlike RFC 1035, there is no 253 character size limit on Name
// string length.
//
// A Name must be unique within the scope of the `rxp` system installation,
// however for any Name that is intended to be used across multiple `rxp`
// system installations, the Name should be globally-unique.
type Name string

// Validate returns an error if the Name is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (n Name) Validate() error {
	if len(n) == 0 {
		return errors.ErrKindNameEmpty
	}
	first := rune(n[0])
	if !unicode.IsLetter(first) && !unicode.IsNumber(first) {
		return errors.ErrKindNameInvalidFirstCharacter
	}
	hasNonValidChars := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return false
		}
		return true
	}
	if strings.ContainsFunc(string(n), hasNonValidChars) {
		return errors.ErrKindNameInvalidCharacters
	}
	if strings.Contains(string(n), "..") {
		return errors.ErrKindNameRepeatedPeriods
	}
	return nil
}
