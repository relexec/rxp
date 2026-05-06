package read

import (
	"context"

	"github.com/relexec/rxp/types"
)

// ObjectReader reads a single [types.Object] from persistent storage.
type ObjectReader interface {
	// ObjectRead reads a single [types.Object] from persistent storage.
	ObjectRead(
		ctx context.Context,
		kv types.KindVersion,
		sel types.Selector,
		opts ...Option,
	) (types.Object, error)
}
