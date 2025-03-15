package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateArgs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		args     []string
		hasError bool
	}{
		{
			name:     "succeed",
			args:     []string{"SELECT * FROM sample.csv"},
			hasError: false,
		},
		{
			name:     "failed: no args",
			args:     []string{},
			hasError: true,
		},
		{
			name:     "failed: empty args",
			args:     []string{""},
			hasError: true,
		},
		{
			name:     "failed: too many args",
			args:     []string{"arg1", "arg2"},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := validateArgs(test.args)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
