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

	readCsv, err := e.csvReader.ReadCSV(filename)
	var records entity.Records = readCsv
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
	if whereConditions.IsWhere() {
		records.FilterRows(whereConditions)
	}

	sortConditions := query.GetOrder()
	if sortConditions.IsOrder() {
		records.SortRows(sortConditions)

	}

	limitConditions := query.GetLimit()
	if limitConditions.IsLimit() {
		records.LimitRows(limitConditions)
	}

	result := &Result{
		Headers: make([]string, len(columnIndices)),
		Rows:    make([][]string, len(records)-1),
	}

	for i, idx := range columnIndices {
		result.Headers[i] = headers[idx]
	}

	for i, row := range records[1:] {
		result.Rows[i] = make([]string, len(columnIndices))
		for j, idx := range columnIndices {
			result.Rows[i][j] = row[idx]
		}
	}

	return result, nil
}
