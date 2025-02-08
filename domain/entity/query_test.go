package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  Query
	}{
		{
			name:  "basic select query",
			input: "SELECT name FROM users.csv",
			want:  Query{Clauses: []string{"SELECT", "name", "FROM", "users.csv"}},
		},
		{
			name:  "query with where clause",
			input: "SELECT name FROM users.csv WHERE id = 10",
			want:  Query{Clauses: []string{"SELECT", "name", "FROM", "users.csv", "WHERE", "id", "=", "10"}},
		},
		{
			name:  "2 queries with where clause",
			input: "SELECT name FROM users.csv WHERE id = 10 and user_name = user10",
			want: Query{Clauses: []string{
				"SELECT", "name", "FROM", "users.csv", "WHERE",
				"id", "=", "10", "AND", "user_name", "=", "user10",
			}},
		},
		{
			name:  "query with order clause",
			input: "SELECT name FROM users.csv ORDER BY id DESC",
			want: Query{Clauses: []string{
				"SELECT", "name", "FROM", "users.csv",
				"ORDER", "BY", "id", "DESC",
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewQuery(test.input)
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
			name:     "valid query",
			query:    Query{Clauses: []string{"SELECT", "name", "FROM", "users.csv"}},
			hasError: false,
		},
		{
			name:     "invalid query: no select",
			query:    Query{Clauses: []string{"FROM", "users.csv"}},
			hasError: true,
		},
		{
			name:     "invalid query: empty query",
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

func TestQuery_GetFile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		query    Query
		want     string
		hasError bool
	}{
		{
			name:     "valid query with filename",
			query:    Query{Clauses: []string{"SELECT", "name", "FROM", "users.csv"}},
			want:     "users.csv",
			hasError: false,
		},
		{
			name:     "failed: no FROM",
			query:    Query{Clauses: []string{"SELECT", "name", "users.csv"}},
			want:     "",
			hasError: true,
		},
		{
			name:     "failed: no csvfile",
			query:    Query{Clauses: []string{"SELECT", "name", "FROM"}},
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

func TestQuery_GetSelect(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  []string
	}{
		{
			name:  "select single column",
			query: Query{Clauses: []string{"SELECT", "name", "FROM", "users.csv"}},
			want:  []string{"name"},
		},
		{
			name:  "select multiple columns",
			query: Query{Clauses: []string{"SELECT", "id", ",", "name", "FROM", "users.csv"}},
			want:  []string{"id", "name"},
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

func TestQuery_GetWhere(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  [][]string
	}{
		{
			name: "query with where clause",
			query: Query{Clauses: []string{
				"SELECT", "id", "FROM", "users.csv", "WHERE", "id", "=", "3",
			}},
			want: [][]string{
				{"id", "=", "3"},
			},
		},
		{
			name: "query without where clause",
			query: Query{Clauses: []string{
				"SELECT", "id", "FROM", "users.csv",
			}},
			want: [][]string{},
		},
		{
			name: "query with order clause",
			query: Query{Clauses: []string{
				"SELECT", "id", "FROM", "users.csv",
				"WHERE", "id", "=", "3", "ORDER", "BY", "id", "DESC",
			}},
			want: [][]string{
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

func TestQuery_GetOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  []string
	}{
		{
			name: "query with order by desc",
			query: Query{Clauses: []string{
				"SELECT", "id", "FROM", "users.csv", "ORDER", "BY", "id", "DESC",
			}},
			want: []string{
				"id", "DESC",
			},
		},
		{
			name: "query with order by",
			query: Query{Clauses: []string{
				"SELECT", "id", "FROM", "users.csv", "ORDER", "BY", "id",
			}},
			want: []string{
				"id", "ASC",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.query.GetOrder()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestQuery_GetLimit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  string
	}{
		{
			name: "query with limit",
			query: Query{Clauses: []string{
				"SELECT", "id", "FROM", "users.csv", "LIMIT", "3",
			}},
			want: "3",
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
