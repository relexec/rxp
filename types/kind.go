package types

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const KindPattern = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`

// Kind is the name of a type of thing managed by rxp.
//
// A valid Kind is a DNS-formatted (RFC 1035-compliant) name type of Meta, e.g.
// `flow.temporal.io`.
//
// Conventionally, Kinds are specified as singular, not plural nouns. So,
// "flow", not "flows".
//
// Furthermore, Kinds are conventionally all lower-cased, with dots separating
// coarser-grained categories/groups. So, "flow.temporal.io", not
// "TemporalFlow".
//
// You can use only alphanumeric characters and hyphens in the Kind name parts,
// separated by periods. Furthermore, the first character of the Kind must be a
// letter or number, not a hyphen or period.
//
// Note that unlike RFC 1035, there is no 253 character size limit on Kind
// string length.
//
// A Kind must be unique within the scope of the `rxp` system installation,
// however for any Kind that is intended to be used across multiple `rxp`
// system installations, the Kind should be globally-unique.
type Kind string

// Validate returns an error if the Kind is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (k Kind) Validate() error {
	if len(k) == 0 {
		return errors.ErrKindEmpty
	}
	first := rune(k[0])
	if !unicode.IsLetter(first) && !unicode.IsNumber(first) {
		return errors.ErrKindInvalidFirstCharacter
	}
	hasNonValidChars := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '.' || r == '-' {
			return false
		}
		return true
	}
	if strings.ContainsFunc(string(k), hasNonValidChars) {
		return errors.ErrKindInvalidCharacters
	}
	if strings.Contains(string(k), "..") {
		return errors.ErrKindRepeatedPeriods
	}
	return nil
}
