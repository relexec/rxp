package schema

import (
	"slices"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/samber/lo"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/types"
)

var (
	FieldPathID     = fieldpath.FromString("$id")
	FieldPathSchema = fieldpath.FromString("$schema")

	FieldPathTitle       = fieldpath.FromString("title")
	FieldPathDescription = fieldpath.FromString("description")

	FieldPathProperties = fieldpath.FromString("properties")
)

// Schema wraps a [jsonschema.Schema] and implements the [types.Differ]
// interface.
type Schema struct {
	jsonschema.Schema
}

// Diff returns a [cmp.Delta] representing the difference between itself and
// something else of the same type.
//
// If the argument is the [cmp.ZeroGen] sentinel, the returned [cmp.Delta]
// represents instructions to create the thing.
func (s Schema) Diff(subject any) (*cmp.Delta, error) {
	var other Schema
	switch subject := subject.(type) {
	case cmp.ZeroGen:
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
		return nil, cmp.CannotCompareTypes(s, subject)
	}

	d := &cmp.Delta{}

	thisJS := s.Schema
	otherJS := other.Schema

	if thisJS.ID != otherJS.ID {
		d.Push(
			cmp.NewDifference(
				FieldPathID,
				cmp.DifferenceTypeModify,
				thisJS.ID,
				otherJS.ID,
			),
		)
	}
	if thisJS.Schema != otherJS.Schema {
		d.Push(
			cmp.NewDifference(
				FieldPathSchema,
				cmp.DifferenceTypeModify,
				thisJS.Schema,
				otherJS.Schema,
			),
		)
	}
	// TODO(jaypipes): Handle $defs
	// TODO(jaypipes): Handle Definitions
	// TODO(jaypipes): Handle DependencySchemas
	// TODO(jaypipes): Handle DependencyStrings

	if thisJS.Title != otherJS.Title {
		d.Push(
			cmp.NewDifference(
				FieldPathTitle,
				cmp.DifferenceTypeModify,
				thisJS.Title,
				otherJS.Title,
			),
		)
	}
	if thisJS.Description != otherJS.Description {
		d.Push(
			cmp.NewDifference(
				FieldPathDescription,
				cmp.DifferenceTypeModify,
				thisJS.Description,
				otherJS.Description,
			),
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
			cmp.NewDifference(
				FieldPathProperties,
				cmp.DifferenceTypeModify,
				thisProps,
				otherProps,
			),
		)
	}

	return d, nil
}

// diffNew returns a [cmp.Delta] containing instructions to make the Schema as
// a new Schema (i.e. for the first generation)
func (s Schema) diffNew() (*cmp.Delta, error) {
	d := &cmp.Delta{}
	if s.Schema.ID != "" {
		d.Push(
			cmp.NewDifference(
				FieldPathID,
				cmp.DifferenceTypeAdd,
				s.Schema.ID,
				nil,
			),
		)
	}
	if s.Schema.Schema != "" {
		d.Push(
			cmp.NewDifference(
				FieldPathSchema,
				cmp.DifferenceTypeAdd,
				s.Schema.Schema,
				nil,
			),
		)
	}
	if s.Schema.Title != "" {
		d.Push(
			cmp.NewDifference(
				FieldPathTitle,
				cmp.DifferenceTypeAdd,
				s.Schema.Title,
				nil,
			),
		)
	}
	if s.Schema.Description != "" {
		d.Push(
			cmp.NewDifference(
				FieldPathDescription,
				cmp.DifferenceTypeAdd,
				s.Schema.Description,
				nil,
			),
		)
	}
	d.Push(
		cmp.NewDifference(
			FieldPathProperties,
			cmp.DifferenceTypeAdd,
			s.Schema.Properties,
			nil,
		),
	)
	return d, nil
}

var _ types.Differ = (*Schema)(nil)
