package fixtures

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/object"
)

var (
	InvalidDomainName    = api.DomainName("invalid domain")
	InvalidNamespaceName = api.NamespaceName("invalid ns")
	InvalidKindName      = api.KindName("invalid kind")
	InvalidKindVersion   = api.KindVersion("invalid kind version")
)

var (
	InvalidKind = kind.New(
		kind.WithName(InvalidKindName),
	)
	InvalidDomain = domain.New(
		domain.WithUUID(DomainUUID),
		domain.WithName(InvalidDomainName),
	)
	InvalidNamespace = namespace.New(
		namespace.WithDomain(Domain),
		namespace.WithUUID(NamespaceUUID),
		namespace.WithName(InvalidNamespaceName),
	)
	InvalidMeta = meta.New(
		meta.WithKind(InvalidKind),
		meta.WithVersion(*SemVer_V1_0_0),
	)
	InvalidObject = object.New(object.WithKindVersion(InvalidKindVersion))
)
