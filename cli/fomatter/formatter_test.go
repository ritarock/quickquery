package fomatter

import (
	"testing"

	"github.com/ritarock/quickquery/model"
	"github.com/stretchr/testify/assert"
)

func TestFormatter_Format(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		columns []string
		rows    [][]string
		want    string
	}{
		{
			name:    "multiple rows",
			columns: []string{"id", "name"},
			rows: [][]string{
				{"1", "name1"},
				{"2", "name2"},
			},
			want: "| id | name  |\n" +
				"+----+-------+\n" +
				"| 1  | name1 |\n" +
				"| 2  | name2 |\n" +
				"(2 rows)",
		},
		{
			name:    "single row",
			columns: []string{"id", "name"},
			rows: [][]string{
				{"1", "name1"},
			},
			want: "| id | name  |\n" +
				"+----+-------+\n" +
				"| 1  | name1 |\n" +
				"(1 rows)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := NewFormatter()
			result := model.NewQueryResult(test.columns, test.rows)
			got := f.Format(result)
			assert.Equal(t, test.want, got)
		})
	}
}
