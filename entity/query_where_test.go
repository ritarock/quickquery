package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetWhere(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  Where
	}{
		{
			name: "query with where",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv", "WHERE", "id", "=", "3",
			}},
			want: Where{
				{"id", "=", "3"},
			},
		},
		{
			name: "query without where",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
			}},
			want: Where{},
		},
		{
			name: "query with order",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
				"WHERE", "id", "=", "3", "ORDER", "BY", "id",
			}},
			want: Where{
				{"id", "=", "3"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.query.GetWhere()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestWhere_IsWhere(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		where Where
		want  bool
	}{
		{
			name: "query with where",
			where: Where{{
				"id", "=", "3",
			}},
			want: true,
		},
		{
			name:  "query without where",
			where: Where{{}},
			want:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.where.IsWhere()
			assert.Equal(t, test.want, got)
		})
	}
}
