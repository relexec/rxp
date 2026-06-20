package fixtures

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/system"
)

const (
	UnknownSystemUUID = "8ccf20f6-df19-45e1-9086-8eff1283fef5"
	UnknownSystemTag  = "unknown system"
	UnknownDomainUUID = "0a081c96-3f30-41aa-b635-30501fe5ef2e"
	UnknownDomainName = api.DomainName("unknown.domain.testing.dxp")
	UnknownKindName   = api.KindName("unknown.testing.rxp")
	UnknownKindUUID   = "2a6fc010-7c4b-4aef-bff4-e4ce98ad15df"
	UnknownObjectUUID = "2f959fc4-e885-4946-ae7b-dc015c185a62"
)

var (
	UnknownKindVersionName = api.NewKindVersionName(UnknownKindName, *SemVer_V1_0_0)
)

var (
	UnknownKind = kind.New(
		kind.WithUUID(UnknownKindUUID),
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
	UnknownKindVersion = kindversion.New(
		kindversion.WithKind(UnknownKind),
		kindversion.WithVersion(*SemVer_V1_0_0),
	)
	UnknownObject = object.New(
		object.WithKindVersionName(UnknownKindVersionName),
		object.WithUUID(UnknownObjectUUID),
		object.WithName("unknown"),
	)
)
