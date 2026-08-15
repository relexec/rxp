package api

// Record is a base struct that describes some piece of data managed by rxp.
type Record struct {
	// systemInternalID contains the internal rxp backend implementation's
	// identifier for the Record. For backends that use relational databases
	// for persistence, this will most likely be an integer type that matches a
	// SERIAL or ROWID in the underlying database table row that stores the
	// record of the Record.
	systemInternalID any
}

// SetSystemInternalID sets the Record's internal system identifier.
func (r *Record) SetSystemInternalID(id any) {
	r.systemInternalID = id
}

// HasSystemInternalID returns whether an rxp backend set a system internal
// identifier on this record.
func (r Record) HasSystemInternalID() bool {
	return r.systemInternalID != nil
}

// SystemInternalID returns the internal rxp backend implementation's
// identifier for thir Record. For backends that use relational databases for
// persistence, this will most likely be an integer type that matches a SERIAL
// or ROWID in the underlying database table row that stores the record of this
// System.
func (r Record) SystemInternalID() any {
	return r.systemInternalID
}

// SystemInternalIDInt64 returns the int64 representation of the internal
// system identifier, or -1 if there is no internal system identifier or the
// underlying interface type was not an int64.
func (r Record) SystemInternalIDInt64() int64 {
	if r.systemInternalID == nil {
		return -1
	}
	v, ok := r.systemInternalID.(int64)
	if !ok {
		return -1
	}
	return v
}

// TreeRecord is a base struct for data managed by rxp that is part of a tree
// structure.
type TreeRecord struct {
	Record
	// nestedSetLeft is the nested set model's left value for this node in the
	// tree.
	nestedSetLeft int64
	// nestedSetRight is the nested set model's right value for this node in
	// the tree.
	nestedSetRight int64
}

// SetNestedSet sets the TreeNode's nested set left and right values.
func (r *TreeRecord) SetNestedSet(left, right int64) {
	r.nestedSetLeft = left
	r.nestedSetRight = right
}

// NestedSetLeft returns the left value of the tree node's nested set.
func (r TreeRecord) NestedSetLeft() int64 {
	return r.nestedSetLeft
}

// NestedSetRight returns the right value of the tree node's nested set.
func (r TreeRecord) NestedSetRight() int64 {
	return r.nestedSetRight
}
