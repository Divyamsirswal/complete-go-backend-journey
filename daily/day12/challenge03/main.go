package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Challenge 3: Accept JSON
	Send data:
		- {"name":"Aman","age":21}
	- print it in server
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on port 8080..")
	http.ListenAndServe(":8080", nil)
}
