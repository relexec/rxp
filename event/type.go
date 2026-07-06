package event

// Type is the type of Event in rxp.
type Type int

const (
	// TypeTaskScheduled occurs when a request to perform some work in the
	// future is persisted to storage.
	TypeTaskScheduled Type = 1
	// TypeTaskStarted occurs when a request to perform some work in the
	// past or at this exact moment in time is persisted to storage.
	TypeTaskStarted = 2
	// TypeTaskCompleted occurs when some work completed with no
	// application-layer errors.
	TypeTaskCompleted = 3
	// TypeTaskFailed occurs when some work failed to complete due to an
	// application-layer error.
	TypeTaskFailed = 4
	// TypeTaskCanceled occurs when a request to cancel some ongoing work
	// is persisted to storage.
	TypeTaskCanceled = 5
	// TypeTaskPaused occurs when a request to pause some ongoing work
	// is persisted to storage.
	TypeTaskPaused = 6
	// TypeTaskResumed occurs when a request to resume some paused work
	// is persisted to storage.
	TypeTaskResumed = 7
)

// String returns the string representation of the Type.
func (t Type) String() string {
	switch t {
	case TypeTaskScheduled:
		return "task.scheduled"
	case TypeTaskStarted:
		return "task.started"
	case TypeTaskCompleted:
		return "task.completed"
	case TypeTaskFailed:
		return "task.failed"
	case TypeTaskCanceled:
		return "task.canceled"
	case TypeTaskPaused:
		return "task.paused"
	case TypeTaskResumed:
		return "task.resumed"
	default:
		return "unknown"
	}
}
