package errors

import "fmt"

var (
	ErrPredicateInvalid       = New("invalid predicate", WithCode(ErrCodeBadRequest))
	ErrPredicateValueRequired = New("value required", WithWrap(ErrPredicateInvalid))
)

// PredicateInvalid returns an Error wrapping ErrPredicateInvalid with the
// supplied message.
func PredicateInvalid(msg string, opts ...option) error {
	e := &Error{
		code:    ErrCodeBadRequest,
		msg:     msg,
		wrapped: ErrPredicateInvalid,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

func PredicateUnsupportedValueType(typ any) error {
	return New(
		fmt.Sprintf("unsupported type for value %T", typ),
		WithWrap(ErrPredicateInvalid),
	)
}
