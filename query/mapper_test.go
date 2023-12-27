package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_ToMap(t *testing.T) {
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
		got := test.table.ToMap()
		assert.Equal(t, test.want, got)
	}
}

func TestMapper_adaptedWhere(t *testing.T) {
	tests := []struct {
		mapper Mapper
		whereQ WhereQ
		want   Mapper
	}{
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			whereQ: WhereQ{
				{"id", "=", "2"},
			},
			want: Mapper{
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
			},
		},
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			whereQ: WhereQ{
				{"id", ">=", "2"},
			},
			want: Mapper{
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
		},
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			whereQ: WhereQ{
				{"id", ">=", "2"},
				{"id", "=", "3"},
			},
			want: Mapper{
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
		},
	}

	for _, test := range tests {
		got := test.mapper.adaptedWhere(test.whereQ)
		assert.Equal(t, test.want, got)
	}
}

func TestMapper_applyCondition(t *testing.T) {
	tests := []struct {
		mapper    Mapper
		key       string
		condition string
		value     string
		want      Mapper
	}{
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			key:       "id",
			condition: ">=",
			value:     "2",
			want: Mapper{
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
		},
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			key:       "id",
			condition: "<=",
			value:     "2",
			want: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
			},
		},
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			key:       "id",
			condition: "=",
			value:     "2",
			want: Mapper{
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
			},
		},
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			key:       "id",
			condition: ">",
			value:     "2",
			want: Mapper{
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
		},
		{
			mapper: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
				map[string]string{"id": "2", "name": "name2", "user": "user2"},
				map[string]string{"id": "3", "name": "name3", "user": "user3"},
			},
			key:       "id",
			condition: "<",
			value:     "2",
			want: Mapper{
				map[string]string{"id": "1", "name": "name1", "user": "user1"},
			},
		},
	}

	for _, test := range tests {
		got := test.mapper.applyCondition(test.key, test.condition, test.value)
		assert.Equal(t, test.want, got)
	}
}
