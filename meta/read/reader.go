package read

import (
	"context"

	"github.com/relexec/rxp/meta/read/option"
	"github.com/relexec/rxp/meta/read/selector"
	"github.com/relexec/rxp/types"
)

// MetaReader reads a single Meta from persistent storage.
type MetaReader interface {
	// MetaRead reads a single Meta from persistent storage.
	MetaRead(
		context.Context,
		selector.Selector,
		...option.Option,
	) (types.Meta, error)
}
