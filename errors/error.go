package errors

import (
	"errors"
	"fmt"
)

// Error is the base struct for all rxp-specific errors. Implements the `error`
// interface.
type Error struct {
	// code is the numeric error code. generally follows HTTP status codes.
	code int
	// msg is the error message
	msg string
	// wrapped is the error that this Error wraps, if any
	wrapped error
}

// Error returns the string error message.
func (e Error) Error() string {
	if e.wrapped != nil {
		return fmt.Sprintf("%v: %s", e.wrapped, e.msg)
	}
	return e.msg
}

// Code returns the numeric error code.
func (e Error) Code() int {
	return e.code
}

// Unwrap returns the underlying error being wrapped.
func (e Error) Unwrap() error {
	return e.wrapped
}

type option func(*Error)

// WithCode sets the returned Error's Code.
func WithCode(code int) option {
	return func(e *Error) {
		e.code = code
	}
}

// WithWrap sets the returned Error's wrapped error.
func WithWrap(wrapped error) option {
	return func(e *Error) {
		e.wrapped = wrapped
	}
}

// New returns a new Error.
func New(
	msg string,
	opts ...option,
) *Error {
	e := &Error{
		msg: msg,
	}
	for _, o := range opts {
		o(e)
	}
	var rxpErr *Error
	if e.wrapped != nil && errors.As(e.wrapped, &rxpErr) {
		e.code = rxpErr.Code()
	}
	if e.code == 0 {
		e.code = ErrCodeInternal
	}
	return e
}
