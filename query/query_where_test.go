package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_GetWhere(t *testing.T) {
	tests := []struct {
		query     Query
		want      WhereQ
		wantFound bool
	}{
		{
			query: Query{
				"WHERE", "id", "=", "2",
			},
			want: WhereQ{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id=2",
			},
			want: WhereQ{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id=", "2",
			},
			want: WhereQ{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "=2",
			},
			want: WhereQ{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">", "2",
			},
			want: WhereQ{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>2",
			},
			want: WhereQ{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>", "2",
			},
			want: WhereQ{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">2",
			},
			want: WhereQ{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<", "2",
			},
			want: WhereQ{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<2",
			},
			want: WhereQ{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<", "2",
			},
			want: WhereQ{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<2",
			},
			want: WhereQ{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">=", "2",
			},
			want: WhereQ{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>=2",
			},
			want: WhereQ{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>=", "2",
			},
			want: WhereQ{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">=2",
			},
			want: WhereQ{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<=", "2",
			},
			want: WhereQ{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<=2",
			},
			want: WhereQ{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<=", "2",
			},
			want: WhereQ{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<=2",
			},
			want: WhereQ{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
	}

	for _, test := range tests {
		got, found := test.query.GetWhere()
		assert.Equal(t, test.wantFound, found)
		assert.Equal(t, test.want, got)
	}
}

func Test_makeCondition(t *testing.T) {
	tests := []struct {
		whereQ    WhereQ
		condition string
		operator  string
		want      WhereQ
	}{
		{
			whereQ:    WhereQ{},
			condition: "id=1",
			operator:  "=",
			want: WhereQ{
				{"id", "=", "1"},
			},
		},
		{
			whereQ: WhereQ{
				{"id", "=", "1"},
			},
			condition: "user>=user1",
			operator:  ">=",
			want: WhereQ{
				{"id", "=", "1"},
				{"user", ">=", "user1"},
			},
		},
	}

	for _, test := range tests {
		got := makeCondition(test.whereQ, test.condition, test.operator)
		assert.Equal(t, test.want, got)
	}
}
