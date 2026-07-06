package run

import (
	"context"

	"github.com/relexec/rxp/run/request"
	"github.com/relexec/rxp/run/response"
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
		request.Request,
		*response.Response, // mutated by Run
	) error
}

// RunnableFunc adapts a pure function or method to be a Runnable
type RunnableFunc struct {
	r func(context.Context, request.Request, *response.Response) error
}

// Run executes a single Request, returning any runtime error that may have
// occurred. Runtime errors are not application-layer errors. They represent
// terminal conditions that stop the flow of execution entirely.  The
// Response.Errors field will contain any application-layer errors (validation
// failures, duplicate/conflict errors, timeouts, etc).
func (f RunnableFunc) Run(
	ctx context.Context,
	req request.Request,
	resp *response.Response,
) error {
	return f.r(ctx, req, resp)
}

// RunnableFrom returns an Runnable from a pure function or method.
func RunnableFrom(
	r func(context.Context, request.Request, *response.Response) error,
) Runnable {
	return RunnableFunc{r}
}
