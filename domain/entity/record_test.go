package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecord_FilterRows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		records   Records
		condition []string
		want      Records
	}{
		{
			name: "match =",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
			condition: []string{"id", "=", "2"},
			want: [][]string{
				{"id", "user_id", "name"},
				{"2", "user2", "name2"},
			},
		},
		{
			name: "match !=",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
			condition: []string{"id", "!=", "2"},
			want: [][]string{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"3", "user3", "name3"},
			},
		},
		{
			name: "match <",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
			condition: []string{"id", "<", "2"},
			want: [][]string{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
			},
		},
		{
			name: "match >",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
			condition: []string{"id", ">", "2"},
			want: [][]string{
				{"id", "user_id", "name"},
				{"3", "user3", "name3"},
			},
		},
		{
			name: "match <=",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
			condition: []string{"id", "<=", "2"},
			want: [][]string{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
			},
		},
		{
			name: "match >=",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
			condition: []string{"id", ">=", "2"},
			want: [][]string{
				{"id", "user_id", "name"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.records.FilterRows(test.condition)
			assert.Equal(t, test.want, test.records)
		})
	}
}

func TestRecords_SortRows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		records Records
		column  string
		order   string
		want    Records
	}{
		{
			name: "order by asc",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"3", "user3", "name3"},
				{"2", "user2", "name2"},
			},
			column: "id",
			order:  "ASC",
			want: [][]string{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
		},
		{
			name: "order by desc",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"3", "user3", "name3"},
				{"2", "user2", "name2"},
			},
			column: "id",
			order:  "DESC",
			want: [][]string{
				{"id", "user_id", "name"},
				{"3", "user3", "name3"},
				{"2", "user2", "name2"},
				{"1", "user1", "name1"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.records.SortRows(test.column, test.order)
			assert.Equal(t, test.want, test.records)
		})
	}
}

func TestRecords_LimitRows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		records Records
		limit   string
		want    Records
	}{
		{
			name: "order by asc",
			records: Records{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
				{"3", "user3", "name3"},
			},
			limit: "2",
			want: [][]string{
				{"id", "user_id", "name"},
				{"1", "user1", "name1"},
				{"2", "user2", "name2"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.records.LimitRows(test.limit)
			assert.Equal(t, test.want, test.records)
		})
	}
}
