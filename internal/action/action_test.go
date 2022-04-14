package action

import (
	"testing"
)

const (
	SELECT_QUERY  = "select * from sample.csv"
	INSERT_QUERY1 = "insert into sample.csv values (1, user1, name1)"
	INSERT_QUERY2 = "insert into sample.csv (id, user) values (1, user1)"
	UPDATE_QUERY1 = "update sample.csv set id = 2 where id = 1"
	UPDATE_QUERY2 = "update sample.csv set id = 2, name = name2 where id = 1"
	DELETE_QUERY  = "delete from sample.csv where id = 1"
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
		want := "sample.csv"
		for _, query := range queries {
			got, _ := parseFileName(query)
			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
	t.Run("failed", func(t *testing.T) {
		_, gotErr := parseFileName("select * from sample.cs")
		wantErr := true
		if (gotErr != nil) != wantErr {
			t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
		}
	})
}

func Test_formingQuery(t *testing.T) {
	assert := func(t *testing.T, gotQuery, wantQuery, gotDml, wantDml string) {
		t.Helper()
		if gotQuery != wantQuery {
			t.Errorf("gotQuery %v, wantQuery %v", gotQuery, wantQuery)
		}
		if gotDml != wantDml {
			t.Errorf("gotDml %v, wantDml %v", gotDml, wantDml)
		}
	}

	t.Run("select query", func(t *testing.T) {
		gotQuery, gotDml := formingQuery(SELECT_QUERY)
		wantQuery := "SELECT * FROM sample"
		wantDml := "SELECT"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})
	t.Run("insert query1", func(t *testing.T) {
		gotQuery, gotDml := formingQuery(INSERT_QUERY1)
		wantQuery := "INSERT INTO sample VALUES (1, user1, name1)"
		wantDml := "INSERT"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})
	t.Run("insert query2", func(t *testing.T) {
		gotQuery, gotDml := formingQuery(INSERT_QUERY2)
		wantQuery := "INSERT INTO sample (id, user) VALUES (1, user1)"
		wantDml := "INSERT"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})
	t.Run("update query1", func(t *testing.T) {
		gotQuery, gotDml := formingQuery(UPDATE_QUERY1)
		wantQuery := "UPDATE sample SET id = 2 WHERE id = 1"
		wantDml := "UPDATE"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})
	t.Run("update query2", func(t *testing.T) {
		gotQuery, gotDml := formingQuery(UPDATE_QUERY2)
		wantQuery := "UPDATE sample SET id = 2, name = name2 WHERE id = 1"
		wantDml := "UPDATE"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})
	t.Run("delete query", func(t *testing.T) {
		gotQuery, gotDml := formingQuery(DELETE_QUERY)
		wantQuery := "DELETE FROM sample WHERE id = 1"
		wantDml := "DELETE"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})
}
