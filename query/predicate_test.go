package query_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/google/uuid"
	"github.com/relexec/rxp/api"
	apikind "github.com/relexec/rxp/api/kind"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/query"
	"github.com/stretchr/testify/require"
)

func TestContainsPredicate(t *testing.T) {
	isKindish := func(p query.Predicate) bool {
		switch p.(type) {
		case
			apikind.NamePredicate,
			apikind.UUIDPredicate,
			apikind.KindPredicate,
			kindversion.NamePredicate,
			kindversion.KindVersionPredicate:
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
	kvn := api.NewKindVersionName(kn, *v)
	kv := api.KindVersion{
		Kind:    k,
		Version: *v,
	}
	cases := []struct {
		name    string
		subject query.Expression
		exp     bool
	}{
		{
			"nil is not a kind predicate",
			nil,
			false,
		},
		{
			"DomainNameEqual",
			domain.NameEqual(api.DomainName("some.domain")),
			false,
		},
		{
			"Or with two DomainNameEquals",
			query.Or(
				domain.NameEqual(api.DomainName("some.domain")),
				domain.NameEqual(api.DomainName("other.domain")),
			),
			false,
		},
		{
			"And with two DomainNameEquals",
			query.And(
				domain.NameEqual(api.DomainName("some.domain")),
				domain.NameEqual(api.DomainName("other.domain")),
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
			"Or with KindNameEqual and DomainNameEqual",
			query.Or(
				apikind.NameEqual(kn),
				domain.NameEqual(api.DomainName("other.domain")),
			),
			true,
		},
		{
			"And with KindNameEqual and DomainNameEqual",
			query.And(
				apikind.NameEqual(kn),
				domain.NameEqual(api.DomainName("other.domain")),
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
			kindversion.Equal(&kv),
			true,
		},
		{
			"KindVersionNameEqual",
			kindversion.NameEqual(kvn),
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			got := query.ContainsPredicate(c.subject, isKindish)
			require.Equal(c.exp, got)
		})
	}
}
