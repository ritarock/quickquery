package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetLimit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  Limit
	}{
		{
			name: "query with LIMIT",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv", "LIMIT", "3",
			}},
			want: Limit(3),
		},
		{
			name: "query without LIMIT",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
			}},
			want: Limit(0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.query.GetLimit()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestLimit_IsLimit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		limit Limit
		want  bool
	}{
		{
			name:  "query with LIMIT",
			limit: Limit(3),
			want:  true,
		},
		{
			name:  "query without LIMIT",
			limit: Limit(0),
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.limit.IsLimit()
			assert.Equal(t, test.want, got)
		})
	}
}
