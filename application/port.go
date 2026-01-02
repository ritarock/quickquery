package application

type CSVReader interface {
	Read(filePath string) ([]string, [][]string, error)
}

type QueryExecutor interface {
	CreateTable(tableName string, columns []string) error
	InsertRows(tableName string, columns []string, rows [][]string) error
	Query(query string) ([]string, [][]string, error)
	Close() error
}
