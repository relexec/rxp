package schema

import (
	"slices"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/samber/lo"

	"github.com/relexec/delta"
	"github.com/relexec/delta/fieldpath"
)

var (
	FieldPathID     = fieldpath.FromString("$id")
	FieldPathSchema = fieldpath.FromString("$schema")

	FieldPathTitle       = fieldpath.FromString("title")
	FieldPathDescription = fieldpath.FromString("description")

	FieldPathProperties = fieldpath.FromString("properties")
)

// Diff returns a [delta.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [delta.ZeroGen] sentinel, the returned [delta.Delta]
// represents instructions to create the thing.
func (s Schema) Diff(subject any) (*delta.Delta, error) {
	var other Schema
	switch subject := subject.(type) {
	case delta.ZeroGen:
		return s.diffNew()
	case Schema:
		other = subject
	case *Schema:
		other = *subject
	case jsonschema.Schema:
		other = Schema{subject}
	case *jsonschema.Schema:
		other = Schema{*subject}
	default:
		return nil, delta.CannotCompareTypes(s, subject)
	}

	d := &delta.Delta{}

	thisJS := s.Schema
	otherJS := other.Schema

	if thisJS.ID != otherJS.ID {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathID,
				Type:      delta.DifferenceTypeModify,
				From:      thisJS.ID,
				To:        otherJS.ID,
			},
		)
	}
	if thisJS.Schema != otherJS.Schema {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathSchema,
				Type:      delta.DifferenceTypeModify,
				From:      thisJS.Schema,
				To:        otherJS.Schema,
			},
		)
	}
	// TODO(jaypipes): Handle $defs
	// TODO(jaypipes): Handle Definitions
	// TODO(jaypipes): Handle DependencySchemas
	// TODO(jaypipes): Handle DependencyStrings

	if thisJS.Title != otherJS.Title {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathTitle,
				Type:      delta.DifferenceTypeModify,
				From:      thisJS.Title,
				To:        otherJS.Title,
			},
		)
	}
	if thisJS.Description != otherJS.Description {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathDescription,
				Type:      delta.DifferenceTypeModify,
				From:      thisJS.Description,
				To:        otherJS.Description,
			},
		)
	}

	thisProps := thisJS.Properties
	otherProps := otherJS.Properties

	thisPropKeys := lo.Keys(thisProps)
	otherPropKeys := lo.Keys(otherProps)

	slices.Sort(thisPropKeys)
	slices.Sort(otherPropKeys)

	if len(thisPropKeys) != len(otherPropKeys) || !slices.Equal(thisPropKeys, otherPropKeys) {
		d.Push(
			delta.Difference{
				FieldPath: FieldPathProperties,
				Type:      delta.DifferenceTypeModify,
				From:      thisProps,
				To:        otherProps,
			},
		)
	}

	return d, nil
}

// diffNew returns a [delta.Delta] containing instructions to make the Schema as
// a new Schema (i.e. for the first generation)
func (s Schema) diffNew() (*delta.Delta, error) {
	d := &delta.Delta{}
	if s.Schema.ID != "" { //nolint:staticcheck
		d.Push(
			delta.Difference{
				FieldPath: FieldPathID,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        s.Schema.ID, //nolint:staticcheck
			},
		)
	}
	if s.Schema.Schema != "" { //nolint:staticcheck
		d.Push(
			delta.Difference{
				FieldPath: FieldPathSchema,
				Type:      delta.DifferenceTypeAdd,
				To:        s.Schema.Schema, //nolint:staticcheck
				From:      nil,
			},
		)
	}
	if s.Schema.Title != "" { //nolint:staticcheck
		d.Push(
			delta.Difference{
				FieldPath: FieldPathTitle,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        s.Schema.Title, //nolint:staticcheck
			},
		)
	}
	if s.Schema.Description != "" { //nolint:staticcheck
		d.Push(
			delta.Difference{
				FieldPath: FieldPathDescription,
				Type:      delta.DifferenceTypeAdd,
				From:      nil,
				To:        s.Schema.Description, //nolint:staticcheck
			},
		)
	}
	d.Push(
		delta.Difference{
			FieldPath: FieldPathProperties,
			Type:      delta.DifferenceTypeAdd,
			From:      nil,
			To:        s.Schema.Properties, //nolint:staticcheck
		},
	)
	return d, nil
}
