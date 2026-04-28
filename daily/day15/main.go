package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

			// route -> GET /users -> i want to get all users

			json.NewEncoder(w).Encode(map[string]any{
				"message": "users fetched successfully",
				"users":   users,
			})

		}
	case http.MethodPost:
		{
			// route -> POST /users -> i want to add a new user

			var user User
			json.NewDecoder(r.Body).Decode(&user)

			user.ID = len(users) + 1
			users = append(users, user)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "user added successfully",
				"user":    user,
			})
		}
	case http.MethodPut:
		{

			// route -> PUT /users?id=value -> here i want to update info of particular user

			id := r.URL.Query().Get("id")

			var newInfo User

			json.NewDecoder(r.Body).Decode(&newInfo)

			var getUser int

			for i := range users {
				if fmt.Sprint(users[i].ID) == id {
					getUser = i
					users[i].Name = newInfo.Name
					users[i].Age = newInfo.Age
				}
			}

			json.NewEncoder(w).Encode(map[string]any{
				"message": "user updated successfully",
				"user":    users[getUser],
			})

		}
	case http.MethodDelete:
		{

			// route -> DELETE /users?id=value -> i have to delete particular user with id = value

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
					break
				}
			}

			json.NewEncoder(w).Encode(map[string]any{
				"message": "user deleted successfully",
				// "user":    users[getUser], // ******this is the problem i m so dumb after deleting i m trying to access it*****
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
