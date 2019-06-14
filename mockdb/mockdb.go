package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
}
func insertToDB(name string) {
	stmt := fmt.Sprintf(`INSERT INTO users (name) VALUES ('%s')`, name)
	_, err := db.Exec(stmt)
	if err != nil {
		fmt.Println(err)
	}
}

func readFromDB(id int) string {
	var name string
	var dbid int
	stmt := fmt.Sprintf(`SELECT id, name FROM users WHERE id=$1;`)
	row := db.QueryRow(stmt, id)
	switch err := row.Scan(&dbid, &name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(dbid, name)
		return name
	default:
		fmt.Println(err)
	}
	return ""
}
func main() {
	initDB()
	insertToDB("testing_aja")
}
