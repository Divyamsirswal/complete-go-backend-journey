package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 1: Users API
	Route : /users
		- GET → return users
		- POST → add user
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(users)

	} else if r.Method == http.MethodPost {
		var user User

		json.NewDecoder(r.Body).Decode(&user)

		users = append(users, user)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "user added successfully",
		})

	}

}

func main() {
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
