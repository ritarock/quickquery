package action

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/ritarock/quickquery/internal/db"
	"github.com/urfave/cli/v2"
)

func RunQuey(ctx *cli.Context) error {
	fileName, err := parseFileName(ctx.Args().First())
	if err != nil {
		return err
	}
	if err := db.CreateTable(fileName); err != nil {
		return err
	}
	defer os.Remove(".sqlite.db")

	query := formingQuery(ctx.Args().First())
	dml := query[0]
	switch dml {
	case "SELECT":
		if err := db.SelectQuery(strings.Join(query, " ")); err != nil {
			return err
		}
	case "INSERT":
		q := convertInsertQuery(strings.Join(query, " "), fileName)
		if err := db.InsertQuery(q); err != nil {
			return err
		}
		db.OutputCsv(fileName)
	case "UPDATE":
		q := convertUpdateQuery(strings.Join(query, " "), fileName)
		if err := db.UpdateQuery(q); err != nil {
			return err
		}
		db.OutputCsv(fileName)
	case "DELETE":
		q := convertDeleteQuery(strings.Join(query, " "), fileName)
		if err := db.DeleteQuery(q); err != nil {
			return err
		}
		db.OutputCsv(fileName)
	default:
		return errors.New("cannot parse query")
	}
	return nil
}

func parseFileName(arg string) (string, error) {
	path := strings.Split(strings.TrimSpace(arg), " ")
	for _, v := range path {
		if strings.Contains(v, ".csv") {
			return v, nil
		}
	}
	return "", errors.New("not find csv file")
}

func formingQuery(arg string) []string {
	q := []string{}

	for _, v := range strings.Split(strings.TrimSpace(arg), " ") {
		switch v {
		case "select", "from", "insert", "values", "into", "update", "set", "where", "delete":
			q = append(q, strings.ToUpper(v))
		default:
			if strings.Contains(v, ".csv") {
				q = append(q, v[:len(v)-4])
				continue
			}
			q = append(q, v)
		}
	}
	return q
}

func convertInsertQuery(query, fileName string) string {
	pt := regexp.MustCompile(`\(.*?\)`)
	insertQ := pt.FindAllString(query, -1)
	if len(insertQ) > 1 {
		return "INSERT INTO " +
			fileName[:len(fileName)-4] + " " +
			insertColumnsValidation(insertQ[0]) +
			" VALUES " +
			insertValuesValidation(insertQ[1])
	} else {
		return "INSERT INTO " +
			fileName[:len(fileName)-4] +
			" VALUES " +
			insertValuesValidation(insertQ[0])
	}

}

func convertUpdateQuery(query, fileName string) string {
	q := query
	pt1 := regexp.MustCompile(`^.*SET `)
	pt2 := regexp.MustCompile(` WHERE .*$`)
	q = pt1.ReplaceAllString(q, "")
	q = pt2.ReplaceAllString(q, "")
	updateValuesValidation(q)

	return pt1.FindString(query) + updateValuesValidation(q) + pt2.FindString(query)

}

func insertColumnsValidation(str string) string {
	str = strings.Replace(str, "(", "", 1)
	str = strings.Replace(str, ")", "", 1)
	var s []string
	for _, v := range strings.Split(str, ",") {
		s = append(s, strings.TrimSpace(v))
	}

	return "(" + strings.Join(s, ",") + ")"
}

func insertValuesValidation(str string) string {
	str = strings.Replace(str, "(", "", 1)
	str = strings.Replace(str, ")", "", 1)
	var s []string
	for _, v := range strings.Split(str, ",") {
		s = append(s, "\""+strings.TrimSpace(v)+"\"")
	}
	return "(" + strings.Join(s, ",") + ")"
}

func updateValuesValidation(str string) string {
	var s []string
	if strings.Contains(str, ",") {
		for _, v := range strings.Split(str, ",") {
			s = append(s,
				strings.Split(v, "=")[0]+
					" = \""+
					strings.TrimSpace(strings.Split(v, "=")[1])+
					"\"")
		}
		return strings.Join(s, ",")
	} else {
		s = append(s,
			strings.Split(str, "=")[0]+
				" = \""+
				strings.TrimSpace(strings.Split(str, "=")[1])+
				"\"")
		return strings.Join(s, ",")
	}
}

func convertDeleteQuery(query, fileName string) string {
	q := query
	pt := regexp.MustCompile(`^DELETE FROM .* WHERE`)
	q = pt.ReplaceAllString(q, "")

	return pt.FindString(query) + deleteValuesValidation(q)
}

func deleteValuesValidation(str string) string {
	var s []string
	s = append(s,
		strings.Split(str, "=")[0]+
			" = \""+
			strings.TrimSpace(strings.Split(str, "=")[1])+
			"\"")

	return strings.Join(s, " ")
}
