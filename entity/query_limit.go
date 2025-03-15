package entity

import "strconv"

type Limit int

func (q Query) GetLimit() Limit {
	var limit Limit
	limitIndex := -1

	for i, v := range q.Clauses {
		if v == "LIMIT" {
			limitIndex = i
			break
		}
	}

	if limitIndex == -1 {
		return limit
	}

	li, err := strconv.Atoi(q.Clauses[limitIndex+1])
	if err != nil {
		return limit
	}
	limit = Limit(li)

	return limit
}

func (l Limit) IsLimit() bool {
	if l > 0 {
		return true
	}
	return false
}
