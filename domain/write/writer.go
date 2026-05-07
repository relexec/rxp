package write

import (
	"context"

	"github.com/relexec/rxp/domain/write/option"
	"github.com/relexec/rxp/types"
)

// DomainWriter is able to write a single Domain to persistent storage.
type DomainWriter interface {
	// DomainWrite persists the single supplied Domain to backend storage.
	DomainWrite(
		context.Context,
		types.Domain,
		...option.Option,
	) error
}
