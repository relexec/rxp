package fixtures

import (
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/types"
)

var (
	InvalidDomainName    = types.DomainName("invalid domain")
	InvalidNamespaceName = types.NamespaceName("invalid ns")
	InvalidKind          = types.Kind("invalid kind")
	InvalidKindVersion   = types.KindVersion("invalid kind version")
)

var (
	InvalidDomain = domain.New(
		domain.WithName(InvalidDomainName),
	)
	InvalidNamespace = namespace.New(
		namespace.WithDomain(Domain),
		namespace.WithName(InvalidNamespaceName),
	)
	InvalidMeta = meta.New(
		meta.WithKindVersion(InvalidKindVersion),
		meta.WithNamescope(types.NamescopeKind),
	)
	InvalidObject = object.New(object.WithKindVersion(InvalidKindVersion))
)
