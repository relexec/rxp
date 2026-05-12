package meta_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	authorv1 "github.com/relexec/rxp/testing/fixtures/author/v1"
	bookv1 "github.com/relexec/rxp/testing/fixtures/book/v1"
	"github.com/relexec/rxp/types"
)

func TestMeta_Diff(t *testing.T) {
	cases := []struct {
		name         string
		a            types.Meta
		b            any
		expError     string
		expDifferent bool
		expDiffs     []cmp.Difference
	}{
		{
			"cannot compare different types",
			bookv1.Meta_V1_0_0,
			"",
			"incompatible type comparison",
			false,
			nil,
		},
		{
			"same meta no diff",
			bookv1.Meta_V1_0_0,
			bookv1.Meta_V1_0_0,
			"",
			false,
			nil,
		},
		{
			"new meta",
			bookv1.Meta_V1_0_0,
			cmp.Zero,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("kindversion"),
					cmp.DifferenceTypeAdd,
					string(bookv1.KindVersion_V1_0_0),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("namescope"),
					cmp.DifferenceTypeAdd,
					types.NamescopeNamespace,
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("schema"),
					cmp.DifferenceTypeAdd,
					bookv1.SchemaJSON_V1_0_0,
					nil,
				),
			},
		},
		{
			"different kind and schema",
			bookv1.Meta_V1_0_0,
			authorv1.Meta_V1_0_0,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("kindversion"),
					cmp.DifferenceTypeModify,
					string(bookv1.KindVersion_V1_0_0),
					string(authorv1.KindVersion_V1_0_0),
				),
				cmp.NewDifference(
					fieldpath.FromString("schema"),
					cmp.DifferenceTypeModify,
					bookv1.SchemaJSON_V1_0_0,
					authorv1.SchemaJSON_V1_0_0,
				),
			},
		},
		{
			"different version and schema",
			bookv1.Meta_V1_0_0,
			bookv1.Meta_V1_0_1,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("kindversion"),
					cmp.DifferenceTypeModify,
					string(bookv1.KindVersion_V1_0_0),
					string(bookv1.KindVersion_V1_0_1),
				),
				cmp.NewDifference(
					fieldpath.FromString("schema"),
					cmp.DifferenceTypeModify,
					bookv1.SchemaJSON_V1_0_0,
					bookv1.SchemaJSON_V1_0_1,
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
