package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecords_LimitRows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		records Records
		limit   Limit
		want    Records
	}{
		{
			name: "query with LIMIT",
			records: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
				{"3", "3", "name3"},
			},
			limit: Limit(2),
			want: Records{
				{"id", "team_id", "name"},
				{"1", "1", "name1"},
				{"2", "2", "name2"},
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
