package schema

import (
	"github.com/google/jsonschema-go/jsonschema"
)

// Schema wraps a [jsonschema.Schema] and provides the ability to diff between
// two Schemas.
type Schema struct {
	jsonschema.Schema
}
