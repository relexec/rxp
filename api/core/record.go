package apicore

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
// identifier for the Record. For backends that use relational databases for
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
