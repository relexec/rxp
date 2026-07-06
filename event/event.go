package event

import (
	"time"
)

// Event describes something that has already occurred.
type Event struct {
	// Sequence stores the Event's sequence number within the EventSet.
	Sequence int
	// Type is the type of the Event.
	Type Type
	// When is the exact time that the Event occurred in UNIX nanoseconds.
	When time.Time
}
