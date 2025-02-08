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
				case "select", "from", "where":
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
				case "select", "from", "where":
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

func (q Query) GetWhere() []string {
	conditions := []string{}
	whereFound := false
	for i := 0; i < len(q.Clauses); i++ {
		if q.Clauses[i] == "WHERE" {
			whereFound = true
			continue
		}
		if whereFound {
			conditions = append(conditions, q.Clauses[i])
		}
	}

	return conditions
}
