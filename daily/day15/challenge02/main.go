package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 2: Validation
		- name required
		- age must be positive

*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		{

			id := r.URL.Query().Get("id")

			if id == "" {

				json.NewEncoder(w).Encode(map[string]any{
					"message": "users fetched successfully",
					"users":   users,
				})

			} else {

				for i := range users {
					if fmt.Sprint(users[i].ID) == id {

						json.NewEncoder(w).Encode(map[string]any{
							"message": "user fetched successfully",
							"user":    users[i],
						})

						return
					}
				}

				json.NewEncoder(w).Encode(map[string]any{
					"message": "no user found",
				})

			}

		}
	case http.MethodPost:
		{

			var user User

			json.NewDecoder(r.Body).Decode(&user)

			user.ID = len(users) + 1

			if user.Name == "" {

				json.NewEncoder(w).Encode(map[string]string{
					"message": "Name is required",
				})
				return
			}

			if user.Age < 0 {

				json.NewEncoder(w).Encode(map[string]string{
					"message": "Age must be positive",
				})
				return
			}

			users = append(users, user)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "user added successfully",
				"user":    user,
			})

		}
	case http.MethodPut:
		{
			id := r.URL.Query().Get("id")

			var newUser User

			json.NewDecoder(r.Body).Decode(&newUser)

			for i := range users {
				if fmt.Sprint(users[i].ID) == id {
					users[i].Name = newUser.Name
					users[i].Age = newUser.Age

					json.NewEncoder(w).Encode(map[string]any{
						"message": "user updated successfully",
						"user":    users[i],
					})
					return
				}
			}

			json.NewEncoder(w).Encode(map[string]any{
				"message": "no user found",
			})

		}
	case http.MethodDelete:
		{

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

					json.NewEncoder(w).Encode(map[string]any{
						"message": "user deleted successfully",
					})
					return
				}
			}

			json.NewEncoder(w).Encode(map[string]string{
				"message": "user not found",
			})

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
