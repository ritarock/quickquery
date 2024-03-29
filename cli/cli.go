package cli

import (
	"errors"
	"flag"
	"fmt"

	q "github.com/ritarock/quickquery/query"
)

func Run() error {
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		fmt.Println(HELP)
		return nil
	}

	args := flag.Args()
	if err := validateArgs(args); err != nil {
		return err
	}

	if args[0] == "help" {
		fmt.Println(HELP)
		return nil
	}

	query := q.ArgToQuery(args[0])
	if err := query.Validate(); err != nil {
		return err
	}

	file, err := query.GetFileName()
	if err != nil {
		return err
	}
	table, err := q.ReadTable(file)
	if err != nil {
		return err
	}

	mapper := table.ToMap()

	selectQ := query.GetSelect()
	if selectQ.IsAll() {
		// set header
		selectQ = table[0]
	}

	whereQ := query.GetWhere()
	orderQ := query.GetOrder(table[0])

	mapper.Result(selectQ, whereQ, orderQ)

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
