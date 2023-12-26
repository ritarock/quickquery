package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_validate(t *testing.T) {
	tests := []struct {
		name     string
		query    Query
		hasError bool
	}{
		{
			name:     "pass",
			query:    Query{"SELECT", "*", "FROM", "./sample.csv"},
			hasError: false,
		},
		{
			name:     "failed: syntax error",
			query:    Query{"SELEC", "*", "FROM", "./sample.csv"},
			hasError: true,
		},
	}

	for _, test := range tests {
		err := test.query.validate()
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
func TestTable_toMap(t *testing.T) {
	tests := []struct {
		table Table
		want  Mapper
	}{
		{
			table: Table{
				{"id", "name", "user"},
				{"1", "name1", "user1"},
				{"2", "name2", "user2"},
				{"3", "name3", "user3"},
			},
			want: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
		},
	}

	for _, test := range tests {
		got := test.table.toMap()
		assert.Equal(t, test.want, got)
	}
}

func TestQuery_getFileName(t *testing.T) {
	tests := []struct {
		name     string
		query    Query
		want     string
		hasError bool
	}{
		{
			name:     "pass",
			query:    Query{"SELECT", "*", "FROM", "./sample.csv"},
			want:     "./sample.csv",
			hasError: false,
		},
		{
			name:     "failed: not exist file name",
			query:    Query{"SELECT", "*", "FROM", ""},
			hasError: true,
		},
	}

	for _, test := range tests {
		file, err := test.query.getFileName()
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, file, test.want)
		}
	}
}

func TestQuery_getSelect(t *testing.T) {
	tests := []struct {
		query Query
		want  SelectClause
	}{
		{
			query: Query{"SELECT", "id", "user", "name", "FROM", "./sample.csv"},
			want:  SelectClause{"id", "user", "name"},
		},
	}

	for _, test := range tests {
		got := test.query.getSelect()
		assert.Equal(t, test.want, got)
	}
}
func TestSelectClause_isAll(t *testing.T) {
	tests := []struct {
		selectClause SelectClause
		want         bool
	}{
		{
			selectClause: SelectClause{"col"},
			want:         false,
		},
		{
			selectClause: SelectClause{"*"},
			want:         true,
		},
	}

	for _, test := range tests {
		got := test.selectClause.isAll()
		assert.Equal(t, test.want, got)
	}
}
func TestMapper_bySelect(t *testing.T) {
	tests := []struct {
		mapper       Mapper
		selectClause SelectClause
		want         Table
	}{
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			selectClause: SelectClause{
				"id", "name", "user",
			},
			want: Table{
				{"id", "name", "user"},
				{"1", "name1", "user1"},
				{"2", "name2", "user2"},
				{"3", "name3", "user3"},
			},
		},
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			selectClause: SelectClause{
				"id", "user",
			},
			want: Table{
				{"id", "user"},
				{"1", "user1"},
				{"2", "user2"},
				{"3", "user3"},
			},
		},
	}

	for _, test := range tests {
		got := test.mapper.bySelect(test.selectClause)
		assert.Equal(t, test.want, got)
	}
}

func Test_argToQuery(t *testing.T) {
	tests := []struct {
		arg  string
		want Query
	}{
		{
			arg:  "select * from ./sample.csv",
			want: Query{"SELECT", "*", "FROM", "./sample.csv"},
		},
	}

	for _, test := range tests {
		got := argToQuery(test.arg)
		assert.Equal(t, test.want, got)
	}
}
