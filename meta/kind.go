package meta

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
)

// KindVersion returns a KindVersion string from a supplied Kind string and
// Version. Version can be either a string or a [semver.Version].
//
// Panics if Version is neither string or [semver.Version] so this is really
// only meant for internal usage.
func KindVersion(kind string, version any) string {
	switch version := version.(type) {
	case string:
		return kind + "@" + strings.TrimPrefix(version, "v")
	case semver.Version:
		return kind + "@" + version.String()
	case *semver.Version:
		return kind + "@" + version.String()
	default:
		msg := fmt.Sprintf("unhandled version type %T", version)
		panic(msg)
	}
}
