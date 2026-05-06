package types_test

import (
	"strings"
	"testing"

	"github.com/relexec/rxp/types"
	"github.com/stretchr/testify/require"
)

func TestDomain(t *testing.T) {
	cases := []struct {
		name             string
		subject          types.Domain
		expValidateError string
	}{
		{
			"empty domain is fine",
			"",
			"",
		},
		{
			"spaces not allowed",
			"spaces not allowed",
			"invalid characters",
		},
		{
			"max length exceeded",
			types.Domain(strings.Repeat("X", types.DomainMaxLength+1)),
			"max length exceeded",
		},
		{
			"cannot start with dot",
			".domain1",
			"first character must be letter or number",
		},
		{
			"valid domain with dots",
			"customer1.domain1",
			"",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require := require.New(t)
			err := c.subject.Validate()
			if c.expValidateError != "" {
				require.ErrorContains(err, c.expValidateError)
			} else {
				require.Nil(err)
			}
		})
	}
}
