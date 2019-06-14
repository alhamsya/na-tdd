package main

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mock, _ = sqlmock.New()

	// Uncomment to test db connection, don't forget to install db first
	// initDB()
	defer db.Close()

	code := m.Run()
	os.Exit(code)
}
func TestShouldRead(t *testing.T) {
	test := "testing"
	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(3, test)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data := readFromDB(3)
	if data != test {
		t.Errorf("should return values:%s got:%s", test, data)
	}
}

func TestShouldInsert(t *testing.T) {
	test := "testing"
	mock.ExpectExec("INSERT INTO users").
		WillReturnResult(sqlmock.NewResult(0, 1))

	insertToDB(test)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectation not met: %s", err)
	}
}
