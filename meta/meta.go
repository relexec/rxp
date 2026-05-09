package meta

import (
	"fmt"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathKindVersion = fieldpath.FromString("kindversion")
	FieldPathNamescope   = fieldpath.FromString("namescope")
	FieldPathSchema      = fieldpath.FromString("schema")
)

type Meta struct {
	// system contains the System containing the Meta.
	system types.System
	// kindVersion is the [types.KindVersion] that uniquely identifies the type
	// and version of Objects represented by this Meta.
	kindVersion types.KindVersion
	// namescope is the uniqueness constraint of the Named.Name of
	// Objects having this Kind+Version.
	namescope types.Namescope
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
	kv := m.kindVersion
	err := kv.Validate()
	if err != nil {
		return err
	}
	err = m.namescope.Validate()
	if err != nil {
		return err
	}
	if m.schema == nil {
		return errors.MetaMissingSchema(kv)
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
	return m.kindVersion
}

// SetKindVersion sets the KindVersion of Meta.
func (m *Meta) SetKindVersion(kv types.KindVersion) {
	m.kindVersion = kv
}

// Namescope returns the name uniqueness constraint for Objects having this
// KindVersion.
func (m Meta) Namescope() types.Namescope {
	return m.namescope
}

// SetNamescope sets the name uniqueness constraint for Objects having this
// KindVersion.
func (m *Meta) SetNamescope(namescope types.Namescope) {
	m.namescope = namescope
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
			"failed to marshal JSON for schema for %q: %w", m.kindVersion, err,
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

	thisKV := string(m.kindVersion)
	otherKV := string(other.KindVersion())
	if thisKV != otherKV {
		d.Push(
			cmp.NewDifference(
				FieldPathKindVersion,
				cmp.DifferenceTypeModify,
				thisKV,
				otherKV,
			),
		)
	}
	if m.namescope != other.Namescope() {
		d.Push(
			cmp.NewDifference(
				FieldPathNamescope,
				cmp.DifferenceTypeModify,
				m.namescope,
				other.Namescope(),
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
					FieldPathNamescope,
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
			FieldPathKindVersion,
			cmp.DifferenceTypeAdd,
			string(m.kindVersion),
			nil,
		),
	)
	d.Push(
		cmp.NewDifference(
			FieldPathNamescope,
			cmp.DifferenceTypeAdd,
			m.namescope,
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
