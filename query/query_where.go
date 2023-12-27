package query

import (
	"regexp"
	"strings"
)

type WhereQ [][]string

func (q Query) GetWhere() (whereQ WhereQ) {
	var condition string
	for i, v := range q {
		if v == "WHERE" {
			condition = strings.Join(q[i+1:], "")
		}
	}
	if condition == "" {
		return nil
	}

	if strings.Contains(condition, "ORDERBY") {
		re := regexp.MustCompile(`ORDERBY.*$`)
		condition = re.ReplaceAllString(condition, "")
	}

	if strings.Contains(condition, "AND") {
		for _, cond := range strings.Split(condition, "AND") {
			if strings.Contains(cond, ">=") {
				whereQ = makeCondition(whereQ, cond, ">=")
				continue
			}
			if strings.Contains(cond, "<=") {
				whereQ = makeCondition(whereQ, cond, "<=")
				continue
			}
			if strings.Contains(cond, "=") {
				whereQ = makeCondition(whereQ, cond, "=")
				continue
			}
			if strings.Contains(cond, ">") {
				whereQ = makeCondition(whereQ, cond, ">")
				continue
			}
			if strings.Contains(cond, "<") {
				whereQ = makeCondition(whereQ, cond, "<")
				continue
			}
		}
		return whereQ
	}

	if strings.Contains(condition, ">=") {
		return makeCondition(whereQ, condition, ">=")
	}
	if strings.Contains(condition, "<=") {
		return makeCondition(whereQ, condition, "<=")
	}
	if strings.Contains(condition, "=") {
		return makeCondition(whereQ, condition, "=")
	}
	if strings.Contains(condition, ">") {
		return makeCondition(whereQ, condition, ">")
	}
	if strings.Contains(condition, "<") {
		return makeCondition(whereQ, condition, "<")
	}

	return whereQ
}

func makeCondition(whereQ WhereQ,
	condition string, operator string) WhereQ {
	q := []string{}
	split := strings.Split(condition, operator)
	q = append(q, []string{split[0], operator, split[1]}...)
	whereQ = append(whereQ, q)
	return whereQ
}
