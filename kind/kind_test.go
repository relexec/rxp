package kind_test

import (
	"testing"

	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/testing/fixtures/author"
	"github.com/relexec/rxp/testing/fixtures/book"
	"github.com/relexec/rxp/types"
	"github.com/stretchr/testify/require"
)

func TestKind_Diff(t *testing.T) {
	cases := []struct {
		name         string
		a            types.Kind
		b            any
		expError     string
		expDifferent bool
		expDiffs     []cmp.Difference
	}{
		{
			"cannot compare different types",
			book.Kind,
			"",
			"incompatible type comparison",
			false,
			nil,
		},
		{
			"same kind no diff",
			book.Kind,
			book.Kind,
			"",
			false,
			nil,
		},
		{
			"new kind",
			book.Kind,
			cmp.Zero,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeAdd,
					string(book.KindName),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("namescope"),
					cmp.DifferenceTypeAdd,
					types.NamescopeNamespace,
					nil,
				),
			},
		},
		{
			"different kind name",
			book.Kind,
			author.Kind,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeModify,
					string(book.KindName),
					string(author.KindName),
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
