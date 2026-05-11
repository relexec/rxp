package book

import (
	"github.com/relexec/rxp/object"
)

// New returns a new Book [Object]
func New(opts ...object.Option) *object.Object {
	return object.New(opts...)
}
