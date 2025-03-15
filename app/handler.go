package app

import (
	"errors"
	"fmt"
	"os"
	"quickquery/entity"
	"quickquery/infra/file"
	"strings"
	"text/tabwriter"
)

type Handler struct {
	queryExecutor QueryExecutor
}

func NewHandler() *Handler {
	csvReader := file.NewCSVReader()
	queryExecutor := NewQueryExecutor(csvReader)
	return &Handler{queryExecutor: *queryExecutor}
}

func (h *Handler) Run(args []string) error {
	if err := validateArgs(args); err != nil {
		return err
	}
	q := entity.NewQuery(args[0])
	result, err := h.queryExecutor.Execute(q)
	if err != nil {
		return err
	}
	h.displayResults(result)
	return nil
}

func (h *Handler) displayResults(result *Result) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, strings.Join(result.Headers, "\t"))

	separators := make([]string, len(result.Headers))
	for i := range separators {
		separators[i] = strings.Repeat("-", len(result.Headers[i]))
	}
	fmt.Fprintln(w, strings.Join(separators, "\t"))

	for _, row := range result.Rows {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}

	w.Flush()
}

func validateArgs(args []string) error {
	if len(args) == 0 || args[0] == "" {
		return errors.New("not enough args")
	}
	if len(args) >= 2 {
		return errors.New("too many args")
	}

	return nil
}
