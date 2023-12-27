package main

import (
	"strings"
)

func (q Query) getWhere() (whereClause WhereClause, found bool) {
	var condition string
	for i, v := range q {
		if v == "WHERE" {
			condition = strings.Join(q[i+1:], "")
		}
	}
	if condition == "" {
		return whereClause, false
	}

	if strings.Contains(condition, ">=") {
		return makeWhereClause(whereClause, condition, ">="), true
	}
	if strings.Contains(condition, "<=") {
		return makeWhereClause(whereClause, condition, "<="), true
	}
	if strings.Contains(condition, "=") {
		return makeWhereClause(whereClause, condition, "="), true
	}
	if strings.Contains(condition, ">") {
		return makeWhereClause(whereClause, condition, ">"), true
	}
	if strings.Contains(condition, "<") {
		return makeWhereClause(whereClause, condition, "<"), true
	}

	return whereClause, false
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
			case ">":
				if mm[key] > value {
					newMapper = append(newMapper, mm)
				}
			case "<":
				if mm[key] < value {
					newMapper = append(newMapper, mm)
				}
			case ">=":
				if mm[key] >= value {
					newMapper = append(newMapper, mm)
				}
			case "<=":
				if mm[key] <= value {
					newMapper = append(newMapper, mm)
				}
			}
		}
	}

	return newMapper
}

func makeWhereClause(whereClause WhereClause, condition string, comparisonOperatorsSet string) WhereClause {
	clause := []string{}
	split := strings.Split(condition, comparisonOperatorsSet)
	clause = append(clause, []string{split[0], comparisonOperatorsSet, split[1]}...)
	whereClause = append(whereClause, clause)
	return whereClause
}
