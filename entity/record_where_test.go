package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecords_filterRows(t *testing.T) {
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
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			condition: []string{
				"id", "=", "2",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"2", "2", "name2"},
			},
		},
		{
			name: "match !=",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			condition: []string{
				"id", "!=", "2",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"3", "3", "name3"},
			},
		},
		{
			name: "match <",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			condition: []string{
				"id", "<", "2",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
			},
		},
		{
			name: "match >",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			condition: []string{
				"id", ">", "2",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"3", "3", "name3"},
			},
		},
		{
			name: "match <=",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			condition: []string{
				"id", "<=", "2",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
			},
		},
		{
			name: "match >=",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			condition: []string{
				"id", ">=", "2",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.records.filterRows(test.condition)
			assert.Equal(t, test.want, test.records)
		})
	}
}

func TestRecords_FilterRows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		records Records
		where   Where
		want    Records
	}{
		{
			name: "query with where",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			where: Where{
				{"id", ">=", "2"},
				{"team_id", "=", "3"},
			},
			want: Records{
				{"id", "team_id", "name"},
				{"3", "3", "name3"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.records.FilterRows(test.where)
			assert.Equal(t, test.want, test.records)
		})
	}
}
