package write

import (
	"context"

	"github.com/relexec/rxp/object/write/option"
	"github.com/relexec/rxp/types"
)

// ObjectWriter is able to write a single Object to persistent storage.
type ObjectWriter interface {
	// ObjectWrite persists the single supplied Object to backend storage.
	ObjectWrite(
		context.Context,
		types.Object,
		...option.Option,
	) error
}
