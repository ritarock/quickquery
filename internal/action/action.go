package action

import (
	"errors"
	"os"
	"strings"

	"github.com/ritarock/quickquery/internal/db"
	qquery "github.com/ritarock/quickquery/internal/query"
)

func Run(arg string) error {
	fileName, err := parseFileName(arg)
	if err != nil {
		return err
	}
	if err := db.CreateTable(fileName); err != nil {
		return err
	}
	defer os.Remove(".sqlite.db")

	query, dml := qquery.FormingQuery(arg)
	switch dml {
	case "SELECT":
		if err := db.SelectQuery(query); err != nil {
			return err
		}
	case "INSERT":
		q := qquery.ConvertInsertQuery(query, fileName)
		if err := db.InsertQuery(q); err != nil {
			return err
		}
		db.OutputCsv(fileName)
	case "UPDATE":
		q := qquery.ConvertUpdateQuery(query)
		if err := db.UpdateQuery(q); err != nil {
			return err
		}
		db.OutputCsv(fileName)
	case "DELETE":
		q := qquery.ConvertDeleteQuery(query)
		if err := db.DeleteQuery(q); err != nil {
			return err
		}
		db.OutputCsv(fileName)
	}

	return nil
}

func parseFileName(input string) (string, error) {
	for _, v := range strings.Split(strings.TrimSpace(input), " ") {
		if strings.Contains(v, ".csv") && fileExists(v) {
			return v, nil
		}
	}

	return "", errors.New("not find csv file")
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
