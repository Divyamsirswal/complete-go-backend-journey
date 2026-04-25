package main

import (
	"fmt"
	"net/http"
)

/*
	Challenge 5 - Monk
	Create route:
		- /user
	returns:
		- User: Aman, Age: 21
*/

func userHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintln(w, "User:", name, ", Age:", age)
}

func main() {
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
