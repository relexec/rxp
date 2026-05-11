package author

import (
	"github.com/relexec/rxp/object"
)

// New returns a new Author [Object].
func New(opts ...object.Option) *object.Object {
	return object.New(opts...)
}
