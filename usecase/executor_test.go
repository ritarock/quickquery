package usecase

import (
	"quickquery/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCSVServer struct {
	data [][]string
	err  error
}

func (m *mockCSVServer) ReadCSV(filename string) ([][]string, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.data, nil
}

func TestQueryExecutor_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		query     entity.Query
		mockData  [][]string
		mockError error
		want      *Result
		hasError  bool
	}{
		{
			name: "select all columns",
			query: entity.Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
			}},
			mockData: [][]string{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			mockError: nil,
			want: &Result{
				Headers: []string{"id", "team_id", "name"},
				Rows: [][]string{
					{"1", "1", "name1"},
					{"2", "2", "name2"},
					{"3", "3", "name3"},
				},
			},
			hasError: false,
		},
		{
			name: "select specific columns",
			query: entity.Query{Clauses: []string{
				"SELECT", "id", "name", "FROM", "users.csv",
			}},
			mockData: [][]string{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			mockError: nil,
			want: &Result{
				Headers: []string{"id", "name"},
				Rows: [][]string{
					{"1", "name1"},
					{"2", "name2"},
					{"3", "name3"},
				},
			},
			hasError: false,
		},
		{
			name: "query with where",
			query: entity.Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
				"WHERE", "id", ">=", "2",
			}},
			mockData: [][]string{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			mockError: nil,
			want: &Result{
				Headers: []string{"id", "team_id", "name"},
				Rows: [][]string{
					{"2", "2", "name2"},
					{"3", "3", "name3"},
				},
			},
			hasError: false,
		},
		{
			name: "query with order",
			query: entity.Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
				"ORDER", "BY", "id", "DESC",
			}},
			mockData: [][]string{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			mockError: nil,
			want: &Result{
				Headers: []string{"id", "team_id", "name"},
				Rows: [][]string{
					{"3", "3", "name3"},
					{"2", "2", "name2"},
					{"1", "1", "name1"},
				},
			},
			hasError: false,
		},
		{
			name: "query with limit",
			query: entity.Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
				"LIMIT", "2",
			}},
			mockData: [][]string{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			mockError: nil,
			want: &Result{
				Headers: []string{"id", "team_id", "name"},
				Rows: [][]string{
					{"1", "1", "name1"},
					{"2", "2", "name2"},
				},
			},
			hasError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockReader := &mockCSVServer{
				data: test.mockData,
				err:  test.mockError,
			}
			executor := NewQueryExecutor(mockReader)

			got, err := executor.Execute(test.query)

			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
