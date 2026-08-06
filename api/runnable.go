package api

import (
	"context"
)

// Runnable standardizes the callable interface for executing code.
type Runnable interface {
	// Run executes a single Request, returning any runtime error that may have
	// occurred. Runtime errors are not application-layer errors. They
	// represent terminal conditions that stopped the flow of execution
	// entirely. The Response.Errors field will contain any application-layer
	// errors (validation failures, duplicate/conflict errors, timeouts, etc).
	Run(
		context.Context,
		RunRequest,
		*RunResponse, // mutated by Run
	) error
}

// RunnableFunc adapts a pure function or method to be a Runnable
type RunnableFunc struct {
	r func(context.Context, RunRequest, *RunResponse) error
}

// Run executes a single Request, returning any runtime error that may have
// occurred. Runtime errors are not application-layer errors. They represent
// terminal conditions that stop the flow of execution entirely.  The
// Response.Errors field will contain any application-layer errors (validation
// failures, duplicate/conflict errors, timeouts, etc).
func (f RunnableFunc) Run(
	ctx context.Context,
	req RunRequest,
	resp *RunResponse,
) error {
	return f.r(ctx, req, resp)
}

// RunnableFrom returns an Runnable from a pure function or method.
func RunnableFrom(
	r func(context.Context, RunRequest, *RunResponse) error,
) Runnable {
	return RunnableFunc{r}
}
