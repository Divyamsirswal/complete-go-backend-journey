package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 3: Query Filter
	URL - /users?name=Aman
		- return only that user

*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		val := r.URL.Query()
		queryLength := len(val)

		if queryLength == 0 {

			json.NewEncoder(w).Encode(users)

		} else {

			name := r.URL.Query().Get("name")

			if name == "" {
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Key:name is not present",
				})
			}

			for _, user := range users {
				if user.Name == name {
					json.NewEncoder(w).Encode(map[string]any{
						"user":    user,
						"message": "user found sucessfully",
					})
				}
			}

		}

	} else if r.Method == http.MethodPost {

		var user User

		json.NewDecoder(r.Body).Decode(&user)

		users = append(users, user)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "User Added Successfully",
		})
	}
}

func main() {
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
