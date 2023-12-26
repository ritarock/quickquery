package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readFile(t *testing.T) {
	tests := []struct {
		path string
		want Table
	}{
		{
			path: "./sample.csv",
			want: [][]string{
				{"id", "name", "user"},
				{"1", "name1", "user1"},
				{"2", "name2", "user2"},
				{"3", "name3", "user3"},
			},
		},
	}

	for _, test := range tests {
		got, err := readFile(test.path)
		assert.NoError(t, err)
		assert.Equal(t, test.want, got)
	}
}
