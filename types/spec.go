package types

// Spec describes the desired state of an Object.
type Spec interface {
	Differ
}
