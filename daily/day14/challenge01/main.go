package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 1: Basic Users API
		- Implement GET + POST
		- Store users in slice
*/

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Age       int    `json:"age"`
}

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		if len(users) == 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "No Users Found",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"message": "users found successfully",
			"users":   users,
		})

	} else if r.Method == http.MethodPost {

		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		users = append(users, user)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user added successfully",
		})

	} else {

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
