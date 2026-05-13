package fixtures

import (
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/types"
)

var (
	InvalidDomainName    = types.DomainName("invalid domain")
	InvalidNamespaceName = types.NamespaceName("invalid ns")
	InvalidKindName      = types.KindName("invalid kind")
	InvalidKindVersion   = types.KindVersion("invalid kind version")
)

var (
	InvalidKind = kind.New(
		kind.WithName(InvalidKindName),
	)
	InvalidDomain = domain.New(
		domain.WithName(InvalidDomainName),
	)
	InvalidNamespace = namespace.New(
		namespace.WithDomain(Domain),
		namespace.WithName(InvalidNamespaceName),
	)
	InvalidMeta = meta.New(
		meta.WithKind(InvalidKind),
		meta.WithVersion(*SemVer_V1_0_0),
	)
	InvalidObject = object.New(object.WithKindVersion(InvalidKindVersion))
)
