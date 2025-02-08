package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		query Query
		want  Order
	}{
		{
			name: "query with order",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv", "ORDER", "BY", "id",
			}},
			want: Order{
				column:    "id",
				condition: "ASC",
			},
		},
		{
			name: "query with order by desc",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv", "ORDER", "BY", "id", "DESC",
			}},
			want: Order{
				column:    "id",
				condition: "DESC",
			},
		},
		{
			name: "query with order by asc",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv", "ORDER", "BY", "id", "ASC",
			}},
			want: Order{
				column:    "id",
				condition: "ASC",
			},
		},
		{
			name: "query without order",
			query: Query{Clauses: []string{
				"SELECT", "*", "FROM", "users.csv",
			}},
			want: Order{
				column:    "",
				condition: "",
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

func TestOrder_IsOrder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		order Order
		want  bool
	}{
		{
			name: "query with order",
			order: Order{
				column:    "id",
				condition: "ASC",
			},
			want: true,
		},
		{
			name: "query without order",
			order: Order{
				column:    "",
				condition: "",
			},
			want: false,
		},
	}

	for _, test := range tests {
		got := test.order.IsOrder()
		assert.Equal(t, test.want, got)
	}
}
