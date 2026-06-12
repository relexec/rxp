package expression_test

import (
	"testing"

	"github.com/relexec/rxp"
	"github.com/relexec/rxp/query/expression"
	"github.com/relexec/rxp/testing/fixtures/platform"
	"github.com/stretchr/testify/require"
)

func TestExpression_ContainsKindPredicate(t *testing.T) {
	cases := []struct {
		name    string
		subject expression.Expression
		exp     bool
	}{
		{
			"nil is not a kind predicate",
			nil,
			false,
		},
		{
			"DomainNameEqual",
			expression.DomainNameEqual(rxp.DomainName("some.domain")),
			false,
		},
		{
			"Or with two DomainNameEquals",
			expression.Or(
				expression.DomainNameEqual(rxp.DomainName("some.domain")),
				expression.DomainNameEqual(rxp.DomainName("other.domain")),
			),
			false,
		},
		{
			"And with two DomainNameEquals",
			expression.And(
				expression.DomainNameEqual(rxp.DomainName("some.domain")),
				expression.DomainNameEqual(rxp.DomainName("other.domain")),
			),
			false,
		},
		{
			"KindNameEqual",
			expression.KindNameEqual(rxp.KindName("some.kind")),
			true,
		},
		{
			"KindUUIDEqual",
			expression.KindUUIDEqual(platform.KindUUID),
			true,
		},
		{
			"Or with KindNameEqual and DomainNameEqual",
			expression.Or(
				expression.KindNameEqual(rxp.KindName("some.kind")),
				expression.DomainNameEqual(rxp.DomainName("other.domain")),
			),
			true,
		},
		{
			"And with KindNameEqual and DomainNameEqual",
			expression.And(
				expression.KindNameEqual(rxp.KindName("some.kind")),
				expression.DomainNameEqual(rxp.DomainName("other.domain")),
			),
			true,
		},
		{
			"KindEqual",
			expression.KindEqual(platform.Kind),
			true,
		},
		{
			"KindVersionEqual",
			expression.KindVersionEqual(platform.LastKindVersion()),
			true,
		},
		{
			"KindVersionNameEqual",
			expression.KindVersionNameEqual(platform.LastKindVersionName()),
			true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			got := expression.ContainsKindPredicate(c.subject)
			require.Equal(c.exp, got)
		})
	}
}
