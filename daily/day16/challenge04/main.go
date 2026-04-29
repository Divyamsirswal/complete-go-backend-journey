package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 5 - Monk
	Create helper:
		- findUserByID(id string) (User, int, error)
	reuse in GET, PUT, DELETE
*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

// generalize response here instead of typing after n after in functions
func respondJSON(w http.ResponseWriter, status int, data any, success string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"success": success,
		"data":    data,
	})
}

func findUserByID(id string) (User, int, error) {
	for i := range users {
		if fmt.Sprint(users[i].ID) == id {
			return users[i], i, nil
		}
	}
	return User{}, 0, fmt.Errorf("user not found")
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	if len(users) == 0 {
		respondJSON(w, http.StatusNotFound, "user not found", "false")
		return
	}
	respondJSON(w, http.StatusOK, users, "true")
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

	respondJSON(w, http.StatusCreated, user, "true")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var updatedUser User

	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var _, i, err2 = findUserByID(id)
	if err2 != nil {
		respondJSON(w, http.StatusNotFound, err2, "false")
		return
	}

	users[i].Name = updatedUser.Name
	users[i].Age = updatedUser.Age

	respondJSON(w, http.StatusOK, users[i], "true")

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, idx, err := findUserByID(id)
	if err != nil {
		respondJSON(w, http.StatusNotFound, err, "false")
		return
	}
	var before []User
	var after []User

	if idx > 0 {
		before = users[:idx]
	}
	if idx+1 < len(users) {
		after = users[idx+1:]
	}
	before = append(before, after...)
	users = before

	respondJSON(w, http.StatusOK, "user deleted successfully", "true")
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
