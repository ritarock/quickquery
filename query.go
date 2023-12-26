package main

import (
	"errors"
	"strings"
)

type Clause string
type Query []Clause
type SelectClause []string
type Mapper []map[string]string

func (q Query) validate() error {
	if q[0] != "SELECT" {
		return errors.New("syntax error")
	}
	return nil
}

func (t Table) toMap() Mapper {
	headers := t[0]
	var mapper Mapper
	for i := range t {
		if i == 0 {
			continue
		}
		m := make(map[string]string)
		for j := range headers {
			m[headers[j]] = t[i][j]
		}
		mapper = append(mapper, m)
	}
	return mapper
}

func (q Query) getFileName() (string, error) {
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

func (q Query) getSelect() SelectClause {
	var selectClause SelectClause
	for _, v := range q {
		if v == "SELECT" {
			continue
		}
		if v == "FROM" {
			break
		}

		for _, vv := range strings.Split(string(v), ",") {
			col := strings.TrimSpace(vv)
			if strings.TrimSpace(col) == "" {
				continue
			}
			selectClause = append(selectClause, col)
		}
	}
	return selectClause
}

func (s SelectClause) isAll() bool {
	return s[0] == "*"
}

func (m Mapper) bySelect(selectClause SelectClause) Table {
	if len(m) == 0 {
		return nil
	}

	var table Table
	table = append(table, selectClause)
	for _, row := range m {
		var line []string
		for _, header := range selectClause {
			line = append(line, row[header])
		}
		table = append(table, line)
	}

	return table
}

func argToQuery(arg string) Query {
	tmp := strings.TrimSpace(arg)
	var query Query

	for _, v := range strings.Split(tmp, " ") {
		clause := strings.ToLower(v)
		switch clause {
		case "select", "from":
			query = append(query, Clause(strings.ToUpper(clause)))
		default:
			query = append(query, Clause(clause))
		}
	}

	return query
}
