package action

import (
	"testing"
)

const (
	SELECT_QUERY  = "select * from test.csv"
	INSERT_QUERY1 = "insert into test.csv values (2, user2, name2)"
	INSERT_QUERY2 = "insert into test.csv (id, user) values (2, user2)"
	UPDATE_QUERY1 = "update test.csv set id = 2 where id = 1"
	UPDATE_QUERY2 = "update test.csv set id = 2, name = name2 where id = 1"
	DELETE_QUERY  = "delete from test.csv where id = 1"
)

func Test_parseFileName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queries := []string{
			SELECT_QUERY,
			INSERT_QUERY1,
			INSERT_QUERY2,
			UPDATE_QUERY1,
			UPDATE_QUERY2,
			DELETE_QUERY,
		}
		want := "test.csv"
		for _, query := range queries {
			got, _ := parseFileName(query)
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})

	t.Run("failed", func(t *testing.T) {
		_, gotErr := parseFileName("select * from test.cs")
		wantErr := true
		if (gotErr != nil) != wantErr {
			t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
		}
	})
}

func Test_fileExists(t *testing.T) {
	got := fileExists("test.csv")
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
