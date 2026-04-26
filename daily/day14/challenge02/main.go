package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 2: Validation
	If:
		- name empty
		- age < 0

	- return error (400)
*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		if len(users) == 0 {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "No users found",
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "users found successfully",
			"users":   users,
		})

	} else if r.Method == http.MethodPost {

		var user User

		json.NewDecoder(r.Body).Decode(&user)

		if user.Name == "" || user.Age < 0 {
			http.Error(w, "Invalid data", http.StatusBadRequest)
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
