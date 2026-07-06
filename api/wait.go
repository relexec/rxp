package api

import "time"

// Wait describes waiter timeouts and behaviour.
type Wait struct {
	// Timeout is the total amount of time that the Waiter will wait for a
	// response. The default value is 0 which means the waiter does *not* wait
	// for a response and will poll for the results of a call at a later time.
	Timeout time.Duration
}
