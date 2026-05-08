package fixtures

import "github.com/relexec/rxp/namespace"

const (
	NamespaceName = "ns1"
)

var (
	Namespace = namespace.New(
		namespace.WithDomain(Domain),
		namespace.WithName(NamespaceName),
	)
)
