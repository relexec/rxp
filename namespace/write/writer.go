package write

import (
	"context"

	"github.com/relexec/rxp/namespace/write/option"
	"github.com/relexec/rxp/types"
)

// NamespaceWriter is able to write a single Namespace to persistent storage.
type NamespaceWriter interface {
	// NamespaceWrite persists the single supplied Namespace to backend storage.
	NamespaceWrite(
		context.Context,
		types.Namespace,
		...option.Option,
	) error
}
