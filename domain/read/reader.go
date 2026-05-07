package read

import (
	"context"

	"github.com/relexec/rxp/domain/read/option"
	"github.com/relexec/rxp/domain/read/selector"
	"github.com/relexec/rxp/types"
)

// DomainReader reads a single Domain from persistent storage.
type DomainReader interface {
	// DomainRead reads a single Domain from persistent storage.
	DomainRead(
		context.Context,
		selector.Selector,
		...option.Option,
	) (types.Domain, error)
}
