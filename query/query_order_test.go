package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetOrder(t *testing.T) {
	tests := []struct {
		query  Query
		header []string
		want   OrderQ
	}{
		{
			query: Query{
				"ORDER", "BY", "id",
			},
			header: []string{"id"},
			want: map[string]string{
				"id": "ASC",
			},
		},
		{
			query: Query{
				"ORDER", "BY", "id", "ASC",
			},
			header: []string{"id"},
			want: map[string]string{
				"id": "ASC",
			},
		},
		{
			query: Query{
				"ORDER", "BY", "id", "DESC",
			},
			header: []string{"id"},
			want: map[string]string{
				"id": "DESC",
			},
		},
		{
			query:  Query{},
			header: []string{"id"},
			want:   nil,
		},
		{
			query: Query{
				"ORDER", "BY", "nam", "DESC",
			},
			header: []string{"id", "name", "user"},
			want:   nil,
		},
	}

	for _, test := range tests {
		got := test.query.GetOrder(test.header)
		assert.Equal(t, test.want, got)
	}
}

func Test_existKeyInHeader(t *testing.T) {
	tests := []struct {
		key    string
		header []string
		want   bool
	}{
		{
			key:    "name",
			header: []string{"id", "name", "user"},
			want:   true,
		},
		{
			key:    "nam",
			header: []string{"id", "name", "user"},
			want:   false,
		},
	}

	for _, test := range tests {
		got := existKeyInHeader(test.key, test.header)
		assert.Equal(t, test.want, got)
	}
}
