package api

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/kind/kindversion/schema"
)

type KindVersion struct {
	// System contains the System containing the KindVersion. This is a pointer
	// to a System to allow for backends to default missing System information
	// to their host system.
	System *System `json:"system,omitempty"`
	// Kind is the [api.Kind] that identifies the type of Objects represented
	// by this KindVersion.
	Kind Kind `json:"kind"`
	// version is the [semver.Version] that identifies the specific version of
	// the Kind of Objects represented by this KindVersion.
	Version semver.Version `json:"version"`
	// schema is the [jsonschema.Schema] that describes the Spec field
	// composition of Object with this Kind+Version.
	Schema *schema.Schema `json:"schema,omitempty"`
	// schemaJSON stores a cache of the marshaled JSON for the
	// [jsonschema.Schema] that describes the Spec field composition of the
	// Objects with this Kind+Version.
	schemaJSON string
}

// Validate returns an error if the KindVersion is not valid.
func (kv KindVersion) Validate() error {
	err := kv.Kind.Validate()
	if err != nil {
		return err
	}
	if kv.Schema == nil {
		return errors.KindVersionMissingSchema(kv.Name())
	}
	return nil
}

// Name returns the KindVersionName of the KindVersion.
func (kv KindVersion) Name() KindVersionName {
	return NewKindVersionName(kv.Kind.Name, kv.Version)
}

// SchemaJSON returns a string containing the [jsonschema.Schema] that
// describes the desired state fields of Objects with this KindVersion.
func (kv *KindVersion) SchemaJSON() (string, error) {
	if kv.schemaJSON != "" {
		return kv.schemaJSON, nil
	}
	if kv.Schema == nil {
		return "", nil
	}
	jsonb, err := kv.Schema.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf(
			"failed to marshal JSON for schema for %q: %w",
			kv.Name(), err,
		)
	}
	kv.schemaJSON = string(jsonb)
	return kv.schemaJSON, nil
}
