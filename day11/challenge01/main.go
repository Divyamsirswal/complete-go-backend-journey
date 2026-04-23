package main

import (
	"fmt"
	"net/http"
)

/*
	Challenge 1: Basic Server
	Create server
	Print:
		- Welcome to my backend journey
*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my backend journey")
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
