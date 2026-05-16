package list

import (
	"context"

	"github.com/relexec/rxp/list/option"
	"github.com/relexec/rxp/types"
)

// ObjectLister lists zero or more [types.Object] from persistent storage.
type ObjectLister interface {
	// ObjectList lists zero or more [types.Object] from persistent storage.
	ObjectList(
		ctx context.Context,
		expr types.Expression,
		opts ...option.Option,
	) (types.ListResult[types.Object], error)
}
