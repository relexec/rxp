package run

import (
	"time"

	"github.com/relexec/rxp/run/request"
)

// Run describes the status of a single execution of some piece of work.
type Run struct {
	// req contains information about the request to execute some piece of
	// work.
	req request.Request
	// root contains the UUID of the root Run. If this is a root Run, this will
	// be the same as req.UUID.
	root string
	// parent points at the Run that spawned this Run, if any.
	parent *Run
	// scheduledOn stores the UNIX nanoseconds for when the Run was scheduled
	// to be executed.
	scheduledOn time.Time
	// startedOn stores the UNIX nanoseconds for when the Run first started
	// executing.
	startedOn time.Time
	// completeddOn stores the UNIX nanoseconds for when the Run completed
	// execution successfully.
	completedOn time.Time
	// faileddOn stores the UNIX nanoseconds for when the Run failed execution.
	failedOn time.Time
	// canceledOn stores the UNIX nanoseconds for when the Run was canceled.
	canceledOn time.Time
	// pausedOn stores the UNIX nanoseconds for when the Run was last paused.
	pausedOn time.Time
	// resumedOn stores the UNIX namoseconds for when the Run was last resumed.
	resumedOn time.Time
}

// Request returns the Run's Request struct.
func (r Run) Request() request.Request {
	return r.req
}

// SetRequest sets the Run's Request struct.
func (r *Run) SetRequest(req request.Request) {
	r.req = req
}

// UUID returns the Run's globally-unique identifier, which is the Run's
// Request UUID.
func (r Run) UUID() string {
	return r.req.UUID
}

// Root returns the UUID of the root Run. If this is the root Run, the returned
// value will be the same as the value returned by UUID().
func (r Run) Root() string {
	if r.root == "" {
		return r.req.UUID
	}
	return r.root
}

// IsRoot returns true if this Run is the root Run. A root Run is the outermost
// execution of some piece of work.
func (r Run) IsRoot() bool {
	return r.root == "" || r.root == r.req.UUID
}

// SetRoot sets the UUID of the root Run. A root Run is the outermost execution
// of some piece of work.
func (r *Run) SetRoot(root string) {
	r.root = root
}

// Parent returns the Run that spawned this Run, if any.
func (r Run) Parent() *Run {
	return r.parent
}

// SetParent sets the Run that spawned this Run.
func (r *Run) SetParent(parent *Run) {
	r.parent = parent
}

// RequestedOn returns the UNIX nanoseconds for when the Run was
// requested/created.
func (r Run) RequestedOn() time.Time {
	return r.req.On
}

// ScheduledOn returns the UNIX nanoseconds for when the Run was scheduled to
// be executed. For Run's where a future scheduled on time was not supplied,
// this will exactly match the RequestedOn timestamp.
func (r Run) ScheduledOn() time.Time {
	if r.scheduledOn.IsZero() {
		return r.RequestedOn()
	}
	return r.scheduledOn
}

// SetScheduledOn sets the UNIX nanoseconds for when the Run was scheduled to
// be executed.
func (r *Run) SetScheduledOn(ts time.Time) {
	r.scheduledOn = ts
}

// StartedOn returns the UNIX nanoseconds for when the Run first started
// executing.
func (r Run) StartedOn() time.Time {
	return r.startedOn
}

// SetStartedOn sets the UNIX nanoseconds for when the Run first started
// executing.
func (r *Run) SetStartedOn(ts time.Time) {
	r.startedOn = ts
}

// CanceledOn returns the UNIX nanoseconds for when the Run was canceled.
func (r Run) CanceledOn() time.Time {
	return r.canceledOn
}

// Canceled returns true if the Run was canceled.
func (r Run) Canceled() bool {
	return !r.canceledOn.IsZero()
}

// SetCanceledOn sets the UNIX nanoseconds for when the Run was canceled.
func (r *Run) SetCanceledOn(ts time.Time) {
	r.canceledOn = ts
}

// CompletedOn returns the UNIX nanoseconds for when the Run completed
// successfully.
func (r Run) CompletedOn() time.Time {
	return r.completedOn
}

// Completed returns true if the run completed successfully.
func (r Run) Completed() bool {
	return !r.completedOn.IsZero()
}

// SetCompletedOn sets the UNIX nanoseconds for when the Run ended.
func (r *Run) SetCompletedOn(ts time.Time) {
	r.completedOn = ts
}

// FailedOn returns the UNIX nanoseconds for when the Run failed.
func (r Run) FailedOn() time.Time {
	return r.failedOn
}

// Failed returns true if the run failed.
func (r Run) Failed() bool {
	return !r.failedOn.IsZero()
}

// SetFailedOn sets the UNIX nanoseconds for when the Run failed.
func (r *Run) SetFailedOn(ts time.Time) {
	r.failedOn = ts
}

// PausedOn returns the UNIX nanoseconds for when the Run was last paused.
func (r Run) PausedOn() time.Time {
	return r.pausedOn
}

// SetPausedOn sets the UNIX nanoseconds for when the Run was last paused.
func (r *Run) SetPausedOn(ts time.Time) {
	r.pausedOn = ts
}

// ResumedOn returns the UNIX nanoseconds for when the Run was last resumed.
func (r Run) ResumedOn() time.Time {
	return r.resumedOn
}

// SetResumedOn sets the UNIX nanoseconds for when the Run was last resumed.
func (r *Run) SetResumedOn(ts time.Time) {
	r.resumedOn = ts
}

// Paused returns true if the Run is currently paused.
func (r Run) Paused() bool {
	if r.Ended() || r.Canceled() {
		return false
	}
	return r.pausedOn.IsZero() || !r.resumedOn.IsZero()
}

// Ended returns true if the Run has ended (either completed successfully,
// completed with a failure or canceled).
func (r Run) Ended() bool {
	return !r.completedOn.IsZero() || !r.failedOn.IsZero() || !r.canceledOn.IsZero()
}

// InProgress returns true if the Run has not ended and is not paused.
func (r Run) InProgress() bool {
	return !r.Ended() && !r.Paused()
}
