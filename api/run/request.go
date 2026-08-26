package apirun

import (
	"time"

	apicore "github.com/relexec/rxp/api/core"
	apierrors "github.com/relexec/rxp/api/errors"
)

// Request describes a single request to execute some work.
type Request struct {
	// UUID is the identifier for a single Run call.
	UUID string
	// On is the UNIX nanoseconds timestamp of when the Request was made.
	On time.Time
	// Options contains per-request settings, encoded as a JSON string.
	Options string
	// Target contains the UUID and Generation of the Object that contains the
	// *definition* of the thing that will be executed by Run.
	Target Target
	// Caller contains information about the calling identity.
	Caller apicore.Caller
	// In contains the values of the input parameters when calling Run, encoded
	// as a JSON string.
	In string
}

// Validate returns an error if the Request is not valid.
func (r Request) Validate() error {
	if r.UUID == "" {
		return apierrors.ErrRunRequestUUIDRequired
	}
	if r.On.IsZero() {
		return apierrors.ErrRunRequestOnRequired
	}
	if err := r.Caller.Validate(); err != nil {
		return err
	}
	return r.Target.Validate()
}
