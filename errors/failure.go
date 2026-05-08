package errors

import (
	"fmt"
	"net/http"
)

var (
	ErrCodeInternal           = http.StatusInternalServerError
	ErrCodeBadRequest         = http.StatusBadRequest
	ErrCodeConflict           = http.StatusConflict
	ErrCodeNotFound           = http.StatusNotFound
	ErrCodePreconditionFailed = http.StatusPreconditionFailed
	ErrCodeTooManyRequests    = http.StatusTooManyRequests
)

var (
	ErrNotFound           = New("not found", WithCode(ErrCodeNotFound))
	ErrMissingIdentity    = New("missing identity", WithCode(ErrCodeBadRequest))
	ErrConflict           = New("conflict", WithCode(ErrCodeConflict))
	ErrPreconditionFailed = New("precondition failed", WithCode(ErrCodePreconditionFailed))
)

var (
	ErrInvalidReadRequest = New("invalid read request", WithCode(ErrCodeBadRequest))
)

// Internal returns an Error with a 500 Internal Server Error code and the
// supplied message.
func Internal(msg string, opts ...option) error {
	e := &Error{
		code: ErrCodeInternal,
		msg:  msg,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

// ExpectedToExist returns an Error wrapping ErrPreconditionFailed that
// explains the supplied thing was expected to exist.
func ExpectedToExist(subject any, opts ...option) error {
	e := &Error{
		code:    ErrCodePreconditionFailed,
		msg:     fmt.Sprintf("expected %q to exist", subject),
		wrapped: ErrPreconditionFailed,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

// ExpectedNotToExist returns an Error wrapping ErrPreconditionFailed that
// explains the supplied thing was expected NOT to exist.
func ExpectedNotToExist(subject any, opts ...option) error {
	e := &Error{
		code:    ErrCodePreconditionFailed,
		msg:     fmt.Sprintf("expected %q not to exist", subject),
		wrapped: ErrPreconditionFailed,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

// ExpectedGeneration returns an Error wrapping ErrPreconditionFailed that
// explains the supplied thing was expected to have a different generation.
func ExpectedGeneration(subject any, expected any, got any, opts ...option) error {
	e := &Error{
		code: ErrCodePreconditionFailed,
		msg: fmt.Sprintf(
			"expected %q to have generation %d but found %d",
			subject, expected, got,
		),
		wrapped: ErrPreconditionFailed,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

// DuplicateKey returns an Error wrapping ErrConflict that explains the supplied
// thing violated a unique key constraint.
func DuplicateKey(typ any, key any, value any, opts ...option) error {
	e := &Error{
		code:    ErrCodeConflict,
		msg:     fmt.Sprintf("%q already exists with key %q of %q", typ, key, value),
		wrapped: ErrConflict,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

// DuplicateName returns an Error wrapping ErrConflict that explains the supplied
// thing violated a name uniqueness constraint.
func DuplicateName(typ any, name any, opts ...option) error {
	e := &Error{
		code:    ErrCodeConflict,
		msg:     fmt.Sprintf("%q already exists with name %q", typ, name),
		wrapped: ErrConflict,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}
