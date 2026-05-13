package platform

import (
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/jsonschema-go/jsonschema"

	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/meta"
	"github.com/relexec/rxp/meta/schema"
	"github.com/relexec/rxp/types"
)

const (
	KindName  = types.KindName("platform.testing.rxp")
	Namescope = types.NamescopeSystem
)

var (
	Kind = kind.New(
		kind.WithName(KindName),
		kind.WithNamescope(Namescope),
	)
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
	Meta_V1_0_0       *meta.Meta
	Schema_V1_0_0     *schema.Schema
	SchemaJSON_V1_0_0 string

	Meta_V1_0_1       *meta.Meta
	Schema_V1_0_1     *schema.Schema
	SchemaJSON_V1_0_1 string
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

// FirstMeta returns the [types.Meta] representing the first known version of
// the meta.
func FirstMeta() types.Meta {
	return Meta_V1_0_0
}

// FirstVersion returns the [semver.Version] representing the first known
// version of the meta.
func FirstVersion() *semver.Version {
	return SemVer_V1_0_0
}

// FirstKindVersion returns the [types.KindVersion] representing the first
// known version of the Kind.
func FirstKindVersion() types.KindVersion {
	return types.NewKindVersion(KindName, *SemVer_V1_0_0)
}

var (
	// firstVersions is a map of major version number (without "v" prefix) to
	// the first known meta version in that major version series.
	firstVersions = map[string]*semver.Version{
		"1": FirstVersion(),
	}
)

// FirstVersionIn returns the [semver.Version] representing the first known
// version of the meta in the supplied major version series.
//
// If the supplied major version string is not known, returns an error.
func FirstVersionIn(major string) (*semver.Version, error) {
	v, ok := firstVersions[major]
	if !ok {
		return nil, fmt.Errorf("unknown major version %s", major)
	}
	return v, nil
}

// LastMeta returns the [types.Meta] representing the last known version of the
// meta.
func LastMeta() types.Meta {
	return Meta_V1_0_1
}

// LastVersion returns the [semver.Version] representing the last known version
// of the meta.
func LastVersion() *semver.Version {
	return SemVer_V1_0_1
}

// LastKindVersion returns the [types.KindVersion] representing the last known
// version of the Kind.
func LastKindVersion() types.KindVersion {
	return types.NewKindVersion(KindName, *SemVer_V1_0_1)
}

var (
	// lastVersions is a map of major version number (without "v" prefix) to
	// the last known meta version in that major version series.
	lastVersions = map[string]*semver.Version{
		"1": SemVer_V1_0_1,
	}
)

// LastVersionIn returns the [semver.Version] representing the last known
// version of the meta in the supplied major version series.
//
// If the supplied major version string is not known, returns an error.
func LastVersionIn(major string) (*semver.Version, error) {
	v, ok := lastVersions[major]
	if !ok {
		return nil, fmt.Errorf("unknown major version %s", major)
	}
	return v, nil
}

func init() {
	var err error

	js, err := jsonschema.For[Spec_V1_0_0](nil)
	if err != nil {
		log.Fatalf(
			"failed to construct jsonschema.Schema for Platform_V1_0_0: %s",
			err.Error(),
		)
	}
	jsonb, err := js.MarshalJSON()
	if err != nil {
		log.Fatalf(
			"failed to marshal JSON for schema for Platform_V1_0_0: %s",
			err.Error(),
		)
	}
	Schema_V1_0_0 = &schema.Schema{Schema: *js}
	SchemaJSON_V1_0_0 = string(jsonb)
	Meta_V1_0_0 = meta.New(
		meta.WithKind(Kind),
		meta.WithVersion(*SemVer_V1_0_0),
		meta.WithSchema(Schema_V1_0_0),
		meta.WithSchemaJSON(SchemaJSON_V1_0_0),
	)

	js, err = jsonschema.For[Spec_V1_0_1](nil)
	if err != nil {
		log.Fatalf(
			"failed to construct jsonschema.Schema for Platform_V1_0_1: %s",
			err.Error(),
		)
	}
	jsonb, err = js.MarshalJSON()
	if err != nil {
		log.Fatalf(
			"failed to marshal JSON for schema for Platform_V1_0_1: %s",
			err.Error(),
		)
	}
	Schema_V1_0_1 = &schema.Schema{Schema: *js}
	SchemaJSON_V1_0_1 = string(jsonb)
	Meta_V1_0_1 = meta.New(
		meta.WithKind(Kind),
		meta.WithVersion(*SemVer_V1_0_1),
		meta.WithSchema(Schema_V1_0_1),
		meta.WithSchemaJSON(SchemaJSON_V1_0_1),
	)
}
