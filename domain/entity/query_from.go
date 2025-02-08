package entity

import "errors"

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
