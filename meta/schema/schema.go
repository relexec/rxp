package schema

import (
	"github.com/google/jsonschema-go/jsonschema"
)

// Schema wraps a [jsonschema.Schema] and implements the [types.Differ]
// interface.
type Schema struct {
	jsonschema.Schema
}
