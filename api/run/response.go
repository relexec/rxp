package apirun

import (
	"time"
)

// Stats contains timing information and other statistics scoped to a single
// call to Run.
type Stats struct {
	// Elapsed is the total amount of wallclock time spent by the Runner to
	// executed the Run call.
	Elapsed time.Duration
	// Attempts is the number of executions made to execute the Target.
	Attempts int
}

// Response contains the result/response for the call to Run.
type Response struct {
	// RequestUUID is the Request's UUID. If the Request.UUID field is empty,
	// the rxp runtime creates a new UUID for the Request and populates this
	// field in the Response struct.
	RequestUUID string
	// Errors contains collected application-layer errors (i.e. not runtime
	// errors) that occurred during the call to Run.
	Errors []error
	// Stats contains timing information and other statistics scoped to a
	// single call to Run.
	Stats Stats
	// Out contains any state fields the Runnable added to the Response,
	// encoded as a JSON string.
	Out string
}
