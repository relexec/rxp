package v1

import (
	"log"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/jsonschema-go/jsonschema"

	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/meta/schema"
	"github.com/relexec/rxp/types"
)

const (
	Kind = types.Kind("book.testing.rxp")
)

const (
	Version_V1_0_0 = "v1.0.0"
	Version_V1_0_1 = "v1.0.1"
)

var (
	SemVer_V1_0_0 = semver.MustParse(Version_V1_0_0)
	SemVer_V1_0_1 = semver.MustParse(Version_V1_0_1)
)

var (
	KindVersion_V1_0_0 = types.KindVersion(string(Kind) + "@" + SemVer_V1_0_0.String())
	KindVersion_V1_0_1 = types.KindVersion(string(Kind) + "@" + SemVer_V1_0_1.String())
)

// LatestMeta returns the [types.Meta] representing the latest known version of
// the Book object meta.
func LatestMeta() types.Meta {
	return Meta_V1_0_1
}

// LatestVersion returns the [semver.Version] representing the latest known
// version of the Book object meta in the v1 major version series.
func LatestVersion() *semver.Version {
	return SemVer_V1_0_1
}

var (
	Schema_V1_0_0     *schema.Schema
	SchemaJSON_V1_0_0 string
	Schema_V1_0_1     *schema.Schema
	SchemaJSON_V1_0_1 string
)

var (
	Meta_V1_0_0 *meta.Meta
	Meta_V1_0_1 *meta.Meta
)

var (
	Metas = map[string]*meta.Meta{
		Version_V1_0_0: Meta_V1_0_0,
		Version_V1_0_1: Meta_V1_0_1,
	}
)

// Meta returns the Meta associated with the supplied version. Version can be
// either a string or a [semver.Version].
func Meta(version any) (*meta.Meta, bool) {
	var vstr string
	switch version := version.(type) {
	case string:
		vstr = version
	case *semver.Version:
		vstr = version.Original()
	}
	if !strings.HasPrefix(vstr, "v") {
		vstr = "v" + vstr
	}
	m, ok := Metas[vstr]
	return m, ok
}

func init() {
	var err error

	js, err := jsonschema.For[Spec_V1_0_0](nil)
	if err != nil {
		log.Fatalf(
			"failed to construct jsonschema.Schema for Book_V1_0_0: %s",
			err.Error(),
		)
	}
	jsonb, err := js.MarshalJSON()
	if err != nil {
		log.Fatalf(
			"failed to marshal JSON for schema for Book_V1_0_0: %s",
			err.Error(),
		)
	}
	Schema_V1_0_0 = &schema.Schema{Schema: *js}
	SchemaJSON_V1_0_0 = string(jsonb)
	Meta_V1_0_0 = meta.New(
		meta.WithKindVersion(KindVersion_V1_0_0),
		meta.WithNamescope(types.NamescopeKind),
		meta.WithSchema(Schema_V1_0_0),
		meta.WithSchemaJSON(SchemaJSON_V1_0_0),
	)

	js, err = jsonschema.For[Spec_V1_0_1](nil)
	if err != nil {
		log.Fatalf(
			"failed to construct jsonschema.Schema for Book_V1_0_1: %s",
			err.Error(),
		)
	}
	jsonb, err = js.MarshalJSON()
	if err != nil {
		log.Fatalf(
			"failed to marshal JSON for schema for Book_V1_0_1: %s",
			err.Error(),
		)
	}
	Schema_V1_0_1 = &schema.Schema{Schema: *js}
	SchemaJSON_V1_0_1 = string(jsonb)
	Meta_V1_0_1 = meta.New(
		meta.WithKindVersion(KindVersion_V1_0_1),
		meta.WithNamescope(types.NamescopeKind),
		meta.WithSchema(Schema_V1_0_1),
		meta.WithSchemaJSON(SchemaJSON_V1_0_1),
	)
}
