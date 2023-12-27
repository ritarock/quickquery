package query

import (
	"io"
	"os"
	"strings"
)

type Table [][]string

func ReadTable(path string) (Table, error) {
	f, err := readFile(path)
	if err != nil {
		return nil, err
	}

	var table Table
	lines := strings.Split(f, "\n")
	for _, cols := range lines {
		columns := []string{}
		for _, col := range strings.Split(cols, ",") {
			tmp := strings.TrimSpace(col)
			columns = append(columns, tmp)
		}
		// ignore line break
		if len(columns) == 1 && columns[0] == "" {
			continue
		}
		table = append(table, columns)
	}

	return table, nil
}

func readFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
