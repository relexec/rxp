package fixtures

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/object"
)

var (
	InvalidDomainName      = api.DomainName("invalid domain")
	InvalidKindName        = api.KindName("invalid kind")
	InvalidKindUUID        = "6b23fab3-3a19-4885-92b7-0138871c4e62"
	InvalidKindVersionName = api.KindVersionName("invalid kind version")
)

var (
	InvalidKind = kind.New(
		kind.WithUUID(InvalidKindUUID),
		kind.WithName(InvalidKindName),
	)
	InvalidDomain = domain.New(
		domain.WithUUID(DomainUUID),
		domain.WithName(InvalidDomainName),
	)
	InvalidKindVersion = kindversion.New(
		kindversion.WithKind(InvalidKind),
		kindversion.WithVersion(*SemVer_V1_0_0),
	)
	InvalidObject = object.New(
		object.WithKindVersionName(InvalidKindVersionName),
	)
)
