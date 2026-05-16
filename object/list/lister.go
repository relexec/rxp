package list

import (
	"context"

	"github.com/relexec/rxp/list"
	"github.com/relexec/rxp/list/option"
	"github.com/relexec/rxp/types"
)

// Lister lists zero or more Objects.
type Lister interface {
	// ObjectList lists zero or more Objects.
	ObjectList(
		ctx context.Context,
		expr types.Expression,
		opts ...option.Option,
	) (list.Result[types.Object], error)
}
