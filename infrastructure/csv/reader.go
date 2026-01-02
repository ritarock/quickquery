package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read(filePath string) ([]string, [][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open csv file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read csv: %w", err)
	}

	if len(records) == 0 {
		return nil, nil, errors.New("csv file is empty")
	}

	columns := records[0]
	rows := records[1:]

	return columns, rows, nil
}
