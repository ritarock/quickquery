package db

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateTable(t *testing.T) {
	defer os.Remove(".sqlite.db")
	gotErr := CreateTable("test.csv")
	success := func(fileName string) bool {
		_, err := os.Stat(fileName)
		return err == nil
	}(".sqlite.db")
	wantErr := false
	if (gotErr != nil) != wantErr {
		t.Errorf("gotErr %v, wantErr %v success %v", gotErr, wantErr, success)
	}
}
