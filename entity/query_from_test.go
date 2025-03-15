package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetFile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		query    Query
		want     string
		hasError bool
	}{
		{
			name:     "query with FROM csvfile",
			query:    Query{[]string{"SELECT", "*", "FROM", "users.csv"}},
			want:     "users.csv",
			hasError: false,
		},
		{
			name:     "query without FROM",
			query:    Query{[]string{"SELECT", "*", "users.csv"}},
			want:     "",
			hasError: true,
		},
		{
			name:     "query without csvfile",
			query:    Query{[]string{"SELECT", "*", "FROM"}},
			want:     "",
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := test.query.GetFile()
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
