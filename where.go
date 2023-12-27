package main

import (
	"fmt"
	"strings"
)

func (q Query) getWhere() (whereClause WhereClause, found bool) {
	var word string
	for i, v := range q {
		if v == "WHERE" {
			word = strings.Join(q[i+1:], "")
		}
	}
	if word == "" {
		return whereClause, false
	}

	clause := []string{}
	if strings.Contains(word, "=") {
		split := strings.Split(word, "=")
		clause = append(clause, []string{split[0], "=", split[1]}...)
	}

	whereClause = append(whereClause, clause)

	fmt.Println(whereClause)
	return whereClause, true
}

func (m Mapper) byWhere(whereClause WhereClause) Mapper {
	if len(m) == 0 {
		return nil
	}

	var newMapper Mapper
	for _, w := range whereClause {
		key, condition, value := w[0], w[1], w[2]
		for _, mm := range m {
			switch condition {
			case "=":
				if mm[key] == value {
					newMapper = append(newMapper, mm)
				}
			}
		}

	}

	return newMapper
}
