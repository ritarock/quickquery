package entity

type Where [][]string

func (q Query) GetWhere() Where {
	conditions := [][]string{}
	whereIndex := -1

	for i, v := range q.Clauses {
		if v == "WHERE" {
			whereIndex = i + 1
			break
		}
	}

	if whereIndex > 0 {
		var where []string
		for _, v := range q.Clauses[whereIndex:] {
			if v == "AND" {
				continue
			}
			where = append(where, v)
		}

		for i := 0; i < len(where); i += 3 {
			if where[i] == "ORDER" {
				break
			}
			if i+2 >= len(where) {
				break
			}
			current := []string{}

			key := where[i]
			operator := where[i+1]
			value := where[i+2]

			current = append(current, key, operator, value)
			conditions = append(conditions, current)
		}

	}

	return conditions
}

func (w Where) IsWhere() bool {
	if len(w) > 0 {
		return true
	}
	return false
}
