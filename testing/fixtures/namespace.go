package fixtures

import "github.com/relexec/rxp/namespace"

const (
	NamespaceUUID = "6ec8833a-c9d6-4bc6-ac90-73db8540ef13"
	NamespaceName = "ns1"
)

var (
	Namespace = namespace.New(
		namespace.WithDomain(Domain),
		namespace.WithUUID(NamespaceUUID),
		namespace.WithName(NamespaceName),
	)
)
