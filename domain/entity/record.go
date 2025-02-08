package entity

import (
	"sort"
	"strconv"
	"strings"
)

type Records [][]string

func (r *Records) FilterRows(condition []string) {
	if len(*r) == 0 {
		return
	}

	column, operator, value := condition[0], condition[1], condition[2]
	columnIndex := -1

	header := (*r)[0]
	for i, col := range header {
		if col == column {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		return
	}

	filtered := Records{header}

	for _, row := range (*r)[1:] {
		if len(row) <= columnIndex {
			continue
		}

		cellValue := row[columnIndex]
		match := false

		switch operator {
		case "=":
			match = cellValue == value
		case "!=":
			match = cellValue != value
		case "<", ">", ">=", "<=":
			cellNum, err1 := strconv.ParseFloat(cellValue, 64)
			valueNum, err2 := strconv.ParseFloat(value, 64)
			if err1 == nil && err2 == nil {
				switch operator {
				case "<":
					match = cellNum < valueNum
				case ">":
					match = cellNum > valueNum
				case "<=":
					match = cellNum <= valueNum
				case ">=":
					match = cellNum >= valueNum
				}
			}
		}

		if match {
			filtered = append(filtered, row)
		}
	}

	*r = filtered
}

func (r *Records) SortRows(column, order string) {
	columnIndex := -1
	for i, header := range (*r)[0] {
		if strings.EqualFold(column, header) {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		return
	}

	isAsc := order == "ASC"
	if isAsc {
		sorted := Records{(*r)[0]}
		rows := (*r)[1:]
		sort.Slice(rows, func(i, j int) bool {
			return rows[i][columnIndex] < rows[j][columnIndex]
		})
		sorted = append(sorted, rows...)
	} else {
		sorted := Records{(*r)[0]}
		rows := (*r)[1:]
		sort.Slice(rows, func(i, j int) bool {
			return rows[i][columnIndex] > rows[j][columnIndex]
		})
		sorted = append(sorted, rows...)
	}
}
