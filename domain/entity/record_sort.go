package entity

import (
	"sort"
	"strings"
)

func (r *Records) SortRows(order Order) {
	columnIndex := -1
	for i, header := range (*r)[0] {
		if strings.EqualFold(order.column, header) {
			columnIndex = i
			break
		}
	}

	if columnIndex == -1 {
		return
	}

	isAsc := order.condition == "ASC"
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
