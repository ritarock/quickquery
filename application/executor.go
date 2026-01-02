package application

import (
	"github.com/ritarock/quickquery/model"
)

type Executor struct {
	reader        CSVReader
	queryExecutor QueryExecutor
}

func NewExecutor(reader CSVReader, queryExecutor QueryExecutor) (*Executor, error) {
	return &Executor{
		reader:        reader,
		queryExecutor: queryExecutor,
	}, nil
}

func (e *Executor) Execute(query *model.Query) (*model.QueryResult, error) {
	defer e.queryExecutor.Close()

	columns, rows, err := e.reader.Read(query.CSVFileName())
	if err != nil {
		return nil, err
	}

	if err := e.queryExecutor.CreateTable(query.TableName(), columns); err != nil {
		return nil, err
	}
	if err := e.queryExecutor.InsertRows(query.TableName(), columns, rows); err != nil {
		return nil, err
	}

	resultColumns, resultRows, err := e.queryExecutor.Query(query.ProcessedQuery())
	if err != nil {
		return nil, err
	}

	return model.NewQueryResult(resultColumns, resultRows), nil
}
