package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

/*
	Challenge 4: Integrate with API
	Replace:
		users []User
	with DB calls
*/

var db *sql.DB

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUserTable() {

	cmd := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	); 
	`

	_, err := db.Exec(cmd)
	if err != nil {
		panic(err)
	}

	fmt.Println("users table created successfully")
}

func createUser(name string, age int) {
	cmd := `INSERT INTO users (name, age) VALUES 
			($1, $2)`

	_, err := db.Exec(cmd, name, age)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func getUsers() *sql.Rows {

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			panic(err)
		}
	}

	return rows

}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users := getUsers()

	fmt.Println(users)
}

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

	// createUserTable() -> only used once

	// createUser("user1", 12)
	// createUser("user2", 22)
	// createUser("user3", 42)
	// createUser("user4", 62)

	http.HandleFunc("/users", getUsersHandler)

}
