package main

import (
	"fmt"
	"net/http"
)

/*
	Challenge 2: Multiple Routes
	Create routes:
		- / → Home
		- /about → About
		- /contact → Contact
*/

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to About Page")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to contact page")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Server is running on port 8080..")

	http.ListenAndServe(":8080", nil)
}
