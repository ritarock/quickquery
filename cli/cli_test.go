package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		hasError bool
	}{
		{
			name:     "pass",
			args:     []string{"select * from ./sample.csv"},
			hasError: false,
		},
		{
			name:     "failed: not enough args",
			args:     []string{},
			hasError: true,
		},
		{
			name:     "failed: not enough args",
			args:     []string{""},
			hasError: true,
		},
		{
			name:     "failed: too many args",
			args:     []string{"select * from ./sample.csv", "test"},
			hasError: true,
		},
	}

	for _, test := range tests {
		err := validateArgs(test.args)
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
