package write

import (
	"context"

	"github.com/relexec/rxp/system/write/option"
	"github.com/relexec/rxp/types"
)

// SystemWriter is able to write a single System to persistent storage.
type SystemWriter interface {
	// SystemWrite persists the single supplied System to backend storage.
	SystemWrite(
		context.Context,
		types.System,
		...option.Option,
	) error
}
