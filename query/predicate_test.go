package query_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/google/uuid"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kindversion"
	"github.com/relexec/rxp/query"
	"github.com/stretchr/testify/require"
)

func TestContainsPredicate(t *testing.T) {
	isKindish := func(p query.Predicate) bool {
		switch p.(type) {
		case
			kind.NamePredicate,
			kind.UUIDPredicate,
			kind.KindPredicate,
			kindversion.NamePredicate,
			kindversion.KindVersionPredicate:
			return true
		default:
			return false
		}
	}
	ku := uuid.NewString()
	kn := kind.Name("SomeKind")
	k := kind.Kind{
		UUID: ku,
		Name: kn,
	}
	v, err := semver.NewVersion("v0.0.1")
	require.Nil(t, err)
	kvn := kindversion.NewName(kn, *v)
	kv := kindversion.KindVersion{
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
			"NameEqual",
			domain.NameEqual(domain.Name("some.domain")),
			false,
		},
		{
			"Or with two NameEquals",
			query.Or(
				domain.NameEqual(domain.Name("some.domain")),
				domain.NameEqual(domain.Name("other.domain")),
			),
			false,
		},
		{
			"And with two NameEquals",
			query.And(
				domain.NameEqual(domain.Name("some.domain")),
				domain.NameEqual(domain.Name("other.domain")),
			),
			false,
		},
		{
			"KindNameEqual",
			kind.NameEqual(kn),
			true,
		},
		{
			"KindUUIDEqual",
			kind.UUIDEqual(ku),
			true,
		},
		{
			"Or with KindNameEqual and NameEqual",
			query.Or(
				kind.NameEqual(kn),
				domain.NameEqual(domain.Name("other.domain")),
			),
			true,
		},
		{
			"And with KindNameEqual and NameEqual",
			query.And(
				kind.NameEqual(kn),
				domain.NameEqual(domain.Name("other.domain")),
			),
			true,
		},
		{
			"KindEqual",
			kind.Equal(&k),
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
