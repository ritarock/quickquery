package sqlite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createTableQuery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		tableName string
		columns   []string
		want      string
	}{
		{
			name:      "pass",
			tableName: "test",
			columns:   []string{"id", "name"},
			want:      `CREATE TABLE "test" ("id" TEXT, "name" TEXT)`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := createTableQuery(test.tableName, test.columns)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_insertRowsQuery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		tableName string
		columns   []string
		want      string
	}{
		{
			name:      "pass",
			tableName: "test",
			columns:   []string{"id", "name"},
			want:      `INSERT INTO "test" (id, name) VALUES (?, ?)`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := insertRowsQuery(test.tableName, test.columns)
			assert.Equal(t, test.want, got)
		})
	}
}
