package errors

import (
	"fmt"
)

var (
	ErrInvalidMeta = New("invalid meta", WithCode(ErrCodeBadRequest))
)

// MetaMissingKind returns a wrapped ErrMetaInvalid indicating the
// supplied Meta type is missing an identifier.
func MetaMissingKind() error {
	return New("missing kind", WithWrap(ErrInvalidMeta))
}

// MetaMissingVersion returns a wrapped ErrMetaInvalid indicating the supplied
// Meta is missing a version.
func MetaMissingVersion(typ any) error {
	return New(
		fmt.Sprintf("missing version for %q", typ),
		WithWrap(ErrInvalidMeta),
	)
}

// MetaMissingSchema returns a wrapped ErrMetaInvalid indicating the supplied
// Meta is missing a schema.
func MetaMissingSchema(typ any) error {
	return New(
		fmt.Sprintf("missing schema for %q", typ),
		WithWrap(ErrInvalidMeta),
	)
}
