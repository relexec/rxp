package api

import (
	"time"

	"github.com/relexec/rxp/errors"
)

// RunOptions describes optional per-request Run timeouts, retry behaviour and
// other configuration options.
type RunOptions struct {
	// Wait describes the behaviour of the caller in waiting for the Runnable
	// to complete. Note that by default, the caller does *not* wait for the
	// Runnable to complete and the only field of the Response that is
	// guaranteed to be populated is the RequestUUID field, which may be used
	// by the caller to poll for the result of the Runnable at a later time.
	Wait Wait
	// Timeout is the duration timeout for the entire call to Run.
	Timeout time.Duration
	// Retry configures retry behaviour for the call to Run.
	Retry Retry
}

// RunTarget uniquely identifies an Object generation that contains the
// *definition* of the thing that will be executed by Run.
type RunTarget struct {
	// KindVersionName is the kind and version identifier for the type of
	// Object describing the Runnable.
	KindVersionName KindVersionName
	// System is System in which the target's Object resides. If empty, the
	// host System for the entity executing the API call is used.
	System *System
	// Domain is the optional Domain the target's Object is scoped to.
	Domain *Domain
	// UUID is the identifier of the Object describing the Runnable.
	UUID string
	// Generation is the Object generation.
	Generation Generation
}

// Validate returns an error if the RunTarget is not valid.
func (t RunTarget) Validate() error {
	if t.KindVersionName == "" {
		return errors.ErrRunTargetKindVersionNameRequired
	}
	if t.UUID == "" {
		return errors.ErrRunTargetUUIDRequired
	}
	if t.Generation == 0 {
		return errors.ErrRunTargetGenerationRequired
	}
	if t.System != nil {
		err := t.System.Validate()
		if err != nil {
			return errors.ErrRunTargetSystemInvalid
		}
	}
	if t.Domain != nil {
		err := t.Domain.Validate()
		if err != nil {
			return errors.ErrRunTargetDomainInvalid
		}
	}
	return nil
}

// RunRequest describes a single request to execute some work.
type RunRequest struct {
	// UUID is the identifier for a single Run call.
	UUID string
	// On is the UNIX nanoseconds timestamp of when the Request was made.
	On time.Time
	// Options contains per-request settings, encoded as a JSON string.
	Options string
	// Target contains the UUID and Generation of the Object that contains the
	// *definition* of the thing that will be executed by Run.
	Target RunTarget
	// Caller contains information about the calling identity.
	Caller Caller
	// In contains the value of the input parameter when calling Run, encoded
	// as a JSON string.
	In string
}

// Validate returns an error if the RunRequest is not valid.
func (r RunRequest) Validate() error {
	if r.UUID == "" {
		return errors.ErrRunRequestUUIDRequired
	}
	if r.On.IsZero() {
		return errors.ErrRunRequestOnRequired
	}
	if err := r.Caller.Validate(); err != nil {
		return err
	}
	return r.Target.Validate()
}

// RunResponse contains the result/response for the call to Run.
type RunResponse struct {
	// RequestUUID is the Request's UUID. If the Request.UUID field is empty,
	// the rxp runtime creates a new UUID for the Request and populates this
	// field in the Response struct.
	RequestUUID string
	// Errors contains collected application-layer errors (i.e. not runtime
	// errors) that occurred during the call to Run.
	Errors []error
	// Stats contains timing information and other statistics scoped to a
	// single call to Run.
	Stats RunStats
	// Out contains any state fields the Runnable added to the Response. Note
	// that for Run requests that were asynchronous (the default), this field
	// will be empty. These values are persisted to the RunLog and callers can
	// retrieve the values by polling the RunResponse via the RequestUUID.
	Out Vars
}

// RunStats contains timing information and other statistics scoped to a single
// call to Run.
type RunStats struct {
	// Elapsed is the total amount of wallclock time spent by the Runner to
	// executed the Run call.
	Elapsed time.Duration
	// Attempts is the number of executions made to execute the Target.
	Attempts int
}

// Run describes a thread of execution of some piece of work.
type Run struct {
	// req contains information about the request to execute some piece of
	// work.
	req RunRequest
	// root contains the UUID of the root Run. If this is a root Run, this will
	// be the same as req.UUID.
	root string
	// parent contains the UUID of the Run that spawned this Run, if any.
	parent string
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

// Validate returns an error if the Run is not valid.
func (r Run) Validate() error {
	err := r.req.Validate()
	if err != nil {
		return err
	}
	if r.scheduledOn.Before(r.startedOn) {
		return errors.ErrRunScheduledBeforeStart
	}
	if !r.canceledOn.IsZero() {
		if r.canceledOn.Before(r.scheduledOn) {
			return errors.ErrRunCanceledBeforeScheduled
		}
	}
	if !r.completedOn.IsZero() {
		if !r.failedOn.IsZero() {
			return errors.ErrRunCompletedAndFailed
		}
	}
	return nil
}

// Request returns the Run's Request struct.
func (r Run) Request() RunRequest {
	return r.req
}

// SetRequest sets the Run's Request struct.
func (r *Run) SetRequest(req RunRequest) {
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

// Parent returns the UUID of the Run that spawned this Run, if any.
func (r Run) Parent() string {
	return r.parent
}

// SetParent sets the UUID of the Run that spawned this Run.
func (r *Run) SetParent(parent string) {
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

// RunEventType describes the type of event that occurred in the timeline of a
// Run.
type RunEventType int

const (
	// RunEventTypeRunScheduled is when the Run was scheduled. Scheduled may be
	// different from Started if the caller has requested execution of a
	// Runnable at a future time.
	RunEventTypeRunScheduled RunEventType = iota
	// RunEventTypeRunStarted is when the Runnable started to execute.
	RunEventTypeRunStarted
	// RunEventTypeRunCompleted is when the Runnable completed execution
	// successfully (no application-layer terminal failures were encountered).
	RunEventTypeRunCompleted
	// RunEventTypeRunFailed is when the Runnable completed execution and
	// encountered an application-layer terminal failure.
	RunEventTypeRunFailed
	// RunEventTypeRunCanceled is when a manual call to cancel an ongoing Run
	// was made. Note that this is different from RunEventTypeRunTimedOut which
	// is for when the execution of a Runnable hit a configured timeout
	// duration.
	RunEventTypeRunCanceled
	// RunEventTypeRunTimedOut is when the execution of a Runnable hit a
	// configured timeout duration. Note that logging a RunEventTypeRunTimedOut
	// event is *not* sufficient information to determine if the entire Run
	// request timed out. To do that, the retry configuration for the
	// RunRequest must be consulted.
	RunEventTypeRunTimedOut
	// RunEventTypeRunPaused is when a manual call to pause an ongoing Run was
	// made.
	RunEventTypeRunPaused
	// RunEventTypeRunResumed is when a manual call to resume a paused Run was
	// made.
	RunEventTypeRunResumed
	// RunEventTypeTimerStarted is when a timer is started within the execution
	// of a Runnable.
	RunEventTypeTimerStarted
	// RunEventTypeTimerFired is when a timer fires (wakes).
	RunEventTypeTimerFired
)

// RunEvent represents a single event that occurred in the timeline of a Run.
type RunEvent struct {
	// Sequence stores the RunEvent's sequence number within the Run.
	Sequence int
	// Type is the RunEventType of the event that occurred.
	Type RunEventType
	// On is the exact time that the RunEvent occurred in UNIX nanoseconds.
	On time.Time
}
