package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgToQuery(t *testing.T) {
	tests := []struct {
		arg  string
		want Query
	}{
		{
			arg: "select * from ./sample.csv where id = 1 and user = user1",
			want: Query{
				"SELECT", "*",
				"FROM", "./sample.csv",
				"WHERE", "id", "=", "1",
				"AND", "user", "=", "user1",
			},
		},
	}

	for _, test := range tests {
		got := ArgToQuery(test.arg)
		assert.Equal(t, test.want, got)
	}
}

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
		err := test.query.Validate()
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestQuery_GetFileName(t *testing.T) {
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
		file, err := test.query.GetFileName()
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, file, test.want)
		}
	}
}
