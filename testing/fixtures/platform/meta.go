package platform

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	v1 "github.com/relexec/rxp/testing/fixtures/platform/v1"
	"github.com/relexec/rxp/types"
)

const (
	Kind      = v1.Kind
	Namescope = v1.Namescope
)

// FirstMeta returns the [types.Meta] representing the first known version of
// the Platform meta.
func FirstMeta() types.Meta {
	return v1.FirstMeta()
}

// FirstVersion returns the [semver.Version] representing the first known
// version of the Platform meta.
func FirstVersion() *semver.Version {
	return v1.FirstVersion()
}

// FirstKindVersion returns the [types.KindVersion] representing the first
// known version of the Platform Kind.
func FirstKindVersion() types.KindVersion {
	return types.KindVersion(Kind + "@" + FirstVersion().String())
}

var (
	// firstVersions is a map of major version number (without "v" prefix) to
	// the first known Platform meta version in that major version series.
	firstVersions = map[string]*semver.Version{
		"1": v1.FirstVersion(),
	}
)

// FirstVersionIn returns the [semver.Version] representing the first known
// version of the Platform meta in the supplied major version series.
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
// Platform meta.
func LastMeta() types.Meta {
	return v1.LastMeta()
}

// LastVersion returns the [semver.Version] representing the last known version
// of the Platform meta.
func LastVersion() *semver.Version {
	return v1.LastVersion()
}

// LastKindVersion returns the [types.KindVersion] representing the last known
// version of the Platform Kind.
func LastKindVersion() types.KindVersion {
	return types.KindVersion(Kind + "@" + LastVersion().String())
}

var (
	// lastVersions is a map of major version number (without "v" prefix) to
	// the last known Platform meta version in that major version series.
	lastVersions = map[string]*semver.Version{
		"1": v1.LastVersion(),
	}
)

// LastVersionIn returns the [semver.Version] representing the last known
// version of the Platform meta in the supplied major version series.
//
// If the supplied major version string is not known, returns an error.
func LastVersionIn(major string) (*semver.Version, error) {
	v, ok := lastVersions[major]
	if !ok {
		return nil, fmt.Errorf("unknown major version %s", major)
	}
	return v, nil
}
