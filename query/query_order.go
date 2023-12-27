package query

import (
	"strings"
)

type OrderQ map[string]string

func (q Query) GetOrder(header []string) (orderQ OrderQ) {
	var condition string
	for i, v := range q {
		if v == "ORDER" {
			condition = strings.Join(q[i+2:], "")
		}
	}
	if condition == "" {
		return nil
	}

	m := map[string]string{}

	if strings.Contains(condition, "ASC") {
		key := strings.Replace(condition, "ASC", "", 1)
		if !existKeyInHeader(key, header) {
			return nil
		}
		m[key] = "ASC"
		return m
	}
	if strings.Contains(condition, "DESC") {
		key := strings.Replace(condition, "DESC", "", 1)
		if !existKeyInHeader(key, header) {
			return nil
		}
		m[key] = "DESC"
		return m
	}

	if !existKeyInHeader(condition, header) {
		return nil
	}
	m[condition] = "ASC"
	return m
}

func existKeyInHeader(key string, header []string) bool {
	for _, v := range header {
		if v == key {
			return true
		}
	}
	return false
}
