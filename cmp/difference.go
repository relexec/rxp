package cmp

import (
	"encoding/json"

	"github.com/relexec/rxp/cmp/fieldpath"
)

// DifferenceType differentiates kinds of Differences
type DifferenceType int

const (
	// DifferenceTypeAdd represents an addition of some field or item in a
	// list/map field value type.
	DifferenceTypeAdd DifferenceType = iota
	// DifferenceTypeRemove represents the removal of some field or item in a
	// list/map field value type.
	DifferenceTypeRemove
	// DifferenceTypeModify represents a modification of some field or item in
	// a list/map field value type.
	DifferenceTypeModify
	// DifferenceTypeOrder represents a change in the order of elements in a
	// sorted list field.
	DifferenceTypeOrder
)

// Difference represents a single difference between two things.
type Difference struct {
	// fp contains the FieldPath where this difference exists.
	fp fieldpath.FieldPath
	// typ contains the type of difference.
	typ DifferenceType
	// from contains the "A" or "from" object's value at the FieldPath.
	from any
	// to contains the "B" or "to" object's value at the FieldPath.
	to any
}

// FieldPath returns the dotted-notation field path, e.g. "spec" or
// "spec.generation", that indicates the specific field in the compared
// things that is different.
func (d Difference) FieldPath() fieldpath.FieldPath {
	return d.fp
}

// Type returns the type of difference.
func (d Difference) Type() DifferenceType {
	return d.typ
}

// FromValue returns the value of the field in the "A" or "from" thing at the
// FieldPath. When Type is DifferenceTypeAdd, returns nil.
func (d Difference) FromValue() any {
	return d.from
}

// ToValue returns the value of the field in the "B" or "to" thing at the
// FieldPath. When Type is DifferenceTypeRemove, returns nil.
func (d Difference) ToValue() any {
	return d.to
}

// MarshalJSON marshals the Delta into a JSON string representation that can
// be unmarshaled with UnmarshalJSON.
func (d Difference) MarshalJSON() ([]byte, error) {
	de := diffExported{
		Path: d.fp.String(),
		Type: int(d.typ),
		From: d.from,
		To:   d.to,
	}
	return json.Marshal(de)
}

// UnmarshalJSON constructs a Delta from a bytestring that was encoded with
// MarshalJSON.
func (d *Difference) UnmarshalJSON(text []byte) error {
	var de diffExported
	if err := json.Unmarshal(text, &de); err != nil {
		return err
	}
	d.fp = fieldpath.FromString(de.Path)
	d.typ = DifferenceType(de.Type)
	if de.From != nil {
		d.from = de.From

	}
	if de.To != nil {
		d.to = de.To
	}
	return nil
}

type diffExported struct {
	Path string `json:"path"`
	Type int    `json:"type"`
	From any    `json:"from,omitempty"`
	To   any    `json:"to,omitempty"`
}

// NewDifference returns a Difference given a DifferenceType, a FieldPath, a
// from and a to field value.
func NewDifference(
	fp fieldpath.FieldPath,
	typ DifferenceType,
	from any,
	to any,
) Difference {
	return Difference{
		fp:   fp,
		typ:  typ,
		from: from,
		to:   to,
	}
}
