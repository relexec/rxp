package api

import (
	"strings"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/errors"
)

const (
	KindVersionSeparator = "@"
)

// KindVersion is a string that can contain a Kind and optionally a SemVer
// version string that uniquely identifies a type of Meta.
//
// A KindVersion string has the format <kind>[@<version>], where <kind> is a
// valid KindName and the optional <version> component must be a valid SemVer
// version string.
type KindVersion string

// Validate returns an error if the KindVersion is invalid.
func (kv KindVersion) Validate() error {
	if len(kv) == 0 {
		return errors.ErrKindNameEmpty
	}
	k := kv.Kind()
	err := k.Validate()
	if err != nil {
		return err
	}
	_, err = kv.Version()
	if err != nil {
		return err
	}
	return nil
}

// Kind returns the Kind identifier of the KindVersion. Note that this does not
// attempt to do any validation of the kind string.
func (kv KindVersion) Kind() KindName {
	parts := strings.SplitN(string(kv), KindVersionSeparator, 2)
	return KindName(parts[0])
}

// VersionString returns the SemVer version string from the optional version
// string component of the KindVersion. Note this does not attempt to do any
// validation of the version string.
func (kv KindVersion) VersionString() string {
	parts := strings.SplitN(string(kv), KindVersionSeparator, 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// Version returns the [semver.Version] object from the optional version string
// component of the KindVersion.
func (kv KindVersion) Version() (*semver.Version, error) {
	vs := kv.VersionString()
	if vs == "" {
		return nil, nil
	}
	return semver.StrictNewVersion(vs)
}

// NewKindVersion returns a KindVersion from a supplied KindName and
// [semver.Version].
func NewKindVersion(kind KindName, ver semver.Version) KindVersion {
	return KindVersion(string(kind) + KindVersionSeparator + ver.String())
}
