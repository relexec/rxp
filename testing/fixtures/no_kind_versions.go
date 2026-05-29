package fixtures

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/kind/kindversion/schema"
)

var (
	NoKindVersionsKindUUID = "7e76e553-f11e-42fd-831d-82819839abda"
	NoKindVersionsKindName = api.KindName("nokindversions.testing.rxp")
)

var (
	// We create this Kind during testing but never create any KindVersions
	// with it. This allows us to check error responses attempting to create a
	// KindVersion with a non-0 version of this Kind.
	NoKindVersionsKind = kind.New(
		kind.WithUUID(NoKindVersionsKindUUID),
		kind.WithName(NoKindVersionsKindName),
	)
	NoKindVersionsKindVersion = kindversion.New(
		kindversion.WithKind(NoKindVersionsKind),
		kindversion.WithVersion(*SemVer_V1_0_1),
		kindversion.WithSchema(&schema.Schema{}),
	)
)
