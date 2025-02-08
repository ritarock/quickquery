package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  string
		want Query
	}{
		{
			name: "basic query",
			arg:  "select * from users.csv",
			want: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
			}},
		},
		{
			name: "basic query",
			arg:  "select id, name from users.csv",
			want: Query{Clauses: []string{
				"SELECT", "id", ",", "name", "FROM", "users.csv",
			}},
		},
		{
			name: "query with where",
			arg:  "select id, name from users.csv where id >= 2",
			want: Query{Clauses: []string{
				"SELECT", "id", ",", "name", "FROM", "users.csv",
				"WHERE", "id", ">=", "2",
			}},
		},
		{
			name: "query with where",
			arg:  "select id, name from users.csv where id >= 2 and team_id = 3",
			want: Query{Clauses: []string{
				"SELECT", "id", ",", "name", "FROM", "users.csv",
				"WHERE", "id", ">=", "2", "AND", "team_id", "=", "3",
			}},
		},
		{
			name: "query with order",
			arg:  "select id, name from users.csv order by id",
			want: Query{Clauses: []string{
				"SELECT", "id", ",", "name", "FROM", "users.csv",
				"ORDER", "BY", "id",
			}},
		},
		{
			name: "query with order",
			arg:  "select id, name from users.csv where id >= 2 order by id",
			want: Query{Clauses: []string{
				"SELECT", "id", ",", "name", "FROM", "users.csv",
				"WHERE", "id", ">=", "2",
				"ORDER", "BY", "id",
			}},
		},
		{
			name: "query with limit",
			arg:  "select id, name from users.csv limit 3",
			want: Query{Clauses: []string{
				"SELECT", "id", ",", "name", "FROM", "users.csv",
				"LIMIT", "3",
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewQuery(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestQuery_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		query    Query
		hasError bool
	}{
		{
			name: "valied query",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
			}},
			hasError: false,
		},
		{
			name: "invalied query: no select",
			query: Query{Clauses: []string{
				"FROM", "users.csv",
			}},
			hasError: true,
		},
		{
			name:     "invalied query: empty query",
			query:    Query{Clauses: []string{}},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.query.Validate()
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
