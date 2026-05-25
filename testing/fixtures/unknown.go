package fixtures

import (
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
	UnknownDomainUUID    = "0a081c96-3f30-41aa-b635-30501fe5ef2e"
	UnknownDomainName    = api.DomainName("unknown.domain.testing.dxp")
	UnknownNamespaceUUID = "8c12b8d9-eada-4c57-aa93-fc2984aea2c0"
	UnknownNamespaceName = api.NamespaceName("unknown.ns")
	UnknownKindName      = api.KindName("unknown.testing.rxp")
	UnknownObjectUUID    = "2f959fc4-e885-4946-ae7b-dc015c185a62"
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
		domain.WithUUID(UnknownDomainUUID),
		domain.WithName(UnknownDomainName),
	)
	UnknownNamespace = namespace.New(
		namespace.WithDomain(Domain),
		namespace.WithUUID(UnknownNamespaceUUID),
		namespace.WithName(UnknownNamespaceName),
	)
	UnknownMeta = meta.New(
		meta.WithKind(UnknownKind),
		meta.WithVersion(*SemVer_V1_0_0),
	)
	UnknownObject = object.New(
		object.WithKindVersion(UnknownKindVersion),
		object.WithUUID(UnknownObjectUUID),
		object.WithName("unknown"),
	)
)
