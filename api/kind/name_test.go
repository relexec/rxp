package apikind_test

import (
	"testing"

	apikind "github.com/relexec/rxp/api/kind"

	"github.com/stretchr/testify/require"
)

func TestKindName(t *testing.T) {
	cases := []struct {
		name             string
		subject          apikind.Name
		expValidateError string
	}{
		{
			"empty kind",
			"",
			"kind name cannot be empty",
		},
		{
			"spaces not allowed",
			"spaces not allowed",
			"invalid characters",
		},
		{
			"cannot start with dot",
			".flow.temporal.io",
			"first character must be letter or number",
		},
		{
			"valid kind name with dots",
			"flow.temporal.io",
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
