package apidomain_test

import (
	"strings"
	"testing"

	apidomain "github.com/relexec/rxp/api/domain"
	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	cases := []struct {
		name             string
		subject          apidomain.Name
		expValidateError string
	}{
		{
			"empty domain name not allowed",
			"",
			"name required",
		},
		{
			"spaces not allowed",
			"spaces not allowed",
			"invalid characters",
		},
		{
			"max length exceeded",
			apidomain.Name(
				strings.Repeat("X", apidomain.NameMaxLength+1),
			),
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
