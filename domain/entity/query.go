package entity

import (
	"errors"
	"strings"
)

type Query struct {
	Clauses []string
}

func NewQuery(arg string) Query {
	tmp := strings.TrimSpace(arg)
	var clauses []string

	var currentToken strings.Builder
	inQuote := false

	for i := 0; i < len(tmp); i++ {
		char := tmp[i]
		if char == '"' {
			inQuote = !inQuote
			continue
		}

		if char == ',' && !inQuote {
			if currentToken.Len() > 0 {
				token := currentToken.String()
				switch strings.ToLower(token) {
				case "select", "from", "where", "and":
					clauses = append(clauses, strings.ToUpper(token))
				default:
					clauses = append(clauses, token)
				}
				currentToken.Reset()
			}
			clauses = append(clauses, ",")
			continue
		}

		if char == ' ' && !inQuote {
			if currentToken.Len() > 0 {
				token := currentToken.String()
				switch strings.ToLower(token) {
				case "select", "from", "where", "and":
					clauses = append(clauses, strings.ToUpper(token))
				default:
					clauses = append(clauses, token)
				}
				currentToken.Reset()
			}
			continue
		}
		currentToken.WriteByte(char)
	}

	if currentToken.Len() > 0 {
		token := currentToken.String()
		switch strings.ToLower(token) {
		case "select", "from", "where":
			clauses = append(clauses, strings.ToUpper(token))
		default:
			clauses = append(clauses, token)
		}
	}

	return Query{Clauses: clauses}
}

func (q Query) Validate() error {
	if len(q.Clauses) == 0 || q.Clauses[0] != "SELECT" {
		return errors.New("syntax error: query must start with SELECT")
	}

	return nil
}

func (q Query) GetFile() (string, error) {
	var file string
	for i, v := range q.Clauses {
		if v == "FROM" && i+1 < len(q.Clauses) {
			file = q.Clauses[i+1]
			break
		}
	}
	if file == "" {
		return "", errors.New("syntax error: FROM must specify a file name")
	}

	return file, nil
}

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

func (q Query) GetWhere() [][]string {
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
