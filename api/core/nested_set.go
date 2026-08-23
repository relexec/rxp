package apicore

// NestedSetRecord is a base struct for data managed by rxp that is part of a
// tree structure that maintains a nested set left/right value for the tree
// node.
type NestedSetRecord struct {
	Record
	// nestedSetLeft is the nested set model's left value for this node in the
	// tree.
	nestedSetLeft int64
	// nestedSetRight is the nested set model's right value for this node in
	// the tree.
	nestedSetRight int64
}

// SetNestedSet sets the NestedSetRecord's nested set left and right values.
func (r *NestedSetRecord) SetNestedSet(left, right int64) {
	r.nestedSetLeft = left
	r.nestedSetRight = right
}

// NestedSetLeft returns the left value of the tree node's nested set.
func (r NestedSetRecord) NestedSetLeft() int64 {
	return r.nestedSetLeft
}

// NestedSetRight returns the right value of the tree node's nested set.
func (r NestedSetRecord) NestedSetRight() int64 {
	return r.nestedSetRight
}
