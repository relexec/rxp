package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/relexec/rxp"
	"github.com/relexec/rxp/cmp"
	"github.com/relexec/rxp/cmp/fieldpath"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/system"
	"github.com/relexec/rxp/testing/fixtures"
	"github.com/stretchr/testify/require"
)

func TestDomain_Diff(t *testing.T) {
	domWithSystemUUID := uuid.NewString()
	domWithSystemName := rxp.DomainName("dom.with.system")
	domWithSystem := domain.New(
		domain.WithUUID(domWithSystemUUID),
		domain.WithSystem(fixtures.System),
		domain.WithName(domWithSystemName),
	)
	domWithParentUUID := uuid.NewString()
	domWithParentName := rxp.DomainName("dom.with.parent")
	domWithParent := domain.New(
		domain.WithUUID(domWithParentUUID),
		domain.WithParent(domWithSystem),
		domain.WithRoot(domWithSystem),
		domain.WithName(domWithParentName),
	)

	cases := []struct {
		name         string
		a            *domain.Domain
		b            any
		expError     string
		expDifferent bool
		expDiffs     []cmp.Difference
	}{
		{
			"cannot compare different types",
			fixtures.Domain,
			"",
			"incompatible type comparison",
			false,
			nil,
		},
		{
			"same domain no diff",
			fixtures.Domain,
			fixtures.Domain,
			"",
			false,
			nil,
		},
		{
			"new domain no system",
			fixtures.Domain,
			cmp.Zero,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("uuid"),
					cmp.DifferenceTypeAdd,
					fixtures.Domain.UUID(),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeAdd,
					string(fixtures.Domain.Name()),
					nil,
				),
			},
		},
		{
			"new domain with system",
			domWithSystem,
			cmp.Zero,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("system"),
					cmp.DifferenceTypeAdd,
					fixtures.System.UUID(),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("uuid"),
					cmp.DifferenceTypeAdd,
					domWithSystemUUID,
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeAdd,
					string(domWithSystemName),
					nil,
				),
			},
		},
		{
			"new domain with parent",
			domWithParent,
			cmp.Zero,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("uuid"),
					cmp.DifferenceTypeAdd,
					domWithParentUUID,
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeAdd,
					string(domWithParentName),
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("root"),
					cmp.DifferenceTypeAdd,
					domWithSystemUUID,
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("parent"),
					cmp.DifferenceTypeAdd,
					domWithSystemUUID,
					nil,
				),
			},
		},
		{
			"different uuid, system, name and parent",
			domWithParent,
			domWithSystem,
			"",
			true,
			[]cmp.Difference{
				cmp.NewDifference(
					fieldpath.FromString("system"),
					cmp.DifferenceTypeAdd,
					nil,
					fixtures.SystemUUID,
				),
				cmp.NewDifference(
					fieldpath.FromString("uuid"),
					cmp.DifferenceTypeModify,
					domWithParentUUID,
					domWithSystemUUID,
				),
				cmp.NewDifference(
					fieldpath.FromString("name"),
					cmp.DifferenceTypeModify,
					string(domWithParentName),
					string(domWithSystemName),
				),
				cmp.NewDifference(
					fieldpath.FromString("root"),
					cmp.DifferenceTypeRemove,
					domWithSystemUUID,
					nil,
				),
				cmp.NewDifference(
					fieldpath.FromString("parent"),
					cmp.DifferenceTypeRemove,
					domWithSystemUUID,
					nil,
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

func TestDomain_Validate(t *testing.T) {
	domWithSystemName := rxp.DomainName("dom.with.system")
	domWithSystem := domain.New(
		domain.WithUUID(uuid.NewString()),
		domain.WithSystem(fixtures.System),
		domain.WithName(domWithSystemName),
	)
	domWithParentNoRootName := rxp.DomainName("dom.with.parent.no.root")
	domWithParentNoRoot := domain.New(
		domain.WithUUID(uuid.NewString()),
		domain.WithParent(domWithSystem),
		domain.WithName(domWithParentNoRootName),
	)
	domWithParentName := rxp.DomainName("dom.with.parent")
	domWithParent := domain.New(
		domain.WithUUID(uuid.NewString()),
		domain.WithParent(domWithSystem),
		domain.WithRoot(domWithSystem),
		domain.WithName(domWithParentName),
	)
	secondSystem := system.New(system.WithUUID(uuid.NewString()))
	domWithParentDiffSystemName := rxp.DomainName("dom.with.diff.system")
	domWithParentDiffSystem := domain.New(
		domain.WithSystem(secondSystem),
		domain.WithUUID(uuid.NewString()),
		domain.WithParent(domWithSystem),
		domain.WithRoot(domWithSystem),
		domain.WithName(domWithParentDiffSystemName),
	)

	cases := []struct {
		name     string
		subject  *domain.Domain
		expError string
	}{
		{
			"known valid domain",
			fixtures.Domain,
			"",
		},
		{
			"parent not empty, system empty",
			domWithParent,
			"",
		},
		{
			"specify parent but no root",
			domWithParentNoRoot,
			"root required when parent specified",
		},
		{
			"different system uuid in root",
			domWithParentDiffSystem,
			"root system must be same",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			err := c.subject.Validate()
			if c.expError != "" {
				require.ErrorContains(err, c.expError)
			} else {
				require.Nil(err)
			}
		})
	}
}
