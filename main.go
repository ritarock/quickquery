package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		fmt.Println("help")
		return
	}

	args := flag.Args()
	if err := validateArgs(args); err != nil {
		fmt.Println(err)
		return
	}

	query := argToQuery(args[0])
	if err := query.validate(); err != nil {
		fmt.Println(err)
		return
	}

	file, err := query.getFileName()
	if err != nil {
		fmt.Println(err)
		return
	}

	table, err := readFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	mapper := table.toMap()

	selectClause := query.getSelect()
	if selectClause.isAll() {
		selectClause = table[0]
	}

	whereClause, found := query.getWhere()
	if !found {
		result := mapper.bySelect(selectClause)
		result.Result()
	}
	mapper.byWhere(whereClause).bySelect(selectClause).Result()

}

func validateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("not enough args")
	}
	if len(args) >= 2 {
		return errors.New("too many args")
	}
	return nil
}
