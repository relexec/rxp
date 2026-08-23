package rxp

import (
	"context"

	apirun "github.com/relexec/rxp/api/run"
)

// DataPlane is the interface that rxp backends implement for data plane
// operations.
type DataPlane interface {
	// RunEventsWrite atomically writes the Events for the supplied Run.
	RunEventsWrite(
		context.Context,
		apirun.Run,
		[]apirun.Event,
	) error
}
