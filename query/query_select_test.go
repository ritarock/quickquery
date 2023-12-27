package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetSelect(t *testing.T) {
	tests := []struct {
		query Query
		want  SelectQ
	}{
		{
			query: Query{"SELECT", "id", "user", "name", "FROM", "./sample.csv"},
			want:  SelectQ{"id", "user", "name"},
		},
	}

	for _, test := range tests {
		got := test.query.GetSelect()
		assert.Equal(t, test.want, got)
	}
}

func TestSelectQ_IsAll(t *testing.T) {
	tests := []struct {
		selectClause SelectQ
		want         bool
	}{
		{
			selectClause: SelectQ{"col"},
			want:         false,
		},
		{
			selectClause: SelectQ{"*"},
			want:         true,
		},
	}

	for _, test := range tests {
		got := test.selectClause.IsAll()
		assert.Equal(t, test.want, got)
	}
}
