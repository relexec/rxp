package types

import "encoding/json"

// Schema describes the field composition for an Object's desired state.
type Schema interface {
	json.Marshaler
	json.Unmarshaler
	Differ
}
