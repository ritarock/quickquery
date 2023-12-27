package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_getWhere(t *testing.T) {
	tests := []struct {
		query     Query
		want      WhereClause
		wantFound bool
	}{
		{
			query: Query{
				"WHERE", "id", "=", "2",
			},
			want: WhereClause{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id=2",
			},
			want: WhereClause{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id=", "2",
			},
			want: WhereClause{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "=2",
			},
			want: WhereClause{
				{"id", "=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">", "2",
			},
			want: WhereClause{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>2",
			},
			want: WhereClause{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>", "2",
			},
			want: WhereClause{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">2",
			},
			want: WhereClause{
				{"id", ">", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<", "2",
			},
			want: WhereClause{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<2",
			},
			want: WhereClause{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<", "2",
			},
			want: WhereClause{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<2",
			},
			want: WhereClause{
				{"id", "<", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">=", "2",
			},
			want: WhereClause{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>=2",
			},
			want: WhereClause{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id>=", "2",
			},
			want: WhereClause{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", ">=2",
			},
			want: WhereClause{
				{"id", ">=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<=", "2",
			},
			want: WhereClause{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<=2",
			},
			want: WhereClause{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id<=", "2",
			},
			want: WhereClause{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
		{
			query: Query{
				"WHERE", "id", "<=2",
			},
			want: WhereClause{
				{"id", "<=", "2"},
			},
			wantFound: true,
		},
	}

	for _, test := range tests {
		got, found := test.query.getWhere()
		assert.Equal(t, test.wantFound, found)
		assert.Equal(t, test.want, got)
	}
}

func TestMapper_byWhere(t *testing.T) {
	tests := []struct {
		mapper      Mapper
		whereClause WhereClause
		want        Mapper
	}{
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			whereClause: WhereClause{
				{
					"id", "=", "2",
				},
			},
			want: Mapper{
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
			},
		},
	}

	for _, test := range tests {
		got := test.mapper.byWhere(test.whereClause)
		assert.Equal(t, test.want, got)
	}
}
