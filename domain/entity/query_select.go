package entity

func (q Query) GetSelect() []string {
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
