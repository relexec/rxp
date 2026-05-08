package fixtures

import (
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/object"
	"github.com/relexec/rxp/types"
)

var (
	InvalidKind        = types.Kind("invalid kind")
	InvalidKindVersion = types.KindVersion("invalid kind version")
)

var (
	InvalidMeta = meta.New(
		meta.WithKindVersion(InvalidKindVersion),
		meta.WithNamescope(types.NamescopeKind),
	)
	InvalidObject = object.New(object.WithKindVersion(InvalidKindVersion))
)
