package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Table [][]string

func readFile(path string) (Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var table Table
	lines := strings.Split(string(b), "\n")
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

func (t Table) Result() {
	for _, v := range t {
		fmt.Println(strings.Join(v, ", "))
	}
}
