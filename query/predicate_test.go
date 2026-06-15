package query_test

import (
	"testing"

	"github.com/relexec/rxp"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/kind"
	"github.com/relexec/rxp/kind/kindversion"
	"github.com/relexec/rxp/query"
	"github.com/relexec/rxp/testing/fixtures/platform"
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
			domain.NameEqual(rxp.DomainName("some.domain")),
			false,
		},
		{
			"Or with two DomainNameEquals",
			query.Or(
				domain.NameEqual(rxp.DomainName("some.domain")),
				domain.NameEqual(rxp.DomainName("other.domain")),
			),
			false,
		},
		{
			"And with two DomainNameEquals",
			query.And(
				domain.NameEqual(rxp.DomainName("some.domain")),
				domain.NameEqual(rxp.DomainName("other.domain")),
			),
			false,
		},
		{
			"KindNameEqual",
			kind.NameEqual(rxp.KindName("some.kind")),
			true,
		},
		{
			"KindUUIDEqual",
			kind.UUIDEqual(platform.KindUUID),
			true,
		},
		{
			"Or with KindNameEqual and DomainNameEqual",
			query.Or(
				kind.NameEqual(rxp.KindName("some.kind")),
				domain.NameEqual(rxp.DomainName("other.domain")),
			),
			true,
		},
		{
			"And with KindNameEqual and DomainNameEqual",
			query.And(
				kind.NameEqual(rxp.KindName("some.kind")),
				domain.NameEqual(rxp.DomainName("other.domain")),
			),
			true,
		},
		{
			"KindEqual",
			kind.Equal(platform.Kind),
			true,
		},
		{
			"KindVersionEqual",
			kindversion.Equal(platform.LastKindVersion()),
			true,
		},
		{
			"KindVersionNameEqual",
			kindversion.NameEqual(platform.LastKindVersionName()),
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
