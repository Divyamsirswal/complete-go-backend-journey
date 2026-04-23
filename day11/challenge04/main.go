package main

import (
	"fmt"
	"net/http"
)

/*
	Challenge 4: Query Parameter
	Try:
		- http://localhost:8080/?name=Aman
	Do:
		- take the name from the query and print it
*/

func homeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintln(w, "Hello", name)
}

func main() {

	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running on port 8080..")

	http.ListenAndServe(":8080", nil)
}
