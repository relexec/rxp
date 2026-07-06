package runners_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/relexec/rxp/run"
	"github.com/relexec/rxp/run/request"
	"github.com/relexec/rxp/run/response"
	"github.com/relexec/rxp/testing/runners"
	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	require := require.New(t)
	r := run.RunnableFrom(runners.Echo)

	ctx := context.TODO()
	caller := request.Caller{
		Identity: "user",
	}
	req := request.Request{
		UUID:   uuid.NewString(),
		Caller: caller,
	}

	var resp response.Response
	err := r.Run(ctx, req, &resp)
	require.Nil(err)
}
