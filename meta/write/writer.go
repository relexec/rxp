package write

import (
	"context"

	"github.com/relexec/rxp/types"
)

// MetaWriter writes a single Meta to persistent storage.
type MetaWriter interface {
	// MetaWrite persists the single supplied Meta to backend storage.
	MetaWrite(
		context.Context,
		types.Meta,
		...Option,
	) error
}
