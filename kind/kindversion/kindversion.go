package kindversion

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion/schema"
	"github.com/relexec/rxp/system"
)

var (
	FieldPathKind    = fieldpath.FromString("kind")
	FieldPathVersion = fieldpath.FromString("version")
	FieldPathSchema  = fieldpath.FromString("schema")
)

type KindVersion struct {
	// system contains the System containing the KindVersion.
	system *system.System
	// kind is the [kind.Kind] that identifies the type of Objects represented
	// by this KindVersion.
	kind *kind.Kind
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
func (kv KindVersion) System() *system.System {
	return kv.system
}

// SetSystem sets the System of KindVersion.
func (kv *KindVersion) SetSystem(system *system.System) {
	kv.system = system
}

// Name returns the KindVersionName of the KindVersion.
func (kv KindVersion) Name() api.KindVersionName {
	return api.NewKindVersionName(kv.kind.Name(), kv.version)
}

// Kind returns the Kind of the KindVersion.
func (kv KindVersion) Kind() *kind.Kind {
	return kv.kind
}

// SetKind sets the Kind of the KindVersion.
func (kv *KindVersion) SetKind(k *kind.Kind) {
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

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func (kv KindVersion) Diff(subject any) (*delta.Delta, error) {
	var other *KindVersion
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return kv.diffNew()
	case KindVersion:
		other = &subject
	case *KindVersion:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(kv, subject)
	}

	d := &delta.Delta{}

	thisKind := string(kv.kind.Name())
	otherKind := string(other.Kind().Name())
	if thisKind != otherKind {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathKind,
				Type:      delta.DifferenceTypeModify,
				From:      thisKind,
				To:        otherKind,
			},
		)
	}
	thisVersion := kv.version.String()
	otherVersion := other.Version().String()
	if thisVersion != otherVersion {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathVersion,
				Type:      delta.DifferenceTypeModify,
				From:      thisVersion,
				To:        otherVersion,
			},
		)
	}
	if kv.schema != nil {
		thisSchemaBytes, err := kv.schema.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("failed marshaling JSONSchema: %w", err)
		}
		thisSchemaJSON := string(thisSchemaBytes)
		if other.Schema() == nil {
			d.Push(
				delta.Difference{
					FieldPath: FieldPathSchema,
					Type:      delta.DifferenceTypeRemove,
					From:      thisSchemaJSON,
					To:        nil,
				},
			)
		} else {
			otherSchemaBytes, err := other.Schema().MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("failed marshaling JSONSchema: %w", err)
			}
			otherSchemaJSON := string(otherSchemaBytes)
			if thisSchemaJSON != otherSchemaJSON {
				d.Push(
					delta.Difference{
						FieldPath: FieldPathSchema,
						Type:      delta.DifferenceTypeModify,
						From:      thisSchemaJSON,
						To:        otherSchemaJSON,
					},
				)
			}
		}
	} else {
		if other.Schema() != nil {
			otherSchemaBytes, err := other.Schema().MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("failed marshaling JSONSchema: %w", err)
			}
			otherSchemaJSON := string(otherSchemaBytes)
			d.Push(
				delta.Difference{
					FieldPath: FieldPathSchema,
					Type:      delta.DifferenceTypeAdd,
					From:      nil,
					To:        otherSchemaJSON,
				},
			)
		}
	}
	return d, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the KindVersion as a
// new KindVersion (i.e. for the first generation)
func (kv KindVersion) diffNew() (*delta.Delta, error) {
	schemaJSON, err := kv.SchemaJSON()
	if err != nil {
		return nil, err
	}
	d := &delta.Delta{}

	d.Push(
		delta.Difference{
			FieldPath: FieldPathKind,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        string(kv.kind.Name()),
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathVersion,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        kv.version.String(),
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathSchema,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        schemaJSON,
		},
	)
	return d, nil
}
