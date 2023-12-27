package query

import (
	"errors"
	"strings"
)

type Query []string

func ArgToQuery(arg string) Query {
	tmp := strings.TrimSpace(arg)
	var query Query

	for _, v := range strings.Split(tmp, " ") {
		clause := strings.ToLower(v)
		switch clause {
		case "select", "from", "where", "and":
			query = append(query, strings.ToUpper(clause))
		default:
			query = append(query, clause)
		}
	}

	return query
}

func (q Query) Validate() error {
	if q[0] != "SELECT" {
		return errors.New("syntax error")
	}
	return nil
}

func (q Query) GetFileName() (string, error) {
	var file string
	for i, v := range q {
		if v == "FROM" {
			file = string(q[i+1])
			break
		}
	}
	if file == "" {
		return "", errors.New("not exist file name")
	}

	return file, nil
}
