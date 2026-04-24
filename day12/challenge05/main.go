package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 5 - Monk
	Create:
		- /users
	- Accept POST
	- Store in slice
	- Return all users
*/

type User struct {
	Name string
	Age  int
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var users []User

	json.NewDecoder(r.Body).Decode(&users)

	for _, user := range users {
		fmt.Println(user.Name, ":", user.Age)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
