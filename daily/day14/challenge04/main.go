package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 4: Pretty JSON Response
	Use:
		- json.NewEncoder(w).SetIndent("", "  ")
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).SetIndent("", " ")

		json.NewEncoder(w).Encode(users)

	} else if r.Method == http.MethodPost {

		user := User{
			Name: "user1",
			Age:  12,
		}

		users = append(users, user)

		json.NewEncoder(w).SetIndent("", " ")

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
