package usecase

import (
	"quickquery/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCSVServer struct {
	data [][]string
	err  error
}

func (m *MockCSVServer) ReadCSV(filename string) ([][]string, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.data, nil
}

func TestQueryExecutor_Execute(t *testing.T) {
	t.Parallel()
	testData := [][]string{
		{"id", "name"},
		{"1", "name1"},
		{"2", "name2"},
		{"3", "name3"},
	}

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
			query: entity.Query{
				Clauses: []string{"SELECT", "*", "FROM", "users.csv"},
			},
			mockData:  testData,
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
			name: "select specific columns",
			query: entity.Query{
				Clauses: []string{"SELECT", "id", ",", "name", "FROM", "users.csv"},
			},
			mockData:  testData,
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockReader := &MockCSVServer{
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
