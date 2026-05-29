package api

type Type string

const (
	TypeSystem      Type = "system"
	TypeDomain      Type = "domain"
	TypeNamespace   Type = "namespace"
	TypeKind        Type = "kind"
	TypeKindVersion Type = "kindversion"
	TypeObject      Type = "object"
)
