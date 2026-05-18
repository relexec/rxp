package read

import (
	"context"

	"github.com/relexec/rxp/read/option"
	"github.com/relexec/rxp/types"
)

// SystemReader reads a single System from persistent storage.
type SystemReader interface {
	// SystemRead reads a single System from persistent storage.
	SystemRead(
		context.Context,
		types.Selector,
		...option.Option,
	) (types.System, error)
}
