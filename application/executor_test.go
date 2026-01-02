package application

import (
	"errors"
	"testing"

	"github.com/ritarock/quickquery/model"
	"github.com/ritarock/quickquery/testing/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestExecutor_Execute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		query       string
		setupMocks  func(*mock.MockCSVReader, *mock.MockQueryExecutor)
		wantColumns []string
		wantRows    [][]string
		hasError    bool
	}{
		{
			name:  "pass",
			query: "SELECT id, name FROM test.csv",
			setupMocks: func(csvReader *mock.MockCSVReader, queryExecutor *mock.MockQueryExecutor) {
				csvReader.EXPECT().
					Read("test.csv").
					Return([]string{"id", "name"}, [][]string{
						{"1", "name1"},
						{"2", "name2"},
					}, nil)
				queryExecutor.EXPECT().
					CreateTable("test", []string{"id", "name"}).
					Return(nil)
				queryExecutor.EXPECT().
					InsertRows("test", []string{"id", "name"}, [][]string{
						{"1", "name1"},
						{"2", "name2"},
					}).
					Return(nil)
				queryExecutor.EXPECT().
					Query("SELECT id, name FROM test").
					Return([]string{"id", "name"}, [][]string{
						{"1", "name1"},
						{"2", "name2"},
					}, nil)
				queryExecutor.EXPECT().Close().Return(nil)
			},
			wantColumns: []string{"id", "name"},
			wantRows: [][]string{
				{"1", "name1"},
				{"2", "name2"},
			},
			hasError: false,
		},
		{
			name:  "failed: csv reader error",
			query: "SELECT id, name FROM test.csv",
			setupMocks: func(csvReader *mock.MockCSVReader, queryExecutor *mock.MockQueryExecutor) {
				csvReader.EXPECT().
					Read("test.csv").
					Return(nil, nil, errors.New("failed to read csv"))
				queryExecutor.EXPECT().Close().Return(nil)
			},
			hasError: true,
		},
		{
			name:  "failed: create table error",
			query: "SELECT id, name FROM test.csv",
			setupMocks: func(csvReader *mock.MockCSVReader, queryExecutor *mock.MockQueryExecutor) {
				csvReader.EXPECT().
					Read("test.csv").
					Return([]string{"id", "name"}, [][]string{{"1", "name1"}}, nil)
				queryExecutor.EXPECT().
					CreateTable("test", []string{"id", "name"}).
					Return(errors.New("failed to create table"))
				queryExecutor.EXPECT().Close().Return(nil)
			},
			hasError: true,
		},
		{
			name:  "faild: insert rows error",
			query: "SELECT * FROM test.csv",
			setupMocks: func(csvReader *mock.MockCSVReader, queryExecutor *mock.MockQueryExecutor) {
				csvReader.EXPECT().
					Read("test.csv").
					Return([]string{"id", "name"}, [][]string{{"1", "name1"}}, nil)
				queryExecutor.EXPECT().
					CreateTable("test", []string{"id", "name"}).
					Return(nil)
				queryExecutor.EXPECT().
					InsertRows("test", []string{"id", "name"}, [][]string{{"1", "name1"}}).
					Return(errors.New("failed to insert rows"))
				queryExecutor.EXPECT().Close().Return(nil)
			},
			hasError: true,
		},
		{
			name:  "failed: query error",
			query: "SELECT * FROM test.csv",
			setupMocks: func(csvReader *mock.MockCSVReader, queryExecutor *mock.MockQueryExecutor) {
				csvReader.EXPECT().
					Read("test.csv").
					Return([]string{"id", "name"}, [][]string{{"1", "name1"}}, nil)
				queryExecutor.EXPECT().
					CreateTable("test", []string{"id", "name"}).
					Return(nil)
				queryExecutor.EXPECT().
					InsertRows("test", []string{"id", "name"}, [][]string{{"1", "name1"}}).
					Return(nil)
				queryExecutor.EXPECT().
					Query("SELECT * FROM test").
					Return(nil, nil, errors.New("failed to execute query"))
				queryExecutor.EXPECT().Close().Return(nil)
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockCSVReader := mock.NewMockCSVReader(ctrl)
			mockQueryExecutor := mock.NewMockQueryExecutor(ctrl)

			test.setupMocks(mockCSVReader, mockQueryExecutor)

			executor, err := NewExecutor(mockCSVReader, mockQueryExecutor)
			assert.NoError(t, err)

			query, err := model.NewQuery(test.query)
			assert.NoError(t, err)

			got, err := executor.Execute(query)

			if test.hasError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NoError(t, err)
				assert.Equal(t, test.wantColumns, got.Columns())
				assert.Equal(t, test.wantRows, got.Rows())
			}
		})
	}
}
