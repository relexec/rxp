package write

import (
	"context"

	"github.com/relexec/rxp/kind/write/option"
	"github.com/relexec/rxp/types"
)

// KindWriter is able to write a single Kind to persistent storage.
type KindWriter interface {
	// KindWrite persists the single supplied Kind to backend storage.
	KindWrite(
		context.Context,
		types.Kind,
		...option.Option,
	) error
}
