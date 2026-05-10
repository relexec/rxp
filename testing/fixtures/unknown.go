package fixtures

import (
	"github.com/google/uuid"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/types"
)

const (
	UnknownDomainName = types.DomainName("unknown.domain.testing.dxp")
	UnknownKind       = types.Kind("unknown.testing.rxp")
)

var (
	UnknownKindVersion = types.KindVersion(string(UnknownKind) + "@" + SemVer_V1_0_0.String())
)

var (
	UnknownDomain = domain.New(
		domain.WithName(UnknownDomainName),
	)
	UnknownMeta = meta.New(
		meta.WithKindVersion(UnknownKindVersion),
		meta.WithNamescope(types.NamescopeKind),
	)
	UnknownObject = object.New(
		object.WithKindVersion(UnknownKindVersion),
		object.WithUUID(uuid.NewString()),
		object.WithName("unknown"),
	)
)
