package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 5 - Monk
	Return response format:
		{
		  "success": true,
		  "data": {...}
		}
*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func getUsers(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "true",
		"data":    users,
	})
}

func getUserByID(w http.ResponseWriter, id string) {
	for i := range users {
		if fmt.Sprint(users[i].ID) == id {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"success": "true",
				"data":    users[i],
			})

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "false",
		"data":    "no user found",
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	user.ID = len(users) + 1

	if user.Name == "" {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"data":    "Name is required",
		})
		return
	}

	if user.Age < 0 {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "false",
			"data":    "Age must be positive",
		})
		return
	}

	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "true",
		"data":    user,
	})
}

func updateUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var newUser User

	json.NewDecoder(r.Body).Decode(&newUser)

	for i := range users {
		if fmt.Sprint(users[i].ID) == id {
			users[i].Name = newUser.Name
			users[i].Age = newUser.Age

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"success": "true",
				"data":    users[i],
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "false",
		"data":    "no user found",
	})

}

func deleteUserByID(w http.ResponseWriter, r *http.Request) {
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

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"success": "true",
				"data":    "user deleted successfully",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"success": "false",
		"data":    "user not found",
	})

}

func userHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		{
			id := r.URL.Query().Get("id")

			if id == "" {
				getUsers(w)
			} else {
				getUserByID(w, id)
			}
		}
	case http.MethodPost:
		{
			createUser(w, r)
		}
	case http.MethodPut:
		{
			updateUserByID(w, r)
		}
	case http.MethodDelete:
		{
			deleteUserByID(w, r)
		}
	default:
		{
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
