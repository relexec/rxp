package types

import "github.com/Masterminds/semver/v3"

// Meta contains the definition of a type and version of an Object.
type Meta interface {
	Validatable
	Differ
	// System returns the system identifier associated with the KindVersion.
	System() System
	// KindVersion returns the Meta's KindVersion.
	KindVersion() KindVersion
	// Kind returns the Meta's Kind.
	Kind() Kind
	// Version returns the Version of the Meta.
	Version() semver.Version
	// Schema returns a [jsonschema.Schema] that describes the desired state
	// fields of Objects with this KindVersion.
	Schema() Schema
	// SchemaJSON returns a string containing the [jsonschema.Schema] that
	// describes the desired state fields of Objects with this KindVersion.
	SchemaJSON() (string, error)
}
