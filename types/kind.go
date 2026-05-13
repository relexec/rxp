package types

import (
	"strings"
	"unicode"

	"github.com/relexec/rxp/errors"
)

const KindNamePattern = `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`

// KindName is the name of a type of thing managed by `rxp`.
//
// A valid KindName is a DNS-formatted (RFC 1035-compliant) name type of Meta,
// e.g.  `flow.temporal.io`.
//
// Conventionally, KindNames are specified as singular, not plural nouns. So,
// "flow", not "flows".
//
// Furthermore, KindNames are conventionally all lower-cased, with dots
// separating coarser-grained categories/groups. So, "flow.temporal.io", not
// "TemporalFlow".
//
// You can use only alphanumeric characters and hyphens in the KindName,
// separated by periods. Furthermore, the first and last character of the
// KindName must be a letter or number, not a hyphen or period.
//
// Note that unlike RFC 1035, there is no 253 character size limit on KindName
// string length.
//
// A KindName must be unique within the scope of the `rxp` system installation,
// however for any KindName that is intended to be used across multiple `rxp`
// system installations, the KindName should be globally-unique.
type KindName string

// Validate returns an error if the KindName is invalid.
//
// Note that we do not use regexp parsing here for performance reasons.
func (k KindName) Validate() error {
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

// Kind represents a type of thing managed by `rxp`.
//
// Kind has a Name which is unique within the scope of the `rxp` system installation.
//
// Kind has a Namescope which defines the scope in which Names of Objects with
// the Kind are unique.
type Kind interface {
	Validatable
	Differ
	// System returns the system identifier associated with the Kind.
	System() System
	// Name returns the name of the Kind.
	Name() KindName
	// Namescope returns the Namescope defining uniqueness constraints for
	// Names of Objects having this Kind.
	Namescope() Namescope
}
