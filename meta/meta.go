package meta

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathKind    = fieldpath.FromString("kind")
	FieldPathVersion = fieldpath.FromString("version")
	FieldPathSchema  = fieldpath.FromString("schema")
)

type Meta struct {
	// system contains the System containing the Meta.
	system types.System
	// kind is the [types.Kind] that identifies the type of Objects represented
	// by this Meta.
	kind types.Kind
	// version is the [semver.Version] that identifies the specific version of
	// the Kind of Objects represented by this Meta.
	version semver.Version
	// schema is the [jsonschema.Schema] that describes the Spec field
	// composition of Object with this Kind+Version.
	schema types.Schema
	// schemaJSON stores a cache of the marshaled JSON for the
	// [jsonschema.Schema] that describes the Spec field composition of the
	// Objects with this Kind+Version.
	schemaJSON string
}

// Validate returns an error if the Meta is not valid.
func (m Meta) Validate() error {
	k := m.kind
	if k == nil {
		return errors.MetaMissingKind()
	}
	err := k.Validate()
	if err != nil {
		return err
	}
	if m.schema == nil {
		return errors.MetaMissingSchema(m.KindVersion())
	}
	return nil
}

// System returns the System of the Meta.
func (m Meta) System() types.System {
	return m.system
}

// SetSystem sets the System of Meta.
func (m *Meta) SetSystem(system types.System) {
	m.system = system
}

// KindVersion returns the KindVersion of the Meta.
func (m Meta) KindVersion() types.KindVersion {
	return types.NewKindVersion(m.kind.Name(), m.version)
}

// Kind returns the Kind of the Meta.
func (m Meta) Kind() types.Kind {
	return m.kind
}

// SetKind sets the Kind of the Meta.
func (m *Meta) SetKind(k types.Kind) {
	m.kind = k
}

// Version returns the Version of the Meta.
func (m Meta) Version() semver.Version {
	return m.version
}

// SetKind sets the Version of the Meta.
func (m *Meta) SetVersion(ver semver.Version) {
	m.version = ver
}

// Schema returns a [jsonschema.Schema] that describes the desired state fields
// of Objects with this KindVersion.
func (m Meta) Schema() types.Schema {
	return m.schema
}

// SetSchema sets the [jsonschema.Schema] that describes the desired state
// fields of Objects with this KindVersion.
func (m *Meta) SetSchema(schema types.Schema) {
	m.schema = schema
}

// SchemaJSON returns a string containing the [jsonschema.Schema] that
// describes the desired state fields of Objects with this KindVersion.
func (m *Meta) SchemaJSON() (string, error) {
	if m.schemaJSON != "" {
		return m.schemaJSON, nil
	}
	if m.schema == nil {
		return "", nil
	}
	jsonb, err := m.schema.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf(
			"failed to marshal JSON for schema for %q: %w",
			m.KindVersion(), err,
		)
	}
	m.schemaJSON = string(jsonb)
	return m.schemaJSON, nil
}

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (m Meta) Diff(subject any) (*cmp.Delta, error) {
	var other types.Meta
	switch subject := subject.(type) {
	case cmp.ZeroGen:
		return m.diffNew()
	case Meta:
		other = &subject
	case *Meta:
		other = subject
	default:
		return nil, cmp.CannotCompareTypes(m, subject)
	}

	d := &cmp.Delta{}

	thisKind := string(m.kind.Name())
	otherKind := string(other.Kind().Name())
	if thisKind != otherKind {
		d.Push(
			cmp.NewDifference(
				FieldPathKind,
				cmp.DifferenceTypeModify,
				thisKind,
				otherKind,
			),
		)
	}
	thisVersion := m.version.String()
	otherVersion := other.Version().String()
	if thisVersion != otherVersion {
		d.Push(
			cmp.NewDifference(
				FieldPathVersion,
				cmp.DifferenceTypeModify,
				thisVersion,
				otherVersion,
			),
		)
	}
	if m.schema != nil {
		thisSchemaBytes, err := m.schema.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("failed marshaling JSONSchema: %w", err)
		}
		thisSchemaJSON := string(thisSchemaBytes)
		if other.Schema() == nil {
			d.Push(
				cmp.NewDifference(
					FieldPathSchema,
					cmp.DifferenceTypeRemove,
					thisSchemaJSON,
					nil,
				),
			)
		} else {
			otherSchemaBytes, err := other.Schema().MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("failed marshaling JSONSchema: %w", err)
			}
			otherSchemaJSON := string(otherSchemaBytes)
			if thisSchemaJSON != otherSchemaJSON {
				d.Push(
					cmp.NewDifference(
						FieldPathSchema,
						cmp.DifferenceTypeModify,
						thisSchemaJSON,
						otherSchemaJSON,
					),
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
				cmp.NewDifference(
					FieldPathSchema,
					cmp.DifferenceTypeAdd,
					nil,
					otherSchemaJSON,
				),
			)
		}
	}
	return d, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the Meta as a
// new Meta (i.e. for the first generation)
func (m Meta) diffNew() (*cmp.Delta, error) {
	schemaJSON, err := m.SchemaJSON()
	if err != nil {
		return nil, err
	}
	d := &cmp.Delta{}

	d.Push(
		cmp.NewDifference(
			FieldPathKind,
			cmp.DifferenceTypeAdd,
			string(m.kind.Name()),
			nil,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathVersion,
			cmp.DifferenceTypeAdd,
			m.version.String(),
			nil,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathSchema,
			cmp.DifferenceTypeAdd,
			schemaJSON,
			nil,
		),
	)
	return d, nil
}

var _ types.Meta = (*Meta)(nil)
