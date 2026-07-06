package runners

import (
	"context"

	"github.com/relexec/rxp/run/request"
	"github.com/relexec/rxp/run/response"
)

// Echo is a runnable function that simply writes the request caller
// information into the response's Out field at the "request.caller" key.
func Echo(
	ctx context.Context,
	req request.Request,
	resp *response.Response,
) error {
	resp.Out.Set("request.caller", req.Caller)
	return nil
}
