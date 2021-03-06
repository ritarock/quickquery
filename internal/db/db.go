package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER      = "sqlite3"
	DATA_SOURCE = ".sqlite.db"
)

func CreateTable(fileName string) error {
	command := fmt.Sprintf(`sqlite3 .sqlite.db \
	".mode csv" \
	".import %v %v"`, fileName, fileName[:len(fileName)-4])

	_, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		os.Remove(".sqlite.db")
		return errors.New("failed create table")
	}
	return nil
}

func OutputCsv(fileName string) {
	command := fmt.Sprintf(`sqlite3 .sqlite.db \
	".headers on" \
	".mode csv" \
	".output %v" \
	"SELECT * FROM %v"`, fileName, fileName[:len(fileName)-4])

	exec.Command("sh", "-c", command).Output()
}

func SelectQuery(query string) error {
	db, err := sql.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		return errors.New("an error occurred")
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return errors.New("failed run query")
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("cannot find record")
	}

	ct, _ := rows.ColumnTypes()
	types := make([]reflect.Type, len(ct))
	for i, typ := range ct {
		types[i] = typ.ScanType()
	}

	values := make([]interface{}, len(ct))
	for i := range values {
		values[i] = reflect.New(types[i]).Interface()
	}

	for {
		err := rows.Scan(values...)
		if err != nil {
			return err
		}
		for _, v := range values {
			fmt.Printf("%v|", reflect.ValueOf(v).Elem().Field(0))
		}
		fmt.Println()
		if !rows.Next() {
			break
		}
	}

	return nil
}

func InsertQuery(query string) error {
	db, err := sql.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		return errors.New("an error occurred")
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		return errors.New("cannot insert query")
	}

	return nil
}

func UpdateQuery(query string) error {
	db, err := sql.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		return errors.New("an error occurred")
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		return errors.New("cannot update query")
	}

	return nil
}

func DeleteQuery(query string) error {
	db, err := sql.Open(DRIVER, DATA_SOURCE)
	if err != nil {
		return errors.New("an error occurred")
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		return errors.New("cannot delete query")
	}

	return nil
}
