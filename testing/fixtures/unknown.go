package fixtures

import (
	"github.com/google/uuid"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/types"
)

const (
	UnknownKind = types.Kind("unknown.testing.rxp")
)

var (
	UnknownKindVersion = types.KindVersion(string(UnknownKind) + "@" + SemVer_V1_0_0.String())
)

var (
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
