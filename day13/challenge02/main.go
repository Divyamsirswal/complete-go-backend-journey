package main

import (
	"fmt"
	"net/http"
)

/*
	Challenge 2: Method Handling
	If wrong method:
		- Method not allowed

*/

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error : This HTTP Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Working fine")
}

func main() {
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
