package errors

import (
	"fmt"
)

var (
	ErrInvalidObject = New("invalid object", WithCode(ErrCodeBadRequest))
	ErrNilObject     = New("nil object parameter", WithWrap(ErrInvalidObject))
)

// ObjectMissingIdentifier returns a wrapped ErrObjectInvalid indicating the
// supplied Object type is missing an identifier.
func ObjectMissingIdentifier(typ any) error {
	return New(
		fmt.Sprintf("missing identifier for %q", typ),
		WithWrap(ErrInvalidObject),
	)
}

// ObjectMissingName returns a wrapped ErrObjectInvalid indicating the supplied
// Object is missing a name.
func ObjectMissingName(typ any, id any) error {
	return New(
		fmt.Sprintf("missing name for %q (%s)", typ, id),
		WithWrap(ErrInvalidObject),
	)
}
