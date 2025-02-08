package usecase

import (
	"quickquery/domain/entity"
	"strings"
)

type QueryExecutor struct {
	csvReader entity.CSVReader
}

type Result struct {
	Headers []string
	Rows    [][]string
}

func NewQueryExecutor(csvReader entity.CSVReader) *QueryExecutor {
	return &QueryExecutor{
		csvReader: csvReader,
	}
}

func (e *QueryExecutor) Execute(query entity.Query) (*Result, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}

	filename, err := query.GetFile()
	if err != nil {
		return nil, err
	}

	records, err := e.csvReader.ReadCSV(filename)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return &Result{}, nil
	}

	selectedColumns := query.GetSelect()
	headers := records[0]
	var columnIndices []int

	if len(selectedColumns) == 1 && selectedColumns[0] == "*" {
		columnIndices = make([]int, len(headers))
		for i := range headers {
			columnIndices[i] = i
		}
	} else {
		for _, col := range selectedColumns {
			for i, header := range headers {
				if strings.EqualFold(col, header) {
					columnIndices = append(columnIndices, i)
					break
				}
			}
		}
	}

	whereConditions := query.GetWhere()
	var filteredRows [][]string
	for i := 1; i < len(records); i++ {
		if e.matchesWhereConditions(records[i], headers, whereConditions) {
			filteredRows = append(filteredRows, records[i])
		}
	}

	result := &Result{
		Headers: make([]string, len(columnIndices)),
		Rows:    make([][]string, len(filteredRows)),
	}

	for i, idx := range columnIndices {
		result.Headers[i] = headers[idx]
	}

	for i, row := range filteredRows {
		result.Rows[i] = make([]string, len(columnIndices))
		for j, idx := range columnIndices {
			result.Rows[i][j] = row[idx]
		}
	}

	return result, nil
}

func (e *QueryExecutor) matchesWhereConditions(row, headers, conditions []string) bool {
	if len(conditions) == 0 {
		return true
	}

	for i := 0; i < len(conditions); i += 3 {
		if i+2 >= len(conditions) {
			break
		}

		column := conditions[i]
		operator := conditions[i+1]
		value := conditions[i+2]

		var columnIndex int
		for j, header := range headers {
			if strings.EqualFold(column, header) {
				columnIndex = j
				break
			}
		}

		if operator == "=" && row[columnIndex] != value {
			return false
		}
	}

	return true
}
