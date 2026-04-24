package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 2: Slice JSON
	Return:
		[
		  {"name":"Aman","age":21},
		  {"name":"Rahul","age":25}
		]
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			Name: "Aman",
			Age:  21,
		},
		{
			Name: "Rahul",
			Age:  25,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on port 8080....")
	http.ListenAndServe(":8080", nil)
}
