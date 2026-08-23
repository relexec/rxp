package apicore

type Type string

const (
	TypeSystem      Type = "system"
	TypeDomain      Type = "domain"
	TypeKind        Type = "kind"
	TypeKindVersion Type = "kindversion"
	TypeObject      Type = "object"
	TypeRun         Type = "run"
	TypeRunEvent    Type = "runevent"
)
