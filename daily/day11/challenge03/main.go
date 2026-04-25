package main

import (
	"fmt"
	"net/http"
)

/*
	Challenge 3: Method Check
	If:
		- GET → print "GET request"
		- else → print "Other request"
*/

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "request")
	fmt.Fprintln(w, "Home Page")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
