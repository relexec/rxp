package rxp

// Driver is the primary interface that rxp backends implement.
type Driver interface {
	ControlPlane
	DataPlane
}
