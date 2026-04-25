package main

import (
	"fmt"
	"net/http"
)

// Basic server in go using net/http
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is my first server.")
}

func myProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintln(w, "Welcome to my profile")
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/myprofile", myProfileHandler)

	fmt.Println("Server is running on port 8080..")

	http.ListenAndServe(":8080", nil)
}
