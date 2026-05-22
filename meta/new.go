package meta

import (
	"encoding/json"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/meta/schema"
	"github.com/relexec/rxp/system"
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
func WithSystem(system *system.System) Option {
	return func(m *Meta) {
		m.system = system
	}
}

// WithKind sets the Meta's Kind.
func WithKind(k *kind.Kind) Option {
	return func(m *Meta) {
		m.kind = k
	}
}

// WithVersion sets the Meta's Version.
func WithVersion(ver semver.Version) Option {
	return func(m *Meta) {
		m.version = ver
	}
}

// WithSchema sets the Meta's Schema.
func WithSchema(schema *schema.Schema) Option {
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
