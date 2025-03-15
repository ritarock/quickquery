package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSVReader_ReadCSV(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile_")
	if err != nil {
		t.Fatalf("failed: create tmp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := []byte("id,name\n1,name1\n2,name2")

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatalf("failed: write tmp file: %v", err)
	}

	tests := []struct {
		name     string
		filename string
		want     [][]string
		hasError bool
	}{
		{
			name:     "succeed",
			filename: tmpFile.Name(),
			want: [][]string{
				{"id", "name"},
				{"1", "name1"},
				{"2", "name2"},
			},
			hasError: false,
		},
		{
			name:     "failed: non-existent file",
			filename: "nonexistent.csv",
			want:     nil,
			hasError: true,
		},
	}

	reader := NewCSVReader()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := reader.ReadCSV(test.filename)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
