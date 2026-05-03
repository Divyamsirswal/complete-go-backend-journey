package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

/*
	Challenge 1: Connect DB
		- connect PostgreSQL
		- print success
*/

var db *sql.DB

func main() {

	connStr := os.Getenv("DATABASE_URL")

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

}
