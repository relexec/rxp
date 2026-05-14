package types

// Label is a key/value pair that may be attached to an Object.
//
// Objects can be selected and categorized using Labels.
type Label struct {
	// Key stores the Label's string key.
	Key string `json:"key"`
	// Value stores the Label's string value.
	Value string `json:"value,omitempty"`
}

// Labels is a collection of Label structs.
//
// Though this is an array, there is no inherent ordering to the list of
// Labels. The reason this is not a map[string]string is so that Selectors can
// select on just a label key or key prefix and not the whole key and value.
type Labels []Label
