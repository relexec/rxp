package rxp

import (
	"context"

	"github.com/relexec/rxp/api"
)

// DataPlane is the interface that rxp backends implement for data plane
// operations.
type DataPlane interface {
	// RunEventsWrite atomically writes the RunEvents for the supplied Run.
	RunEventsWrite(
		context.Context,
		api.Run,
		[]api.RunEvent,
	) error
}
