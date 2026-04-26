package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 5 - Monk
	Add:
		- /users?id=1
	- return specific user
*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		// /users?id=1

		id := r.URL.Query().Get("id")

		for _, user := range users {
			if fmt.Sprint(user.ID) == id {

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]any{
					"message": "user found successfully",
					"user":    user,
				})
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user not found",
		})

	} else if r.Method == http.MethodPost {

		var user User

		json.NewDecoder(r.Body).Decode(&user)

		users = append(users, user)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "user added successfully",
		})

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
