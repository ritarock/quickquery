package model

type QueryResult struct {
	columns []string
	rows    [][]string
}

func NewQueryResult(columns []string, rows [][]string) *QueryResult {
	return &QueryResult{columns: columns, rows: rows}
}

func (r *QueryResult) Columns() []string {
	return r.columns
}

func (r *QueryResult) Rows() [][]string {
	return r.rows
}

func (r *QueryResult) RowCount() int {
	return len(r.rows)
}
