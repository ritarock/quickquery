package csv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader_Read(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		setupFile   func(t *testing.T) string
		wantColumns []string
		wantRows    [][]string
		hasError    bool
	}{
		{
			name: "pass",
			setupFile: func(t *testing.T) string {
				tmpDir := t.TempDir()
				tmpFile := filepath.Join(tmpDir, "test.csv")
				content := "id,name\n1,name1\n2,name2"
				if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return tmpFile
			},
			wantColumns: []string{"id", "name"},
			wantRows: [][]string{
				{"1", "name1"},
				{"2", "name2"},
			},
			hasError: false,
		},
		{
			name: "failed: file does not exist",
			setupFile: func(t *testing.T) string {
				return "/not/exist/file.csv"
			},
			hasError: true,
		},
		{
			name: "failed: csv file is empty",
			setupFile: func(t *testing.T) string {
				tmpDir := t.TempDir()
				tmpFile := filepath.Join(tmpDir, "test.csv")
				if err := os.WriteFile(tmpFile, []byte(""), 0644); err != nil {
					t.Fatalf("failed to create test file: %v", err)
				}
				return tmpFile
			},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			reader := NewReader()
			filePath := test.setupFile(t)
			gotColumns, gotRows, err := reader.Read(filePath)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.wantColumns, gotColumns)
				assert.Equal(t, test.wantRows, gotRows)
			}
		})
	}
}
