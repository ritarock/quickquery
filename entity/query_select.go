package entity

type Select []string

func (q Query) GetSelect() Select {
	var columns []string
	for i := 1; i < len(q.Clauses); i++ {
		if q.Clauses[i] == "FROM" {
			break
		}
		if q.Clauses[i] != "," {
			columns = append(columns, q.Clauses[i])
		}
	}

	return columns
}
