package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecords_SortRows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		records Records
		order   Order
		want    Records
	}{
		{
			name: "query with ORDER BY ASC",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"5", "5", "name5"},
				{"2", "2", "name2"},
			},
			order: Order{
				column:    "id",
				condition: "ASC",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"5", "5", "name5"},
			},
		},
		{
			name: "query with ORDER BY DESC",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"5", "5", "name5"},
				{"2", "2", "name2"},
			},
			order: Order{
				column:    "id",
				condition: "DESC",
			},
			want: Records{
				{"id", "team_id", "name"},
				{"5", "5", "name5"},
				{"2", "2", "name2"},
				{"1", "1", "name1"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.records.SortRows(test.order)
			assert.Equal(t, test.want, test.records)
		})
	}
}
