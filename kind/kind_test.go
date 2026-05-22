package kind_test

import (
	"testing"

	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/testing/fixtures/application"
	"github.com/relexec/rxp/testing/fixtures/platform"
	"github.com/stretchr/testify/require"
)

func TestKind_Diff(t *testing.T) {
	cases := []struct {
		name         string
		a            *kind.Kind
		b            any
		expError     string
		expDifferent bool
		expDiffs     []cmp.Difference
	}{
		{
			"cannot compare different types",
			platform.Kind,
			"",
			"incompatible type comparison",
			false,
			nil,
		},
		{
			"same kind no diff",
			platform.Kind,
			platform.Kind,
			"",
			false,
			nil,
		},
		{
			"new kind",
			platform.Kind,
			cmp.Zero,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeAdd,
					string(platform.KindName),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("scope"),
					cmp.DifferenceTypeAdd,
					api.ScopeSystem,
					nil,
				),
			},
		},
		{
			"different kind name and scope",
			platform.Kind,
			application.Kind,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeModify,
					string(platform.KindName),
					string(application.KindName),
				),
				cmp.NewDifference(
					fieldpath.FromString("scope"),
					cmp.DifferenceTypeModify,
					api.ScopeSystem,
					api.ScopeDomain,
				),
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			d, err := c.a.Diff(c.b)
			if c.expError != "" {
				require.ErrorContains(err, c.expError)
			} else {
				require.Nil(err)
				require.NotNil(d)
				require.Equal(c.expDifferent, d.Different())
				if c.expDifferent {
					require.Equal(c.expDiffs, d.Differences())
				}
			}
		})
	}
}
