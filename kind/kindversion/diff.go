package kindversion

import (
	"fmt"

	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"

	"github.com/relexec/rxp/api"
)

var (
	FieldPathKind    = fieldpath.FromString("kind")
	FieldPathVersion = fieldpath.FromString("version")
	FieldPathSchema  = fieldpath.FromString("schema")
)

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func Diff(kv api.KindVersion, subject any) (*delta.Delta, error) {
	var other *api.KindVersion
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return diffNew(kv)
	case api.KindVersion:
		other = &subject
	case *api.KindVersion:
		other = subject
	default:
		return nil, delta.CannotCompareTypes(kv, subject)
	}

	d := &delta.Delta{}

	thisKind := string(kv.Kind().Name())
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
	thisVersion := kv.Version().String()
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
	if kv.Schema() != nil {
		thisSchemaBytes, err := kv.Schema().MarshalJSON()
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
func diffNew(kv api.KindVersion) (*delta.Delta, error) {
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
			To:        string(kv.Kind().Name()),
		},
	)
	d.Push(
		delta.Difference{
			FieldPath: FieldPathVersion,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        kv.Version().String(),
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
