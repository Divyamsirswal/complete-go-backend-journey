package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 4: Validation
	If:
		- name empty → error
		- age < 0 → error
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "" {
		fmt.Fprintln(w, "Error : Name can't be empty")
		return
	}
	if user.Age < 0 {
		fmt.Fprintln(w, "Error : Age must be valid")
		return
	}

	fmt.Fprintln(w, user)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
