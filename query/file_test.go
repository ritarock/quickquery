package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTable(t *testing.T) {
	tests := []struct {
		path string
		want Table
	}{
		{
			path: "../sample.csv",
			want: [][]string{
				{"id", "name", "user"},
				{"1", "name1", "user1"},
				{"2", "name2", "user2"},
				{"3", "name3", "user3"},
				{"4", "name4", "user4"},
				{"5", "name5", "user5"},
			},
		},
	}

	for _, test := range tests {
		got, err := ReadTable(test.path)
		assert.NoError(t, err)
		assert.Equal(t, test.want, got)
	}
}

func Test_readFile(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{
			path: "../sample.csv",
			want: "id, name, user\n1, name1, user1\n2, name2, user2\n3, name3, user3\n4, name4, user4\n5, name5, user5\n",
		},
	}

	for _, test := range tests {
		got, err := readFile(test.path)
		assert.NoError(t, err)
		assert.Equal(t, test.want, got)
	}

}
