package kindversion

import (
	"encoding/json"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion/schema"
	"github.com/relexec/rxp/system"
)

// Option modifies a KindVersion returned from New.
type Option func(*KindVersion)

// New returns a new [KindVersion].
func New(opts ...Option) *KindVersion {
	kv := &KindVersion{}
	for _, opt := range opts {
		opt(kv)
	}
	return kv
}

// WithSystem sets the KindVersion's System.
func WithSystem(system *system.System) Option {
	return func(kv *KindVersion) {
		kv.system = system
	}
}

// WithKind sets the KindVersion's Kind.
func WithKind(k *kind.Kind) Option {
	return func(kv *KindVersion) {
		kv.kind = k
	}
}

// WithVersion sets the KindVersion's Version.
func WithVersion(ver semver.Version) Option {
	return func(kv *KindVersion) {
		kv.version = ver
	}
}

// WithSchema sets the KindVersion's Schema.
func WithSchema(schema *schema.Schema) Option {
	return func(kv *KindVersion) {
		kv.schema = schema
	}
}

// WithSchemaJSON sets the KindVersion's Schema JSON string.
func WithSchemaJSON(schemaJSON string) Option {
	return func(kv *KindVersion) {
		kv.schemaJSON = schemaJSON
		if kv.schema == nil {
			_ = json.Unmarshal([]byte(schemaJSON), kv.schema)
		}
	}
}
