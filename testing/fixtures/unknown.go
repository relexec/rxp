package fixtures

import (
	"github.com/google/uuid"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/system"
	"github.com/relexec/rxp/types"
)

const (
	UnknownSystemUUID = "8ccf20f6-df19-45e1-9086-8eff1283fef5"
	UnknownSystemName = "unknown system"
	UnknownDomainName = types.DomainName("unknown.domain.testing.dxp")
	UnknownKind       = types.Kind("unknown.testing.rxp")
)

var (
	UnknownKindVersion = types.KindVersion(string(UnknownKind) + "@" + SemVer_V1_0_0.String())
)

var (
	UnknownSystem = system.New(
		system.WithUUID(UnknownSystemUUID),
		system.WithName(UnknownSystemName),
	)
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
