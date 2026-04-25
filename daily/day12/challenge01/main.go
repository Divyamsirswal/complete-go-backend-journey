package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 1: Return JSON
	Create route:
		- /user
	Return:
		- {
			"name":"Aman",
			"age":21
		  }
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Aman",
		Age:  12,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
