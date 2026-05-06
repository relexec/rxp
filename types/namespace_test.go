package types_test

import (
	"testing"

	"github.com/relexec/rxp/types"
	"github.com/stretchr/testify/require"
)

func TestNamespace(t *testing.T) {
	cases := []struct {
		name             string
		subject          types.Namespace
		expValidateError string
	}{
		{
			"empty namespace is fine",
			"",
			"",
		},
		{
			"spaces not allowed",
			"spaces not allowed",
			"invalid characters",
		},
		{
			"cannot start with dot",
			".namespace1",
			"first character must be letter or number",
		},
		{
			"valid namespace with dots",
			"customer1.namespace1",
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
