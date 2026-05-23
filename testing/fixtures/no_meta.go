package fixtures

import (
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/meta/schema"
)

var (
	NoMetaKindName = api.KindName("nometa.testing.rxp")
)

var (
	// We create this Kind during testing but never create any Metas
	// (KindVersions) with it. This allows us to check error responses
	// attempting to create a KindVersion with a non-0 version of this Kind.
	NoMetaKind = kind.New(
		kind.WithName(NoMetaKindName),
	)
	NoMetaMeta = meta.New(
		meta.WithKind(NoMetaKind),
		meta.WithVersion(*SemVer_V1_0_1),
		meta.WithSchema(&schema.Schema{}),
	)
)
