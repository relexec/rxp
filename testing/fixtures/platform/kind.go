package platform

import (
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/jsonschema-go/jsonschema"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/kind/kindversion/schema"
)

const (
	KindName = api.KindName("platform.testing.rxp")
	KindUUID = "a9912ad6-8436-4085-9e58-64bdb8f06fc2"
	Scope    = api.ScopeSystem
)

var (
	Kind = kind.New(
		kind.WithUUID(KindUUID),
		kind.WithName(KindName),
		kind.WithScope(Scope),
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
	KindVersion_V1_0_0 *kindversion.KindVersion
	Schema_V1_0_0      *schema.Schema
	SchemaJSON_V1_0_0  string

	KindVersion_V1_0_1 *kindversion.KindVersion
	Schema_V1_0_1      *schema.Schema
	SchemaJSON_V1_0_1  string
)

var (
	KindVersions = map[string]*kindversion.KindVersion{
		Version_V1_0_0: KindVersion_V1_0_0,
		Version_V1_0_1: KindVersion_V1_0_1,
	}
)

// KindVersion returns the KindVersion associated with the supplied version.
// Version can be either a string or a [semver.Version].
func KindVersion(version any) (*kindversion.KindVersion, bool) {
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
	m, ok := KindVersions[vstr]
	return m, ok
}

// FirstKindVersion returns the [*kindversion.KindVersion] representing the
// first known version of the kindversion.
func FirstKindVersion() *kindversion.KindVersion {
	return KindVersion_V1_0_0
}

// FirstVersion returns the [semver.Version] representing the first known
// version of the kindversion.
func FirstVersion() *semver.Version {
	return SemVer_V1_0_0
}

// FirstKindVersionName returns the [api.KindVersionName] representing the
// first known version of the Kind.
func FirstKindVersionName() api.KindVersionName {
	return api.NewKindVersionName(KindName, *SemVer_V1_0_0)
}

var (
	// firstVersions is a map of major version number (without "v" prefix) to
	// the first known kindversion version in that major version series.
	firstVersions = map[string]*semver.Version{
		"1": FirstVersion(),
	}
)

// FirstVersionIn returns the [semver.Version] representing the first known
// version of the kindversion in the supplied major version series.
//
// If the supplied major version string is not known, returns an error.
func FirstVersionIn(major string) (*semver.Version, error) {
	v, ok := firstVersions[major]
	if !ok {
		return nil, fmt.Errorf("unknown major version %s", major)
	}
	return v, nil
}

// LastKindVersion returns the [*kindversion.KindVersion] representing the last
// known version of the kindversion.
func LastKindVersion() *kindversion.KindVersion {
	return KindVersion_V1_0_1
}

// LastVersion returns the [semver.Version] representing the last known version
// of the kindversion.
func LastVersion() *semver.Version {
	return SemVer_V1_0_1
}

// LastKindVersion returns the [api.KindVersionName] representing the last
// known version of the Kind.
func LastKindVersionName() api.KindVersionName {
	return api.NewKindVersionName(KindName, *SemVer_V1_0_1)
}

var (
	// lastVersions is a map of major version number (without "v" prefix) to
	// the last known kindversion version in that major version series.
	lastVersions = map[string]*semver.Version{
		"1": SemVer_V1_0_1,
	}
)

// LastVersionIn returns the [semver.Version] representing the last known
// version of the kindversion in the supplied major version series.
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
	KindVersion_V1_0_0 = kindversion.New(
		kindversion.WithKind(Kind),
		kindversion.WithVersion(*SemVer_V1_0_0),
		kindversion.WithSchema(Schema_V1_0_0),
		kindversion.WithSchemaJSON(SchemaJSON_V1_0_0),
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
	KindVersion_V1_0_1 = kindversion.New(
		kindversion.WithKind(Kind),
		kindversion.WithVersion(*SemVer_V1_0_1),
		kindversion.WithSchema(Schema_V1_0_1),
		kindversion.WithSchemaJSON(SchemaJSON_V1_0_1),
	)
}
