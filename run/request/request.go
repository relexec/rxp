package request

import (
	"time"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
)

// Request describes a single request to execute some work.
type Request struct {
	// UUID is the identifier for a single Run call.
	UUID string
	// On is the UNIX nanoseconds timestamp of when the Request was made.
	On time.Time
	// Options contains per-request settings.
	Options Options
	// Target contains the *definition* of the thing that will be executed by
	// Run.
	Target *api.Object
	// Caller contains information about the calling identity.
	Caller api.Caller
	// In contains the value of the input parameter when calling Run.
	In api.Vars
}

// Validate returns an error if the Request is not valid.
func (r Request) Validate() error {
	if r.UUID == "" {
		return errors.ErrRunRequestUUIDRequired
	}
	if r.Target == nil {
		return errors.ErrRunRequestTargetRequired
	}
	if r.On.IsZero() {
		return errors.ErrRunRequestOnRequired
	}
	return nil
}
