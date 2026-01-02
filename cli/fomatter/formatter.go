package fomatter

import (
	"fmt"
	"strings"

	"github.com/ritarock/quickquery/model"
)

type Formatter struct{}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Format(result *model.QueryResult) string {
	columns := result.Columns()
	rows := result.Rows()

	if len(columns) == 0 {
		return "no results"
	}

	widths := make([]int, len(columns))
	for i, col := range columns {
		widths[i] = len(col)
	}
	for _, row := range rows {
		for i, val := range row {
			if len(val) > widths[i] {
				widths[i] = len(val)
			}
		}
	}

	var sb strings.Builder
	sb.WriteString(f.formatRow(columns, widths))
	sb.WriteString("\n")

	sb.WriteString(f.formatSeparator(widths))
	sb.WriteString("\n")

	for _, row := range rows {
		sb.WriteString(f.formatRow(row, widths))
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("(%d rows)", result.RowCount()))

	return sb.String()
}

func (f *Formatter) formatRow(values []string, widths []int) string {
	cells := make([]string, len(values))
	for i, val := range values {
		cells[i] = fmt.Sprintf("%-*s", widths[i], val)
	}
	return "| " + strings.Join(cells, " | ") + " |"
}

func (f *Formatter) formatSeparator(widths []int) string {
	parts := make([]string, len(widths))
	for i, w := range widths {
		parts[i] = strings.Repeat("-", w)
	}
	return "+-" + strings.Join(parts, "-+-") + "-+"
}
