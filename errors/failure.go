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
	ErrInvalidReadRequest  = New("invalid read request", WithCode(ErrCodeBadRequest))
	ErrInvalidWriteRequest = New("invalid write request", WithCode(ErrCodeBadRequest))
	ErrInvalidQueryRequest = New("invalid query request", WithCode(ErrCodeBadRequest))
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

// RequiredParameterNil returns an Error with a 400 Bad Request Error code that
// explains the supplied required non-nil parameter was nil.
func RequiredParameterNil(param string, opts ...option) error {
	e := &Error{
		code: ErrCodeBadRequest,
		msg:  fmt.Sprintf("require parameter %q was nil or empty", param),
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

// ExpectedFirstVersionInSeries returns an Error wrapping ErrPreconditionFailed
// that explains the supplied kindversion was expected to have a zero minor and
// patch version number.
func ExpectedFirstVersionInSeries(kv any, opts ...option) error {
	e := &Error{
		code: ErrCodePreconditionFailed,
		msg: fmt.Sprintf(
			"expected %q to have minor and patch version of 0", kv,
		),
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
