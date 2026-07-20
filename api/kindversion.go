package api

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/kind/kindversion/schema"
)

type KindVersion struct {
	// system contains the System containing the KindVersion.
	system *System
	// kind is the [api.Kind] that identifies the type of Objects represented
	// by this KindVersion.
	kind *Kind
	// version is the [semver.Version] that identifies the specific version of
	// the Kind of Objects represented by this KindVersion.
	version semver.Version
	// schema is the [jsonschema.Schema] that describes the Spec field
	// composition of Object with this Kind+Version.
	schema *schema.Schema
	// schemaJSON stores a cache of the marshaled JSON for the
	// [jsonschema.Schema] that describes the Spec field composition of the
	// Objects with this Kind+Version.
	schemaJSON string
}

// Validate returns an error if the KindVersion is not valid.
func (kv KindVersion) Validate() error {
	k := kv.kind
	if k == nil {
		return errors.KindVersionMissingKind()
	}
	err := k.Validate()
	if err != nil {
		return err
	}
	if kv.schema == nil {
		return errors.KindVersionMissingSchema(kv.Name())
	}
	return nil
}

// System returns the System of the KindVersion.
func (kv KindVersion) System() *System {
	return kv.system
}

// SetSystem sets the System of KindVersion.
func (kv *KindVersion) SetSystem(system *System) {
	kv.system = system
}

// Name returns the KindVersionName of the KindVersion.
func (kv KindVersion) Name() KindVersionName {
	return NewKindVersionName(kv.kind.Name(), kv.version)
}

// Kind returns the Kind of the KindVersion.
func (kv KindVersion) Kind() *Kind {
	return kv.kind
}

// SetKind sets the Kind of the KindVersion.
func (kv *KindVersion) SetKind(k *Kind) {
	kv.kind = k
}

// Version returns the Version of the KindVersion.
func (kv KindVersion) Version() semver.Version {
	return kv.version
}

// SetKind sets the Version of the KindVersion.
func (kv *KindVersion) SetVersion(ver semver.Version) {
	kv.version = ver
}

// Schema returns a [jsonschema.Schema] that describes the desired state fields
// of Objects with this KindVersion.
func (kv KindVersion) Schema() *schema.Schema {
	return kv.schema
}

// SetSchema sets the [jsonschema.Schema] that describes the desired state
// fields of Objects with this KindVersion.
func (kv *KindVersion) SetSchema(schema *schema.Schema) {
	kv.schema = schema
}

// SchemaJSON returns a string containing the [jsonschema.Schema] that
// describes the desired state fields of Objects with this KindVersion.
func (kv *KindVersion) SchemaJSON() (string, error) {
	if kv.schemaJSON != "" {
		return kv.schemaJSON, nil
	}
	if kv.schema == nil {
		return "", nil
	}
	jsonb, err := kv.schema.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf(
			"failed to marshal JSON for schema for %q: %w",
			kv.Name(), err,
		)
	}
	kv.schemaJSON = string(jsonb)
	return kv.schemaJSON, nil
}
