package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/relexec/rxp/api"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/system"
	"github.com/stretchr/testify/require"
)

func TestDomain_Validate(t *testing.T) {
	sys := system.New(system.WithUUID(uuid.NewString()))
	domWithSystemName := api.DomainName("dom.with.system")
	domWithSystem := domain.New(
		domain.WithUUID(uuid.NewString()),
		domain.WithSystem(sys),
		domain.WithName(domWithSystemName),
	)
	domWithParentNoRootName := api.DomainName("dom.with.parent.no.root")
	domWithParentNoRoot := domain.New(
		domain.WithUUID(uuid.NewString()),
		domain.WithParent(domWithSystem),
		domain.WithName(domWithParentNoRootName),
	)
	domWithParentName := api.DomainName("dom.with.parent")
	domWithParent := domain.New(
		domain.WithUUID(uuid.NewString()),
		domain.WithParent(domWithSystem),
		domain.WithRoot(domWithSystem),
		domain.WithName(domWithParentName),
	)
	secondSystem := system.New(system.WithUUID(uuid.NewString()))
	domWithParentDiffSystemName := api.DomainName("dom.with.diff.system")
	domWithParentDiffSystem := domain.New(
		domain.WithSystem(secondSystem),
		domain.WithUUID(uuid.NewString()),
		domain.WithParent(domWithSystem),
		domain.WithRoot(domWithSystem),
		domain.WithName(domWithParentDiffSystemName),
	)

	cases := []struct {
		name     string
		subject  *api.Domain
		expError string
	}{
		{
			"known valid domain",
			domWithSystem,
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
