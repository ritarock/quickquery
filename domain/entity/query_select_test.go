package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetSelect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  Select
	}{
		{
			name: "select single column",
			query: Query{Clauses: []string{
				"SELECT", "id", "FROM", "users.csv",
			}},
			want: []string{"id"},
		},
		{
			name: "select multiple columns",
			query: Query{Clauses: []string{
				"SELECT", "id", "name", "FROM", "users.csv",
			}},
			want: []string{"id", "name"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.query.GetSelect()
			assert.Equal(t, test.want, got)
		})
	}
}
