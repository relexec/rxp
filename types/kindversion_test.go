package types_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/relexec/rxp/types"
	"github.com/stretchr/testify/require"
)

func TestKindVersion(t *testing.T) {
	cases := []struct {
		name             string
		subject          types.KindVersion
		expValidateError string
		expKind          string
		expVersionString string
		expVersion       *semver.Version
		expVersionError  string
	}{
		{
			"empty kind version",
			"",
			"invalid kind: kind cannot be empty",
			"",
			"",
			nil,
			"",
		},
		{
			"spaces not allowed",
			"spaces not allowed",
			"invalid kind: invalid characters",
			"spaces not allowed",
			"",
			nil,
			"",
		},
		{
			"cannot start with dot",
			".flow.temporal.io",
			"first character must be letter or number",
			".flow.temporal.io",
			"",
			nil,
			"",
		},
		{
			"double periods not allowed",
			"a..b",
			"invalid kind: repeated periods",
			"a..b",
			"",
			nil,
			"",
		},
		{
			"multiple @ signs",
			"flow.temporal.io@1.0.0@1.1.0",
			"invalid characters in version",
			"flow.temporal.io",
			"1.0.0@1.1.0",
			nil,
			"invalid characters in version",
		},
		{
			"v prefix is not valid semver",
			"flow.temporal.io@v1.1.0",
			"invalid characters in version",
			"flow.temporal.io",
			"v1.1.0",
			nil,
			"invalid characters in version",
		},
		{
			"valid kind and version",
			"flow.temporal.io@1.1.0",
			"",
			"flow.temporal.io",
			"1.1.0",
			nil,
			"",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			err := c.subject.Validate()
			if c.expValidateError != "" {
				require.ErrorContains(err, c.expValidateError)
			} else {
				require.Nil(err)
			}
			k := string(c.subject.Kind())
			require.Equal(c.expKind, k)
			vs := c.subject.VersionString()
			require.Equal(c.expVersionString, vs)
			v, err := c.subject.Version()
			if c.expVersionError != "" {
				require.ErrorContains(err, c.expVersionError)
			} else {
				require.Nil(err)
				if c.expVersion != nil {
					require.True(c.expVersion.Equal(v))
				}
			}
		})
	}
}
