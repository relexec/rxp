package query_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/google/uuid"
	apidomain "github.com/relexec/rxp/api/domain"
	apikind "github.com/relexec/rxp/api/kind"
	apikindversion "github.com/relexec/rxp/api/kindversion"
	apiquery "github.com/relexec/rxp/api/query"
	"github.com/stretchr/testify/require"
)

func TestContainsPredicate(t *testing.T) {
	isKindish := func(p apiquery.Predicate) bool {
		switch p.(type) {
		case
			apikind.NamePredicate,
			apikind.UUIDPredicate,
			apikind.KindPredicate,
			apikindversion.NamePredicate,
			apikindversion.KindVersionPredicate:
			return true
		default:
			return false
		}
	}
	ku := uuid.NewString()
	kn := apikind.Name("SomeKind")
	k := apikind.Kind{
		UUID: ku,
		Name: kn,
	}
	v, err := semver.NewVersion("v0.0.1")
	require.Nil(t, err)
	kvn := apikindversion.NewName(kn, *v)
	kv := apikindversion.KindVersion{
		Kind:    k,
		Version: *v,
	}
	cases := []struct {
		name    string
		subject apiquery.Expression
		exp     bool
	}{
		{
			"nil is not a kind predicate",
			nil,
			false,
		},
		{
			"NameEqual",
			apidomain.NameEqual(apidomain.Name("some.domain")),
			false,
		},
		{
			"Or with two NameEquals",
			apiquery.Or(
				apidomain.NameEqual(apidomain.Name("some.domain")),
				apidomain.NameEqual(apidomain.Name("other.domain")),
			),
			false,
		},
		{
			"And with two NameEquals",
			apiquery.And(
				apidomain.NameEqual(apidomain.Name("some.domain")),
				apidomain.NameEqual(apidomain.Name("other.domain")),
			),
			false,
		},
		{
			"KindNameEqual",
			apikind.NameEqual(kn),
			true,
		},
		{
			"KindUUIDEqual",
			apikind.UUIDEqual(ku),
			true,
		},
		{
			"Or with KindNameEqual and NameEqual",
			apiquery.Or(
				apikind.NameEqual(kn),
				apidomain.NameEqual(apidomain.Name("other.domain")),
			),
			true,
		},
		{
			"And with KindNameEqual and NameEqual",
			apiquery.And(
				apikind.NameEqual(kn),
				apidomain.NameEqual(apidomain.Name("other.domain")),
			),
			true,
		},
		{
			"KindEqual",
			apikind.Equal(&k),
			true,
		},
		{
			"KindVersionEqual",
			apikindversion.Equal(&kv),
			true,
		},
		{
			"KindVersionNameEqual",
			apikindversion.NameEqual(kvn),
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			got := apiquery.ContainsPredicate(c.subject, isKindish)
			require.Equal(c.exp, got)
		})
	}
}
