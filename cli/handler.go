package cli

import (
	"errors"
	"fmt"

	"github.com/ritarock/quickquery/application"
	"github.com/ritarock/quickquery/cli/fomatter"
	"github.com/ritarock/quickquery/infrastructure/csv"
	"github.com/ritarock/quickquery/infrastructure/sqlite"
	"github.com/ritarock/quickquery/model"
)

type Handler struct {
	formatter *fomatter.Formatter
}

func NewHandler() *Handler {
	return &Handler{formatter: fomatter.NewFormatter()}
}

func (h *Handler) Run(args []string) error {
	if err := validateArgs(args); err != nil {
		return err
	}

	query, err := model.NewQuery(args[0])
	if err != nil {
		return err
	}

	reader := csv.NewReader()
	queryExecutor, err := sqlite.NewExecutor()
	if err != nil {
		return err
	}
	executor := application.NewExecutor(reader, queryExecutor)

	result, err := executor.Execute(query)
	if err != nil {
		return err
	}

	fmt.Println(h.formatter.Format(result))

	return nil
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
