package fixtures

import (
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/types"
)

var (
	InvalidDomainName  = types.DomainName("invalid domain")
	InvalidKind        = types.Kind("invalid kind")
	InvalidKindVersion = types.KindVersion("invalid kind version")
)

var (
	InvalidDomain = domain.New(
		domain.WithName(InvalidDomainName),
	)
	InvalidMeta = meta.New(
		meta.WithKindVersion(InvalidKindVersion),
		meta.WithNamescope(types.NamescopeKind),
	)
	InvalidObject = object.New(object.WithKindVersion(InvalidKindVersion))
)
