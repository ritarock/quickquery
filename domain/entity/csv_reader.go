package entity

type CSVReader interface {
	ReadCSV(filename string) ([][]string, error)
}
