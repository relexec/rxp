package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/relexec/rxp/api"
	apisystem "github.com/relexec/rxp/api/system"
	"github.com/stretchr/testify/require"
)

func TestDomain_Validate(t *testing.T) {
	sys := apisystem.System{UUID: uuid.NewString()}
	domWithSystemName := api.DomainName("dom.with.system")
	domWithSystem := api.Domain{
		UUID:   uuid.NewString(),
		System: &sys,
		Name:   domWithSystemName,
	}
	domWithParentNoRootName := api.DomainName("dom.with.parent.no.root")
	domWithParentNoRoot := api.Domain{
		UUID:   uuid.NewString(),
		Parent: &domWithSystem,
		Name:   domWithParentNoRootName,
	}
	domWithParentName := api.DomainName("dom.with.parent")
	domWithParent := api.Domain{
		UUID:   uuid.NewString(),
		Parent: &domWithSystem,
		Root:   &domWithSystem,
		Name:   domWithParentName,
	}
	secondSystem := apisystem.System{UUID: uuid.NewString()}
	domWithParentDiffSystemName := api.DomainName("dom.with.diff.system")
	domWithParentDiffSystem := api.Domain{
		System: &secondSystem,
		UUID:   uuid.NewString(),
		Parent: &domWithSystem,
		Root:   &domWithSystem,
		Name:   domWithParentDiffSystemName,
	}

	cases := []struct {
		name     string
		subject  api.Domain
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
