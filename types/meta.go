package types

// Meta contains the definition of a type and version of an Object.
type Meta interface {
	Validatable
	Differ
	// System returns the system identifier associated with the KindVersion.
	System() System
	// KindVersion returns the Meta's KindVersion.
	KindVersion() KindVersion
	// Namescope returns the name uniqueness constraint for Objects having this
	// KindVersion.
	Namescope() Namescope
	// Schema returns a [jsonschema.Schema] that describes the desired state
	// fields of Objects with this KindVersion.
	Schema() Schema
	// SchemaJSON returns a string containing the [jsonschema.Schema] that
	// describes the desired state fields of Objects with this KindVersion.
	SchemaJSON() (string, error)
}
