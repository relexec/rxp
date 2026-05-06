package types_test

import (
	"testing"

	"github.com/relexec/rxp/types"
	"github.com/stretchr/testify/require"
)

func TestKind(t *testing.T) {
	cases := []struct {
		name             string
		subject          types.Kind
		expValidateError string
	}{
		{
			"empty kind version",
			"",
			"kind cannot be empty",
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
			"valid kind with dots",
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
