package fixtures

import (
	"github.com/google/uuid"
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/namespace"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/system"
)

const (
	UnknownSystemUUID    = "8ccf20f6-df19-45e1-9086-8eff1283fef5"
	UnknownSystemTag     = "unknown system"
	UnknownDomainName    = api.DomainName("unknown.domain.testing.dxp")
	UnknownNamespaceName = api.NamespaceName("unknown.ns")
	UnknownKindName      = api.KindName("unknown.testing.rxp")
)

var (
	UnknownKindVersion = api.NewKindVersion(UnknownKindName, *SemVer_V1_0_0)
)

var (
	UnknownKind = kind.New(
		kind.WithName(UnknownKindName),
	)
	UnknownSystem = system.New(
		system.WithUUID(UnknownSystemUUID),
		system.WithTag(UnknownSystemTag),
	)
	UnknownDomain = domain.New(
		domain.WithName(UnknownDomainName),
	)
	UnknownNamespace = namespace.New(
		namespace.WithDomain(Domain),
		namespace.WithName(UnknownNamespaceName),
	)
	UnknownMeta = meta.New(
		meta.WithKind(UnknownKind),
		meta.WithVersion(*SemVer_V1_0_0),
	)
	UnknownObject = object.New(
		object.WithKindVersion(UnknownKindVersion),
		object.WithUUID(uuid.NewString()),
		object.WithName("unknown"),
	)
)
