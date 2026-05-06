package author

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	v1 "github.com/relexec/rxp/testing/fixtures/author/v1"
	"github.com/relexec/rxp/types"
)

const (
	Kind = "author.testing.rxp"
)

// LatestVersion returns the [semver.Version] representing the latest known
// version of the Author object meta.
func LatestVersion() *semver.Version {
	return v1.LatestVersion()
}

// LatestKindVersion returns the [types.KindVersion] representing the latest
// known version of the Author Kind.
func LatestKindVersion() types.KindVersion {
	return types.KindVersion(Kind + "@" + LatestVersion().Original())
}

var (
	// latestVersions is a map of major version number (without "v" prefix) to
	// the latest known Author object meta version in that major version
	// series.
	latestVersions = map[string]*semver.Version{
		"1": v1.LatestVersion(),
	}
)

// LatestVersionIn returns the [semver.Version] representing the latest know
// version of the Author object meta in the supplied major version series.
//
// If the supplied major version string is not known, returns an error.
func LatestVersionIn(major string) (*semver.Version, error) {
	v, ok := latestVersions[major]
	if !ok {
		return nil, fmt.Errorf("unknown major version %s", major)
	}
	return v, nil
}
