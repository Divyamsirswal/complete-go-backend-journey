package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 5 - Monk
	Create:
		 - /users/{id} GET
	return user by ID (manual parsing)
*/

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func postHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "user added successfully",
		"user":    user,
	})

}

func getHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path

	var req string = ""

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			break
		}
		req += string(path[i])
	}

	var id string = ""
	for i := len(req) - 1; i >= 0; i-- {
		id += string(req[i])
	}

	for _, user := range users {
		if fmt.Sprint(user.Id) == id {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "user found successfully",
				"Id":      user.Id,
				"Name":    user.Name,
				"Age":     user.Age,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "user not found",
	})

}

func main() {

	http.HandleFunc("/users", postHandler)
	http.HandleFunc("/users/{id}", getHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)

}
