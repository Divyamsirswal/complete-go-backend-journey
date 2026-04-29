package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 2: Use Helper Function
	Implement:
		 - respondJSON()
*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

// generalize response here instead of typing after n after in functions
func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	if len(users) == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"data":    "no users found",
		})
		return
	}
	respondJSON(w, http.StatusOK, users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	respondJSON(w, http.StatusCreated, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var updatedUser User

	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range users {
		if fmt.Sprint(users[i].ID) == id {

			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age

			respondJSON(w, http.StatusOK, users[i])
			return
		}
	}

	respondJSON(w, http.StatusNotFound, "user not found")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for i := range users {
		if fmt.Sprint(users[i].ID) == id {
			var before []User
			var after []User

			if i > 0 {
				before = users[:i]
			}
			if i+1 < len(users) {
				after = users[i+1:]
			}
			before = append(before, after...)
			users = before

			respondJSON(w, http.StatusOK, "user deleted successfully")
			return
		}
	}
	respondJSON(w, http.StatusNotFound, "user not found")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			getUsers(w, r)

		}
	case http.MethodPost:
		{
			createUser(w, r)
		}
	case http.MethodPut:
		{
			updateUser(w, r)
		}
	case http.MethodDelete:
		{
			deleteUser(w, r)
		}
	default:
		{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func main() {

	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
