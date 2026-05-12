package errors

import (
	"fmt"
)

var (
	ErrInvalidObject            = New("invalid object", WithCode(ErrCodeBadRequest))
	ErrNilObject                = New("nil object parameter", WithWrap(ErrInvalidObject))
	ErrObjectMissingKindVersion = New("object missing kindversion", WithWrap(ErrInvalidObject))
	ErrObjectDomainRequired     = New("domain required", WithWrap(ErrInvalidObject))
	ErrObjectNamespaceRequired  = New("namespace required", WithWrap(ErrInvalidObject))
)

// ObjectMissingUUID returns a wrapped ErrObjectInvalid indicating the
// supplied Object type is missing an identifier.
func ObjectMissingUUID(typ any) error {
	return New(
		fmt.Sprintf("missing uuid for %q", typ),
		WithWrap(ErrInvalidObject),
	)
}

// ObjectMissingName returns a wrapped ErrObjectInvalid indicating the supplied
// Object is missing a name.
func ObjectMissingName(typ any, uuid any) error {
	return New(
		fmt.Sprintf("missing name for %q (%s)", typ, uuid),
		WithWrap(ErrInvalidObject),
	)
}
