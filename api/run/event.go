package apirun

import (
	"time"
)

// EventType describes the type of event that occurred in the timeline of a
// Run.
type EventType int

const (
	// EventTypeRunScheduled is when the Run was scheduled. Scheduled may be
	// different from Started if the caller has requested execution of a
	// Runnable at a future time.
	EventTypeRunScheduled EventType = iota
	// EventTypeRunStarted is when the Runnable started to execute.
	EventTypeRunStarted
	// EventTypeRunCompleted is when the Runnable completed execution
	// successfully (no application-layer terminal failures were encountered).
	EventTypeRunCompleted
	// EventTypeRunFailed is when the Runnable completed execution and
	// encountered an application-layer terminal failure.
	EventTypeRunFailed
	// EventTypeRunCanceled is when a manual call to cancel an ongoing Run was
	// made. Note that this is different from EventTypeRunTimedOut which is for
	// when the execution of a Runnable hit a configured timeout duration.
	EventTypeRunCanceled
	// EventTypeRunTimedOut is when the execution of a Runnable hit a
	// configured timeout duration. Note that logging a EventTypeRunTimedOut
	// event is *not* sufficient information to determine if the entire Run
	// request timed out. To do that, the retry configuration for the
	// RunRequest must be consulted.
	EventTypeRunTimedOut
	// EventTypeRunPaused is when a manual call to pause an ongoing Run was
	// made.
	EventTypeRunPaused
	// EventTypeRunResumed is when a manual call to resume a paused Run was
	// made.
	EventTypeRunResumed
	// EventTypeTimerStarted is when a timer is started within the execution of
	// a Runnable.
	EventTypeTimerStarted
	// EventTypeTimerFired is when a timer fires (wakes).
	EventTypeTimerFired
)

// Event represents a single event that occurred in the timeline of a Run.
type Event struct {
	// Sequence stores the Event's sequence number within the Run.
	Sequence int
	// Type is the EventType of the event that occurred.
	Type EventType
	// On is the exact time that the Event occurred in UNIX nanoseconds.
	On time.Time
}
