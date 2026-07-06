package api

import "time"

// Retry configures retry behaviour.
type Retry struct {
	// MaxAttempts is the number of attempts to try.
	MaxAttempts int
	// Backoff configures the exponential backoff retry behaviour.
	Backoff Backoff
}

// Backoff configures the backoff retry behaviour.
type Backoff struct {
	// Coefficient is the exponential backoff coefficient. If zero, defaults to
	// 2.0.
	Coefficient float64
	// InitialIntervalDuration is the duration of the first backoff interval.
	// If zero, defaults to 1 second.
	InitialIntervalDuration time.Duration
	// MaxIntervalDuration is the maximum duration of a backoff interval. If
	// zero, defaults to 100 times the initial interval duration.
	MaxIntervalDuration time.Duration
}
