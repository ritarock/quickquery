package model

import (
	"errors"
	"strings"
)

type Query struct {
	processedQuery string
	tableName      string
}

func NewQuery(rawQuery string) (*Query, error) {
	if rawQuery == "" {
		return nil, errors.New("query cannot be empty")
	}

	processedQuery, tableName, err := extractTableName(rawQuery)

	if err != nil {
		return nil, err
	}

	return &Query{processedQuery: processedQuery, tableName: tableName}, nil
}

func (q *Query) ProcessedQuery() string {
	return q.processedQuery
}

func (q *Query) TableName() string {
	return q.tableName
}

func (q *Query) CSVFileName() string {
	return q.tableName + ".csv"
}

func extractTableName(raw string) (string, string, error) {
	tmp := strings.TrimSpace(raw)
	clauses := strings.Fields(tmp)
	var tableName string

	var index int
	for i, v := range clauses {
		if strings.EqualFold(v, "FROM") && i+1 < len(clauses) {
			tableName = clauses[i+1]
			index = i + 1
			break
		}
	}
	if tableName == "" {
		return "", "", errors.New("could not extract table name from query")
	}

	tableName = strings.TrimSuffix(tableName, ".csv")
	clauses[index] = tableName

	return strings.Join(clauses, " "), tableName, nil
}
