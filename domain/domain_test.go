package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/relexec/rxp/domain"
	"github.com/relexec/rxp/system"
	"github.com/stretchr/testify/require"
)

func TestDomain_Validate(t *testing.T) {
	sys := system.System{UUID: uuid.NewString()}
	domWithSystemName := domain.Name("dom.with.system")
	domWithSystem := domain.Domain{
		UUID:   uuid.NewString(),
		System: &sys,
		Name:   domWithSystemName,
	}
	domWithParentNoRootName := domain.Name("dom.with.parent.no.root")
	domWithParentNoRoot := domain.Domain{
		UUID:   uuid.NewString(),
		Parent: &domWithSystem,
		Name:   domWithParentNoRootName,
	}
	domWithParentName := domain.Name("dom.with.parent")
	domWithParent := domain.Domain{
		UUID:   uuid.NewString(),
		Parent: &domWithSystem,
		Root:   &domWithSystem,
		Name:   domWithParentName,
	}
	secondSystem := system.System{UUID: uuid.NewString()}
	domWithParentDiffSystemName := domain.Name("dom.with.diff.system")
	domWithParentDiffSystem := domain.Domain{
		System: &secondSystem,
		UUID:   uuid.NewString(),
		Parent: &domWithSystem,
		Root:   &domWithSystem,
		Name:   domWithParentDiffSystemName,
	}

	cases := []struct {
		name     string
		subject  domain.Domain
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
