package query

import (
	"fmt"
	"sort"
	"strings"
)

type Mapper []map[string]string

func (t Table) ToMap() Mapper {
	headers := t[0]
	var mapper Mapper
	for i := range t {
		if i == 0 {
			continue
		}
		m := make(map[string]string)
		for j := range headers {
			m[headers[j]] = t[i][j]
		}
		mapper = append(mapper, m)
	}
	return mapper
}

func (m Mapper) Result(selectQ SelectQ, whereQ WhereQ, orderQ OrderQ) {
	if len(whereQ) != 0 {
		m = m.adaptedWhere(whereQ)
	}
	var table Table
	table = append(table, selectQ)
	for _, row := range m {
		var line []string
		for _, header := range selectQ {
			line = append(line, row[header])
		}
		table = append(table, line)
	}

	var index int
	var order string
	for k, v := range orderQ {
		for i, vv := range table[0] {
			if k == vv {
				index = i
				break
			}
		}
		order = v
	}
	t := table[1:]
	sort.Slice(t, func(i, j int) bool {
		if order == "ASC" {
			return t[i][index] < t[j][index]
		} else {
			return t[i][index] > t[j][index]
		}
	})

	fmt.Println(strings.Join(table[0], ", "))
	for _, v := range table[1:] {
		fmt.Println(strings.Join(v, ", "))
	}
}

func (m Mapper) adaptedWhere(whereQ WhereQ) Mapper {
	for _, w := range whereQ {
		key, condition, value := w[0], w[1], w[2]
		m = m.applyCondition(key, condition, value)
	}
	return m
}

func (m Mapper) applyCondition(key, condition, value string) Mapper {
	var newMapper Mapper
	for _, mm := range m {
		switch condition {
		case ">=":
			if mm[key] >= value {
				newMapper = append(newMapper, mm)
			}
		case "<=":
			if mm[key] <= value {
				newMapper = append(newMapper, mm)
			}
		case "=":
			if mm[key] == value {
				newMapper = append(newMapper, mm)
			}
		case ">":
			if mm[key] > value {
				newMapper = append(newMapper, mm)
			}
		case "<":
			if mm[key] < value {
				newMapper = append(newMapper, mm)
			}
		}
	}
	return newMapper
}
