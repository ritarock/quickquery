package query

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

func TestFormingQuery(t *testing.T) {
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
		gotQuery, gotDml := FormingQuery(SELECT_QUERY)
		wantQuery := "SELECT * FROM test"
		wantDml := "SELECT"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})

	t.Run("insert query1", func(t *testing.T) {
		gotQuery, gotDml := FormingQuery(INSERT_QUERY1)
		wantQuery := "INSERT INTO test VALUES (2, user2, name2)"
		wantDml := "INSERT"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})

	t.Run("insert query2", func(t *testing.T) {
		gotQuery, gotDml := FormingQuery(INSERT_QUERY2)
		wantQuery := "INSERT INTO test (id, user) VALUES (2, user2)"
		wantDml := "INSERT"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})

	t.Run("update query1", func(t *testing.T) {
		gotQuery, gotDml := FormingQuery(UPDATE_QUERY1)
		wantQuery := "UPDATE test SET id = 2 WHERE id = 1"
		wantDml := "UPDATE"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})

	t.Run("update query2", func(t *testing.T) {
		gotQuery, gotDml := FormingQuery(UPDATE_QUERY2)
		wantQuery := "UPDATE test SET id = 2, name = name2 WHERE id = 1"
		wantDml := "UPDATE"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})

	t.Run("delete query", func(t *testing.T) {
		gotQuery, gotDml := FormingQuery(DELETE_QUERY)
		wantQuery := "DELETE FROM test WHERE id = 1"
		wantDml := "DELETE"
		assert(t, gotQuery, wantQuery, gotDml, wantDml)
	})
}

func TestConvertInsertQuery(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("insert query1", func(t *testing.T) {
		got := ConvertInsertQuery(INSERT_QUERY1, "test.csv")
		want := "INSERT INTO test VALUES (\"2\",\"user2\",\"name2\")"
		assert(t, got, want)
	})
	t.Run("insert query2", func(t *testing.T) {
		got := ConvertInsertQuery(INSERT_QUERY2, "test.csv")
		want := "INSERT INTO test (id,user) VALUES (\"2\",\"user2\")"
		assert(t, got, want)
	})
}

func Test_insertColumnsValidation(t *testing.T) {
	t.Run("value=false", func(t *testing.T) {
		got := insertQueryValidation("(id, user)", false)
		want := "(id,user)"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("value=true", func(t *testing.T) {
		got := insertQueryValidation("(1, user1)", true)
		want := "(\"1\",\"user1\")"
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestConvertUpdateQuery(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("update query1", func(t *testing.T) {
		got := ConvertUpdateQuery("UPDATE test SET id = 2 WHERE id = 1")
		want := "UPDATE test SET id = \"2\" WHERE id = \"1\""
		assert(t, got, want)
	})
	t.Run("update query2", func(t *testing.T) {
		got := ConvertUpdateQuery("UPDATE test SET id = 2, name = name2 WHERE id = 1")
		want := "UPDATE test SET id = \"2\",name = \"name2\" WHERE id = \"1\""
		assert(t, got, want)
	})
}

func Test_updateQueryValidation(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("id = 1", func(t *testing.T) {
		got := updateQueryValidation("id = 1")
		want := "id = \"1\""
		assert(t, got, want)
	})
	t.Run("id = 1, user = 2", func(t *testing.T) {
		got := updateQueryValidation("id = 1, user = 2")
		want := "id = \"1\",user = \"2\""
		assert(t, got, want)
	})
}

func TestConvertDeleteQuery(t *testing.T) {
	got := ConvertDeleteQuery("DELETE FROM test WHERE id = 1")
	want := "DELETE FROM test WHERE id = \"1\""
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_deleteQueryValidation(t *testing.T) {
	got := deleteQueryValidation("id = 1")
	want := "id = \"1\""
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
