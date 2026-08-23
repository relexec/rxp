package apirun

import (
	"time"

	apicore "github.com/relexec/rxp/api/core"
)

// Options describes optional per-request Run timeouts, retry behaviour and
// other configuration options.
type Options struct {
	// Wait describes the behaviour of the caller in waiting for the Runnable
	// to complete. Note that by default, the caller does *not* wait for the
	// Runnable to complete and the only field of the Response that is
	// guaranteed to be populated is the RequestUUID field, which may be used
	// by the caller to poll for the result of the Runnable at a later time.
	Wait apicore.Wait `json:"wait"`
	// Timeout is the duration timeout for the entire call to Run.
	Timeout time.Duration `json:"timeout"`
	// Retry configures retry behaviour for the call to Run.
	Retry apicore.Retry `json:"retry"`
}
