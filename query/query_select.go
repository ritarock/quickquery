package query

import "strings"

type SelectQ []string

func (q Query) GetSelect() SelectQ {
	var selectQ []string
	for _, v := range q {
		if v == "SELECT" {
			continue
		}
		if v == "FROM" {
			break
		}

		for _, vv := range strings.Split(v, ",") {
			col := strings.TrimSpace(vv)
			if col == "" {
				continue
			}
			selectQ = append(selectQ, col)
		}
	}
	return selectQ
}

func (s SelectQ) IsAll() bool {
	return s[0] == "*"
}
