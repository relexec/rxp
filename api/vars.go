package api

// Vars is used as a generic container for input and output field variables.
//
// NOTE(jaypipes): This struct is not thread-safe.
type Vars map[string]any

// Set sets the supplied value at the supplied key, ensuring that if Vars is
// nil, we allocate a new map.
func (v *Vars) Set(key string, value any) {
	var m map[string]any
	if v == nil {
		m = map[string]any{}
	} else {
		m = *v
	}
	m[key] = value
	*v = m
}
