package meta

import (
	"encoding/json"

	"github.com/google/jsonschema-go/jsonschema"

	"github.com/relexec/rxp/types"
)

// Option modifies a Meta returned from New.
type Option func(*Meta)

// New returns a new [Meta].
func New(opts ...Option) *Meta {
	m := &Meta{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithSystem sets the Meta's System.
func WithSystem(system types.System) Option {
	return func(m *Meta) {
		m.system = system
	}
}

// WithKindVersion sets the Meta's KindVersion.
func WithKindVersion(kv types.KindVersion) Option {
	return func(m *Meta) {
		m.kindVersion = kv
	}
}

// WithNamescope sets the Meta's Namescope.
func WithNamescope(namescope types.Namescope) Option {
	return func(m *Meta) {
		m.namescope = namescope
	}
}

// WithSchema sets the Meta's Schema.
func WithSchema(schema *jsonschema.Schema) Option {
	return func(m *Meta) {
		m.schema = schema
	}
}

// WithSchemaJSON sets the Meta's Schema JSON string.
func WithSchemaJSON(schemaJSON string) Option {
	return func(m *Meta) {
		m.schemaJSON = schemaJSON
		if m.schema == nil {
			_ = json.Unmarshal([]byte(schemaJSON), m.schema)
		}
	}
}
