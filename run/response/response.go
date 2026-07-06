package response

import (
	"time"

	"github.com/relexec/rxp/api"
)

// Response contains the result/response for the call to Run.
type Response struct {
	// RequestUUID is the Request's UUID. If the Request.UUID field is empty,
	// the the rxp runtime creates a new UUID for the Request and populates
	// this field in the Response struct.
	RequestUUID string
	// Errors contains collected application-layer errors (i.e. not runtime
	// errors) that occurred during the call to Run.
	Errors []error
	// Stats contains timing information and other statistics scoped to a
	// single call to Run.
	Stats Stats
	// Out contains any state fields the Runnable added to the Response. Note
	// that for Run requests that were asynchronous (the default), this field
	// will be empty. These values are persisted to the RunLog and callers can
	// retrieve the values by polling the RunResponse via the RequestUUID.
	Out api.Vars
}

// Stats contains timing information and other statistics scoped to a single
// call to Run.
type Stats struct {
	// Elapsed is the total amount of wallclock time spent by the Runner to
	// executed the Run call.
	Elapsed time.Duration
	// Attempts is the number of attempts to execute the Runnable.
	Attempts int
}
