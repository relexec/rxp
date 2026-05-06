package read

import (
	"context"

	"github.com/relexec/rxp/types"
)

// MetaReader reads a single Meta from persistent storage.
type MetaReader interface {
	// MetaRead reads a single Meta from persistent storage.
	MetaRead(
		context.Context,
		types.KindVersion,
		...Option,
	) (types.Meta, error)
}
