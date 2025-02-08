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
				case "select", "from", "where", "and", "order", "by", "desc", "asc", "limit":
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
				case "select", "from", "where", "and", "order", "by", "desc", "asc", "limit":
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
		case "select", "from", "where", "and", "order", "by", "desc", "asc", "limit":
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
