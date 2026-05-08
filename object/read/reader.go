package read

import (
	"context"

	"github.com/relexec/rxp/object/read/option"
	"github.com/relexec/rxp/object/read/selector"
	"github.com/relexec/rxp/types"
)

// ObjectReader reads a single [types.Object] from persistent storage.
type ObjectReader interface {
	// ObjectRead reads a single [types.Object] from persistent storage.
	ObjectRead(
		ctx context.Context,
		sel selector.Selector,
		opts ...option.Option,
	) (types.Object, error)
}
