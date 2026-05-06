package write

import (
	"context"

	"github.com/relexec/rxp/types"
)

// ObjectWriter is able to write a *single* Object to persistent storage.
type ObjectWriter interface {
	// WriteObject persists the *single* supplied Object to backend storage.
	//
	// WriteObject returns the Object that was saved, mutated with any desired
	// state defaults.
	WriteObject(
		context.Context,
		types.Object,
		...Option,
	) (types.Object, error)
}
