package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 3: Auto ID
	Instead of manual ID:
		- generate automatically
*/

/*

	1. global ctr = 1

*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

var ctr int = 1

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {

		json.NewEncoder(w).Encode(map[string]any{
			"message": "users found successfully",
			"users":   users,
		})

	} else if r.Method == http.MethodPost {
		var user User

		json.NewDecoder(r.Body).Decode(&user)

		user.ID = ctr
		ctr++

		users = append(users, user)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user added successfully",
		})

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
