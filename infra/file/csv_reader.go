package file

import (
	"encoding/csv"
	"os"
	"quickquery/entity"
)

type CSVReader struct{}

var _ entity.CSVReader = (*CSVReader)(nil)

func NewCSVReader() *CSVReader {
	return &CSVReader{}
}

func (r *CSVReader) ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
