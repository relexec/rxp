package api

import "context"

type contextKey string

const (
	contextKeyCaller contextKey = "rxp.caller"
)

// Caller contains information about the identity of the caller.
type Caller struct {
	// Identity is the identifier of the calling entity. This will be
	// system-specific. It could be a UUID, a username, an email or any other
	// type of identifier.
	Identity string
	// System is System that the caller's API request was routed through. If
	// empty, the host System for the entity executing the API call is used.
	System *System
	// Domain is the optional Domain the caller should have their rxp API calls
	// scoped to. Calling systems may add this Domain to the API call request
	// context automatically during authentication to avoid a caller needing to
	// manually specify a Domain UUID in queries.
	Domain *Domain
}

// CallerToContext sets the supplied Caller information into the supplied
// context, returning the mutated context.
func CallerToContext(
	ctx context.Context,
	caller Caller,
) context.Context {
	return context.WithValue(ctx, contextKeyCaller, &caller)
}

// CallerFromContext returns the Caller contained in the supplied context.
// Returns nil if no Caller was set in the context.
func CallerFromContext(ctx context.Context) *Caller {
	if ctx == nil {
		return nil
	}
	if v := ctx.Value(contextKeyCaller); v != nil {
		return v.(*Caller)
	}
	return nil
}
