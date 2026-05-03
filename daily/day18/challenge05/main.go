package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

/*
	Challenge 5 : Monk

	Implement:
		GET → from DB
		POST → insert to DB
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

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		users := getUsers()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]any{
			"message": "users fetched successfully",
			"users":   users,
		})

	} else if r.Method == http.MethodPost {

		w.Header().Set("Content-Type", "application/json")

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		createUser(user.Name, user.Age)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "user created successfully",
		})

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	http.HandleFunc("/users", userHandler)

}
