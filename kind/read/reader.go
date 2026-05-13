package read

import (
	"context"

	"github.com/relexec/rxp/kind/read/option"
	"github.com/relexec/rxp/kind/read/selector"
	"github.com/relexec/rxp/types"
)

// KindReader reads a single Kind from persistent storage.
type KindReader interface {
	// KindRead reads a single Kind from persistent storage.
	KindRead(
		context.Context,
		selector.Selector,
		...option.Option,
	) (types.Kind, error)
}
