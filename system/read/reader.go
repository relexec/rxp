package read

import (
	"context"

	"github.com/relexec/rxp/system/read/option"
	"github.com/relexec/rxp/system/read/selector"
	"github.com/relexec/rxp/types"
)

// SystemReader reads a single System from persistent storage.
type SystemReader interface {
	// SystemRead reads a single System from persistent storage.
	SystemRead(
		context.Context,
		selector.Selector,
		...option.Option,
	) (types.System, error)
}
