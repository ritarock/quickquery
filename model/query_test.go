package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractTableName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		rawQuery  string
		wantQuery string
		wantTable string
		hasError  bool
	}{
		{
			name:      "pass",
			rawQuery:  "SELECT * FROM test.csv",
			wantQuery: "SELECT * FROM test",
			wantTable: "test",
			hasError:  false,
		},
		{
			name:      "pass: lowercase from",
			rawQuery:  "select * from test.csv",
			wantQuery: "select * from test",
			wantTable: "test",
			hasError:  false,
		},
		{
			name:      "pass: mixed case from",
			rawQuery:  "Select * From test.csv",
			wantQuery: "Select * From test",
			wantTable: "test",
			hasError:  false,
		},
		{
			name:      "pass: extra spaces",
			rawQuery:  "SELECT   *   FROM   test.csv",
			wantQuery: "SELECT * FROM test",
			wantTable: "test",
			hasError:  false,
		},
		{
			name:     "failed: without csv suffix",
			rawQuery: "SELECT * FROM",
			hasError: true,
		},
		{
			name:     "failed: no from clause",
			rawQuery: "SELECT * test.csv",
			hasError: true,
		},
		{
			name:     "failed: from without table name",
			rawQuery: "SELECT * from",
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			gotQuery, gotTable, err := extractTableName(test.rawQuery)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.wantQuery, gotQuery)
				assert.Equal(t, test.wantTable, gotTable)
			}
		})
	}
}
