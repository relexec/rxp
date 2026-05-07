package read

import (
	"context"

	"github.com/relexec/rxp/namespace/read/option"
	"github.com/relexec/rxp/namespace/read/selector"
	"github.com/relexec/rxp/types"
)

// NamespaceReader reads a single Namespace from persistent storage.
type NamespaceReader interface {
	// NamespaceRead reads a single Namespace from persistent storage.
	NamespaceRead(
		context.Context,
		selector.Selector,
		...option.Option,
	) (types.Namespace, error)
}
