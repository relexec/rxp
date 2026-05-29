package errors

import (
	"fmt"
)

var (
	ErrInvalidKindVersion = New("invalid kindversion", WithCode(ErrCodeBadRequest))
)

// KindVersionMissingKind returns a wrapped ErrKindVersionInvalid indicating the
// supplied KindVersion type is missing an identifier.
func KindVersionMissingKind() error {
	return New("missing kind", WithWrap(ErrInvalidKindVersion))
}

// KindVersionMissingVersion returns a wrapped ErrKindVersionInvalid indicating the supplied
// KindVersion is missing a version.
func KindVersionMissingVersion(typ any) error {
	return New(
		fmt.Sprintf("missing version for %q", typ),
		WithWrap(ErrInvalidKindVersion),
	)
}

// KindVersionMissingSchema returns a wrapped ErrKindVersionInvalid indicating the supplied
// KindVersion is missing a schema.
func KindVersionMissingSchema(typ any) error {
	return New(
		fmt.Sprintf("missing schema for %q", typ),
		WithWrap(ErrInvalidKindVersion),
	)
}
