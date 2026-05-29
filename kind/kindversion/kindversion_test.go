package kindversion_test

import (
	"testing"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/testing/fixtures/application"
	"github.com/relexec/rxp/testing/fixtures/platform"
	"github.com/stretchr/testify/require"
)

func TestKindVersion_Diff(t *testing.T) {
	cases := []struct {
		name         string
		a            *kindversion.KindVersion
		b            any
		expError     string
		expDifferent bool
		expDiffs     []cmp.Difference
	}{
		{
			"cannot compare different types",
			platform.KindVersion_V1_0_0,
			"",
			"incompatible type comparison",
			false,
			nil,
		},
		{
			"same kindversion no diff",
			platform.KindVersion_V1_0_0,
			platform.KindVersion_V1_0_0,
			"",
			false,
			nil,
		},
		{
			"new kindversion",
			platform.KindVersion_V1_0_0,
			cmp.Zero,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("kind"),
					cmp.DifferenceTypeAdd,
					string(platform.KindName),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("version"),
					cmp.DifferenceTypeAdd,
					platform.SemVer_V1_0_0.String(),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("schema"),
					cmp.DifferenceTypeAdd,
					platform.SchemaJSON_V1_0_0,
					nil,
				),
			},
		},
		{
			"different kind and schema",
			platform.KindVersion_V1_0_0,
			application.KindVersion_V1_0_0,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("kind"),
					cmp.DifferenceTypeModify,
					string(platform.KindName),
					string(application.KindName),
				),
				cmp.NewDifference(
					fieldpath.FromString("schema"),
					cmp.DifferenceTypeModify,
					platform.SchemaJSON_V1_0_0,
					application.SchemaJSON_V1_0_0,
				),
			},
		},
		{
			"different version and schema",
			platform.KindVersion_V1_0_0,
			platform.KindVersion_V1_0_1,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("version"),
					cmp.DifferenceTypeModify,
					string(platform.SemVer_V1_0_0.String()),
					string(platform.SemVer_V1_0_1.String()),
				),
				cmp.NewDifference(
					fieldpath.FromString("schema"),
					cmp.DifferenceTypeModify,
					platform.SchemaJSON_V1_0_0,
					platform.SchemaJSON_V1_0_1,
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
