package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ritarock/quickquery/application"
)

type Executor struct {
	db *sql.DB
}

var _ application.QueryExecutor = (*Executor)(nil)

func NewExecutor() (*Executor, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite: %w", err)
	}

	return &Executor{db: db}, nil
}

func (e *Executor) CreateTable(tableName string, columns []string) error {
	query := createTableQuery(tableName, columns)
	_, err := e.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func (e *Executor) InsertRows(tableName string, columns []string, rows [][]string) error {
	if len(rows) == 0 {
		return nil
	}

	query := insertRowsQuery(tableName, columns)

	stmt, err := e.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare insert: %w", err)
	}
	defer stmt.Close()

	for _, row := range rows {
		args := make([]any, len(row))
		for i, v := range row {
			args[i] = v
		}
		if _, err := stmt.Exec(args...); err != nil {
			return fmt.Errorf("failed to insert row: %w", err)
		}
	}

	return nil
}

func (e *Executor) Query(query string) ([]string, [][]string, error) {
	rows, err := e.db.Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get columns: %w", err)
	}

	var results [][]string
	for rows.Next() {
		values := make([]any, len(columns))
		valuePtrs := make([]any, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}

		row := make([]string, len(columns))
		for i, v := range values {
			if v == nil {
				row[i] = "NULL"
			} else {
				row[i] = fmt.Sprintf("%v", v)
			}
		}

		results = append(results, row)
	}

	return columns, results, nil
}

func (e *Executor) Close() error {
	return e.db.Close()
}

func createTableQuery(tableName string, columns []string) string {
	columnDefs := make([]string, len(columns))
	for i, col := range columns {
		columnDefs[i] = fmt.Sprintf(`"%s" TEXT`, col)
	}

	return fmt.Sprintf(
		`CREATE TABLE "%s" (%s)`,
		tableName,
		strings.Join(columnDefs, ", "),
	)
}

func insertRowsQuery(tableName string, columns []string) string {
	placeholders := make([]string, len(columns))
	for i := range columns {
		placeholders[i] = "?"
	}

	return fmt.Sprintf(
		`INSERT INTO "%s" (%s) VALUES (%s)`,
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

}
