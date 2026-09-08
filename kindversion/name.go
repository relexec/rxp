package kindversion

import (
	"strings"

	"github.com/Masterminds/semver/v3"

	"github.com/relexec/rxp/errors"
	"github.com/relexec/rxp/kind"
)

const (
	NameSeparator = "@"
)

// Name is a string that can contain a Kind and a SemVer version
// string that uniquely identifies a type of Object.
//
// A Name string has the format <kind>@<version>, where <kind> is a
// valid kind.Name and <version> is a valid SemVer version string.
type Name string

// Validate returns an error if the Name is invalid.
func (n Name) Validate() error {
	if len(n) == 0 {
		return errors.ErrKindNameEmpty
	}
	k := n.Kind()
	err := k.Validate()
	if err != nil {
		return err
	}
	_, err = n.Version()
	if err != nil {
		return err
	}
	return nil
}

// Kind returns the Kind identifier of the Name. Note that this does
// not attempt to do any validation of the kind string.
func (n Name) Kind() kind.Name {
	parts := strings.SplitN(string(n), NameSeparator, 2)
	return kind.Name(parts[0])
}

// VersionString returns the SemVer version string from the optional version
// string component of the Name. Note this does not attempt to do
// any validation of the version string.
func (n Name) VersionString() string {
	parts := strings.SplitN(string(n), NameSeparator, 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// Version returns the [semver.Version] object from the version string
// component of the Name.
func (n Name) Version() (*semver.Version, error) {
	vs := n.VersionString()
	if vs == "" {
		return nil, nil
	}
	return semver.StrictNewVersion(vs)
}

// NewName returns a Name from a supplied kind.Name and
// [semver.Version].
func NewName(kind kind.Name, ver semver.Version) Name {
	return Name(string(kind) + NameSeparator + ver.String())
}
