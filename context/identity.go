package context

import "context"

type contextKey string

const (
	contextKeyIdentity = "rxp.identity"
)

// SetIdentity sets the supplied Identity into the supplied context.
func SetIdentity(ctx context.Context, identity string) context.Context {
	return context.WithValue(ctx, contextKeyIdentity, identity)
}

// Identity returns the Identity contained in the supplied context.
func Identity(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v := ctx.Value(contextKeyIdentity); v != nil {
		return v.(string)
	}
	return ""
}
