package query

import (
	"regexp"
	"strings"
)

func FormingQuery(input string) (quey, dml string) {
	q := []string{}
	for _, v := range strings.Split(strings.TrimSpace(input), " ") {
		switch v {
		case "select", "insert", "update", "delete", "from", "values", "into", "set", "where":
			q = append(q, strings.ToUpper(v))
		default:
			if strings.Contains(v, ".csv") {
				q = append(q, v[:len(v)-4])
				continue
			}
			q = append(q, v)
		}
	}

	return strings.Join(q, " "), q[0]
}

func ConvertInsertQuery(query, fileName string) string {
	pt := regexp.MustCompile(`\(.*?\)`)
	insertQ := pt.FindAllString(query, -1)
	if len(insertQ) > 1 {
		return "INSERT INTO " +
			fileName[:len(fileName)-4] + " " +
			insertQueryValidation(insertQ[0], false) +
			" VALUES " +
			insertQueryValidation(insertQ[1], true)
	} else {
		return "INSERT INTO " +
			fileName[:len(fileName)-4] +
			" VALUES " +
			insertQueryValidation(insertQ[0], true)
	}
}

func insertQueryValidation(str string, value bool) string {
	str = strings.Replace(str, "(", "", 1)
	str = strings.Replace(str, ")", "", 1)
	var s []string
	for _, v := range strings.Split(str, ",") {
		if value {
			s = append(s, "\""+strings.TrimSpace(v)+"\"")
		} else {
			s = append(s, strings.TrimSpace(v))
		}
	}

	return "(" + strings.Join(s, ",") + ")"
}

func ConvertUpdateQuery(query string) string {
	q := query
	pt1 := regexp.MustCompile(`^.*SET `)
	pt2 := regexp.MustCompile(` WHERE .*$`)
	q = pt1.ReplaceAllString(q, "")
	q = pt2.ReplaceAllString(q, "")

	return pt1.FindString(query) + updateQueryValidation(q) + " " + updateQueryValidation(pt2.FindString(query))
}

func updateQueryValidation(str string) string {
	var s []string
	if strings.Contains(str, ",") {
		for _, v := range strings.Split(str, ",") {
			s = append(s,
				strings.TrimSpace(strings.Split(v, "=")[0])+
					" = \""+
					strings.TrimSpace(strings.Split(v, "=")[1])+
					"\"")
		}
		return strings.Join(s, ",")
	} else {
		s = append(s,
			strings.TrimSpace(strings.Split(str, "=")[0])+
				" = \""+
				strings.TrimSpace(strings.Split(str, "=")[1])+
				"\"")
		return strings.Join(s, ",")
	}
}

func ConvertDeleteQuery(query string) string {
	q := query
	pt := regexp.MustCompile(`^DELETE FROM .* WHERE`)
	q = pt.ReplaceAllString(q, "")

	return pt.FindString(query) + " " + deleteQueryValidation(q)
}

func deleteQueryValidation(str string) string {
	var s []string
	s = append(s,
		strings.TrimSpace(strings.Split(str, "=")[0])+
			" = \""+
			strings.TrimSpace(strings.Split(str, "=")[1])+
			"\"")

	return strings.Join(s, " ")
}
