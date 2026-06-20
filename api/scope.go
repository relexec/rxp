package api

// Scope refers to the extent to which Names of instances of a Type of thing
// are unique.
type Scope int

const (
	// ScopeDomain means instances of a thing of this Type have a Name that
	// is unique within its Domain.
	ScopeDomain Scope = iota
	// ScopeSystem means instances of a thing of this Type have a Name that
	// is unique within its System.
	ScopeSystem
	// ScopeGlobal means instances of a thing of this Type can only be
	// identified by UUID and never by Name.
	ScopeGlobal
)

func (s Scope) String() string {
	switch s {
	case ScopeDomain:
		return "domain"
	case ScopeSystem:
		return "system"
	case ScopeGlobal:
		return "global"
	default:
		return "unknown"
	}
}
