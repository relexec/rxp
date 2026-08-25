package api

import (
	"strings"

	"github.com/Masterminds/semver/v3"

	apikind "github.com/relexec/rxp/api/kind"
	"github.com/relexec/rxp/errors"
)

const (
	KindVersionNameSeparator = "@"
)

// KindVersionName is a string that can contain a Kind and a SemVer version
// string that uniquely identifies a type of Object.
//
// A KindVersionName string has the format <kind>@<version>, where <kind> is a
// valid apikind.Name and <version> is a valid SemVer version string.
type KindVersionName string

// Validate returns an error if the KindVersionName is invalid.
func (kv KindVersionName) Validate() error {
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

// Kind returns the Kind identifier of the KindVersionName. Note that this does
// not attempt to do any validation of the kind string.
func (kv KindVersionName) Kind() apikind.Name {
	parts := strings.SplitN(string(kv), KindVersionNameSeparator, 2)
	return apikind.Name(parts[0])
}

// VersionString returns the SemVer version string from the optional version
// string component of the KindVersionName. Note this does not attempt to do
// any validation of the version string.
func (kv KindVersionName) VersionString() string {
	parts := strings.SplitN(string(kv), KindVersionNameSeparator, 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// Version returns the [semver.Version] object from the version string
// component of the KindVersionName.
func (kv KindVersionName) Version() (*semver.Version, error) {
	vs := kv.VersionString()
	if vs == "" {
		return nil, nil
	}
	return semver.StrictNewVersion(vs)
}

// NewKindVersionName returns a KindVersionName from a supplied apikind.Name and
// [semver.Version].
func NewKindVersionName(kind apikind.Name, ver semver.Version) KindVersionName {
	return KindVersionName(string(kind) + KindVersionNameSeparator + ver.String())
}
