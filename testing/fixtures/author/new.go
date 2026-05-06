package author

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/object"
	v1 "github.com/relexec/rxp/testing/fixtures/author/v1"
	"github.com/relexec/rxp/types"
	"github.com/relexec/rxp/version"
)

// New returns a new Author [Object] with the latest known Version.
func New(opts ...object.Option) *object.Object {
	kv := LatestKindVersion()
	return object.New(kv, opts...)
}

// NewAtVersion returns a new Author [Object] with the supplied version.  The
// supplied version parameter can be either a string or [semver.Version].
// String parameters can be either a full SemVer string or a single major
// version series, e.g. "2" or "v2". If the string is just a major version
// series, the latest version in that major version series is used.
func NewAtVersion(ver any, opts ...object.Option) (*object.Object, error) {
	var v *semver.Version
	switch ver := ver.(type) {
	case semver.Version:
		v = &ver
	case *semver.Version:
		v = ver
	case string:
		sv, err := semver.NewVersion(ver)
		if err == nil {
			v = sv
			break
		}
		// user may have passed just a major version. If so, grab
		// the latest version from that major version series.
		major, err := version.Major(ver)
		if err != nil {
			return nil, err
		}
		sv, err = LatestVersionIn(major)
		if err != nil {
			return nil, err
		}
		v = sv
	default:
		return nil, fmt.Errorf(
			"unsupported type %T passed to NewAtVersion", ver,
		)
	}
	var kv types.KindVersion

	vstr := v.Original()
	switch vstr {
	case v1.Version_V1_0_0:
		kv = v1.Meta_V1_0_0.KindVersion()
	case v1.Version_V1_0_1:
		kv = v1.Meta_V1_0_1.KindVersion()
	default:
		return nil, fmt.Errorf(
			"unknown version %s of kind %q",
			vstr, Kind,
		)
	}

	return object.New(kv, opts...), nil
}
