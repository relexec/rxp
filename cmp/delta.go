package cmp

import (
	"encoding/json"
	"strings"

	"github.com/relexec/rxp/cmp/fieldpath"
)

// ZeroGen is a sentinel type used in creating Deltas for new things.
type ZeroGen struct{}

var (
	Zero = ZeroGen{}
)

// Delta represents the difference between two things of the same type.
type Delta struct {
	// diffs contains the set of the Differences between two things of the same
	// type.
	diffs []Difference
}

// Push adds a Difference to the Delta.
func (d *Delta) Push(diff Difference) {
	d.diffs = append(d.diffs, diff)
}

// Differences returns the set of differences between two things.
func (d Delta) Differences() []Difference {
	return d.diffs
}

// DifferenceAt returns the Differences in the Delta at the supplied path, or
// nil if there was no Difference at the supplied path.
func (d Delta) DifferenceAt(subject fieldpath.FieldPath) *Difference {
	if len(d.diffs) == 0 {
		return nil
	}
	for _, diff := range d.diffs {
		if diff.FieldPath().Contains(subject) {
			return &diff
		}
	}
	return nil
}

// DifferencesAtAny returns the Differences in the Delta at any of the supplied
// paths, or nil if there was no Difference at the supplied path.
func (d Delta) DifferencesAtAny(subjects ...fieldpath.FieldPath) []Difference {
	if len(d.diffs) == 0 {
		return nil
	}
	res := []Difference{}
	for _, diff := range d.diffs {
		if diff.FieldPath().ContainsAny(subjects) {
			res = append(res, diff)
		}
	}
	return res
}

// Different returns true if there are any differences in the Delta.
func (d Delta) Different() bool {
	return len(d.diffs) > 0
}

// DifferentAt returns whether there are differences at the supplied path in
// the things under comparison.
func (d Delta) DifferentAt(subject fieldpath.FieldPath) bool {
	if len(d.diffs) == 0 {
		return false
	}
	for _, diff := range d.diffs {
		if diff.FieldPath().Contains(subject) {
			return true
		}
	}
	return false
}

// DifferentAt returns whether there are differences at any of the the supplied
// paths in the things under comparison.
func (d Delta) DifferentAtAny(subjects ...fieldpath.FieldPath) bool {
	if len(d.diffs) == 0 {
		return false
	}
	for _, diff := range d.diffs {
		if diff.FieldPath().ContainsAny(subjects) {
			return true
		}
	}
	return false
}

// DifferentExcept returns true if there are differences in the things under
// comparison at any path *except* the supplied paths.
//
// This method is useful when comparing different objects where you only care
// about whether certain fields (or sets of fields) are different but not
// others. For example, assume some object with the following pseudo-schema:
//
// | {
// |   "spec": {
// |     "replicas": <int>,
// |     "generation": <int>
// |   },
// |   "status": {
// |     "readyReplicas": <int>,
// |     "lastHeartbeat": <timestamp>
// |   },
// | }
//
// If we wanted to know whether two objects with the above schema had different
// spec.replicas or spec.generation field values but don't care if the
// status.lastHeartbeat or status.readyReplicas field values are different, we
// could call `Delta.DifferentExcept("status") to determine that.
func (d Delta) DifferentExcept(
	exceptPaths ...fieldpath.FieldPath,
) bool {
	numDiffs := len(d.diffs)
	if numDiffs == 0 {
		return false
	} else if numDiffs > len(exceptPaths) {
		return true
	}
	foundExcepts := 0
	for _, diff := range d.diffs {
		for _, exceptPath := range exceptPaths {
			if diff.FieldPath().Contains(exceptPath) {
				foundExcepts++
			}
		}
	}
	return foundExcepts != numDiffs
}

// String returns a nicely-formatted string showing the differences contained
// in the Delta.
func (d Delta) String() string {
	if !d.Different() {
		return "<empty>"
	}
	// TODO
	b := strings.Builder{}
	return b.String()
}

// MarshalJSON marshals the Delta into a JSON string representation that can
// be unmarshaled with UnmarshalJSON.
func (d Delta) MarshalJSON() ([]byte, error) {
	if len(d.diffs) == 0 {
		return nil, nil
	}
	return json.Marshal(d.diffs)
}

// UnmarshalJSON constructs a Delta from a bytestring that was encoded with
// MarshalJSON.
func (d *Delta) UnmarshalJSON(text []byte) error {
	return json.Unmarshal(text, d)
}

// NewDelta returns a Delta given zero or more Differences.
func NewDelta(diffs ...Difference) Delta {
	return Delta{diffs}
}
