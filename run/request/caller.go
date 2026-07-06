package request

import (
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/system"
)

// Caller contains information about the identity of the caller.
type Caller struct {
	// Identity is the identifier of the calling entity. This will be
	// system-specific. It could be a UUID, a username, an email or any other
	// type of identifier.
	Identity string
	// System is the UUID of the System that the caller's Run request was
	// routed through. If empty, the host System for the entity executing the
	// Runnable is used.
	System *system.System
	// Domain is the UUID of the Domain that the caller's Run request was
	// associated with. If empty, the Domain associated with the entity
	// executing the Runnable is used.
	Domain *domain.Domain
}
