package entity

type Order struct {
	column    string
	condition string
}

func (q Query) GetOrder() Order {
	var order Order
	orderIndex := -1

	for i, v := range q.Clauses {
		if v == "ORDER" {
			orderIndex = i
			break
		}
	}

	if orderIndex > 0 {
		order.column = q.Clauses[orderIndex+2]
		if len(q.Clauses[orderIndex:]) == 3 {
			order.condition = "ASC"
			return order
		}
		switch q.Clauses[orderIndex+3] {
		case "ASC":
			order.condition = "ASC"
		case "DESC":
			order.condition = "DESC"
		default:
			order.condition = "ASC"
		}
	}

	return order
}

func (o Order) IsOrder() bool {
	if o == (Order{}) {
		return false
	}
	return true
}
